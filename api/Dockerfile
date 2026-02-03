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


# # ---------- build ----------
# FROM golang:1.23.4 AS builder
# WORKDIR /app

# COPY go.mod go.sum ./
# RUN go mod download

# COPY . .
# RUN CGO_ENABLED=0 GOOS=linux go build -o server .

# # ---------- run ----------
# FROM gcr.io/distroless/base-debian12
# WORKDIR /app
# COPY --from=builder /app/server /app/server

# EXPOSE 8080
# ENV PORT=8080

# CMD ["/app/server"]

