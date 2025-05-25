FROM golang:alpine

WORKDIR /workspace

COPY go.work ./
COPY services/apis/bookapi ./services/apis/bookapi

WORKDIR /workspace/services/apis/bookapi

RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    go mod download && go build -o ./bin/api ./cmd/api

CMD ["/workspace/services/apis/bookapi/bin/api"]

EXPOSE 8080
