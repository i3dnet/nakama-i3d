FROM heroiclabs/nakama-pluginbuilder:3.26.0 AS builder

ENV CGO_ENABLED 1
ENV GO111MODULE on
ENV GOPRIVATE "github.com/heroiclabs/nakama-project-template"

WORKDIR /plugins

COPY . .
RUN ls -halp

#RUN go get -u github.com/heroiclabs/nakama-common/runtime
RUN go mod download
RUN go mod tidy
RUN go build --trimpath --buildmode=plugin -o ./i3dplugin.so

FROM heroiclabs/nakama:3.26.0

COPY --from=builder /plugins/i3dplugin.so /nakama/data/modules/
COPY ./local.yml /nakama/data/

