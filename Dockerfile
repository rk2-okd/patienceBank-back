# ========================
# 1. ビルドフェーズ
# ========================
FROM golang:1.23-alpine AS builder

# 必要なパッケージ（Cライブラリ用）を追加
RUN apk add --no-cache gcc musl-dev

WORKDIR /app

# 依存関係取得
COPY go.mod go.sum ./
RUN go mod download

# アプリ全体をコピー
COPY . .

# Goアプリをビルド（main.goを含めた構成前提）
RUN go build -o app .

# ========================
# 2. 実行フェーズ
# ========================
FROM alpine:3.20

WORKDIR /app

# 実行バイナリだけコピー
COPY --from=builder /app/app .

# 必要に応じてポートを公開（例：8080）
EXPOSE 8080

# 起動コマンド
CMD ["./app"]