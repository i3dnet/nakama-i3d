FROM heroiclabs/nakama-pluginbuilder:3.26.0 AS builder

ENV CGO_ENABLED 1
ENV GO111MODULE on
ENV GOPRIVATE "github.com/heroiclabs/nakama-project-template"

WORKDIR /plugins

COPY . .
RUN ls -halp

#RUN go get -u github.com/heroiclabs/nakama-common/runtime

WORKDIR /plugins/example/src
RUN ls -halp
RUN go mod download
RUN go mod tidy

WORKDIR /plugins
RUN go build -C ./example/src/cmd/main --trimpath --buildmode=plugin  -o /plugins/i3dplugin.so
RUN ls -halp
FROM heroiclabs/nakama:3.26.0

COPY --from=builder /plugins/i3dplugin.so /nakama/data/modules/
COPY example/src/local.development.yml /nakama/data/


