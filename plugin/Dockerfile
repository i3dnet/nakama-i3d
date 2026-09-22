FROM heroiclabs/nakama-pluginbuilder:3.41.0@sha256:eb506121ec2e67f39febee3a4252ae05598a0b39d3517b713ec1577714a9e5b5 AS builder

ENV CGO_ENABLED=1 GO111MODULE=on
WORKDIR /plugins
COPY go.mod go.sum ./
COPY config ./config
COPY internal ./internal
COPY *.go ./
COPY plugin/example/src ./plugin/example/src
WORKDIR /plugins/plugin/example/src
RUN go mod download
ARG PLUGIN_BUILD_TAGS=""
RUN go build -tags="$PLUGIN_BUILD_TAGS" -mod=readonly --trimpath --buildmode=plugin -o /plugins/i3dplugin.so ./cmd/main

FROM heroiclabs/nakama:3.41.0@sha256:ca9f3fe65c90f56860fe5d7cd024868a395cb1a001abb9912bf87ef47d605349
COPY --from=builder /plugins/i3dplugin.so /nakama/data/modules/
COPY plugin/example/src/local.development.yml /nakama/data/local.yml
