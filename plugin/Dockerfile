FROM heroiclabs/nakama-pluginbuilder:3.26.0 AS builder

ENV CGO_ENABLED=1 GO111MODULE=on
WORKDIR /plugins
COPY go.mod go.sum ./
COPY config ./config
COPY internal ./internal
COPY *.go ./
COPY plugin/example/src ./plugin/example/src
WORKDIR /plugins/plugin/example/src
RUN go mod download
RUN go build -mod=readonly --trimpath --buildmode=plugin -o /plugins/i3dplugin.so ./cmd/main

FROM heroiclabs/nakama:3.26.0
COPY --from=builder /plugins/i3dplugin.so /nakama/data/modules/
COPY plugin/example/src/local.development.yml /nakama/data/local.yml
