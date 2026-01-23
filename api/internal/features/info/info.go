package info

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// --- Europe PMC JSONの必要部分だけ定義 ---
type epmcResp struct {
	ResultList struct {
		Result []struct {
			ID     string `json:"id"`
			Source string `json:"source"` //  PMC
			Title  string `json:"title"`
			DOI    string `json:"doi"` // 論文に割り振られる世界共通の永久ID
		} `json:"result"`
	} `json:"resultList"`
}

type Article struct {
	ID    string
	Title string
	Link  string
}

func GetMuscleInfo() {
	ctx := context.Background()

	// 環境変数からDB接続文字列を取る
	// 例: user:pass@tcp(mysql:3306)/dbname?parseTime=true&charset=utf8mb4&loc=Asia%2FTokyo
	dsn := os.Getenv("MYSQL_DSN")
	if dsn == "" {
		panic("MYSQL_DSN is required")
	}

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if err := db.PingContext(ctx); err != nil {
		panic(err)
	}
	// 外部API（Europe PMC）から論文を取ってきて、Goで扱いやすい形に変換する関数
	articles, err := fetchLatestJapanStrengthTraining(ctx, 3)
	if err != nil {
		panic(err)
	}
	// DBに論文を保存する責務だけを持つ関数
	if err := upsertArticles(ctx, db, articles); err != nil {
		panic(err)
	}

	fmt.Printf("saved %d articles\n", len(articles))
}

// Europe PMCから "strength training" AND AFF:"Japan" を新着順で3件取る
func fetchLatestJapanStrengthTraining(ctx context.Context, limit int) ([]Article, error) {
	base := "https://www.ebi.ac.uk/europepmc/webservices/rest/search"

	query := `"strength training" AND AFF:"Japan"`
	u, _ := url.Parse(base)
	q := u.Query()
	q.Set("query", query)
	q.Set("format", "json")
	q.Set("sort_date", "y") // 新しい順
	q.Set("pageSize", fmt.Sprint(limit))
	u.RawQuery = q.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{Timeout: 20 * time.Second}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return nil, fmt.Errorf("epmc API status: %s", res.Status)
	}

	var parsed epmcResp
	if err := json.NewDecoder(res.Body).Decode(&parsed); err != nil {
		return nil, err
	}

	out := make([]Article, 0, len(parsed.ResultList.Result))
	for _, r := range parsed.ResultList.Result {
		if r.ID == "" || r.Source == "" || r.Title == "" {
			// どれか欠けてたらスキップ（必要ならログ）
			continue
		}
		extID := r.Source + ":" + r.ID

		link := ""
		if r.DOI != "" {
			link = "https://doi.org/" + r.DOI
		} else {
			// DOIがなければEurope PMCのページへ
			// source/id を使ったURLが安定
			link = "https://europepmc.org/article/" + r.Source + "/" + r.ID
		}

		out = append(out, Article{
			ID:    extID,
			Title: r.Title,
			Link:  link,
		})
	}

	if len(out) == 0 {
		return nil, errors.New("no results (query too strict or temporary issue)")
	}
	return out, nil
}

// PRIMARY KEY(id) で UPSERT
func upsertArticles(ctx context.Context, db *sql.DB, articles []Article) error {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := tx.PrepareContext(ctx, `
		INSERT INTO articles (id, title, link)
		VALUES (?, ?, ?)
		ON DUPLICATE KEY UPDATE
			title = VALUES(title),
			link  = VALUES(link)
	`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, a := range articles {
		if a.ID == "" || a.Title == "" || a.Link == "" {
			continue
		}
		if _, err := stmt.ExecContext(ctx, a.ID, a.Title, a.Link); err != nil {
			return err
		}
	}

	return tx.Commit()
}
