FROM golang:1.24.1 AS builder

WORKDIR /app

COPY ./src .

RUN go mod download
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -C ./cmd/main  -o /app/mock_api_server

RUN ls -la /app

FROM alpine:3.20.3

ARG SERVICE_USER=usr_i3d
ARG SERVICE_GROUP=grp_i3d

RUN apk --no-cache add ca-certificates shadow
RUN addgroup -S ${SERVICE_GROUP} && adduser -S ${SERVICE_USER} -G ${SERVICE_GROUP}

EXPOSE 8181
WORKDIR /app

COPY --from=builder /app/mock_api_server .
RUN chmod +x /app/mock_api_server
RUN chown -R ${SERVICE_USER}:${SERVICE_GROUP} /app

USER ${SERVICE_USER}
CMD  [ "/app/mock_api_server" ]