FROM golang:1.24-alpine

WORKDIR /workspace

COPY go.work ./

COPY services/apis/bookapi/go.mod services/apis/bookapi/go.sum ./services/apis/bookapi/
COPY shared/go/configs/go.mod shared/go/configs/go.sum ./shared/go/configs/
COPY shared/go/errors/go.mod shared/go/errors/go.sum ./shared/go/errors/
COPY shared/go/logger/go.mod shared/go/logger/go.sum ./shared/go/logger/
COPY shared/go/utils/go.mod shared/go/utils/go.sum ./shared/go/utils/
COPY shared/go/models/go.mod shared/go/models/go.sum ./shared/go/models/
COPY shared/go/repositories/go.mod shared/go/repositories/go.sum ./shared/go/repositories/

# required to avoid err: cannot load module xxx listed in go.work file
COPY tools/db/dbmigrate/go.mod tools/db/dbmigrate/go.sum ./tools/db/dbmigrate/

RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    cd shared/go/configs && go mod download && \
    cd ../errors && go mod download && \
    cd ../logger && go mod download && \
    cd ../utils && go mod download && \
    cd ../models && go mod download && \
    cd ../repositories && go mod download && \
    cd ../../../services/apis/bookapi && go mod download

COPY services/apis/bookapi ./services/apis/bookapi
COPY shared/go/configs ./shared/go/configs
COPY shared/go/errors ./shared/go/errors
COPY shared/go/logger ./shared/go/logger
COPY shared/go/utils ./shared/go/utils
COPY shared/go/models ./shared/go/models
COPY shared/go/repositories ./shared/go/repositories

WORKDIR /workspace/services/apis/bookapi
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    go build -ldflags '-w -s' -a -o ./bin/app ./cmd/app

CMD ["/workspace/services/apis/bookapi/bin/app"]

EXPOSE 8080
