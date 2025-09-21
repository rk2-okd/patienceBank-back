
FROM golang:1.21
WORKDIR /api
COPY go.mod .
COPY go.sum .
RUN go mod download
RUN go mod tidy
COPY . .
RUN go build -o main .
CMD [ "./main" ]