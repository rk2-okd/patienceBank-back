FROM golang:1.23.4
WORKDIR /api

COPY ./go.mod .
COPY ./go.sum .

# 依存取得と整理
RUN go mod download
# ソースコードをコピー
COPY . .

# ビルド
RUN go build -o main .

# 実行
CMD ["./main"]