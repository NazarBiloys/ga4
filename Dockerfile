FROM golang:1.18-buster as build

RUN apt-get update && apt-get install -y ca-certificates

WORKDIR src

ADD go.mod go.sum ./
RUN go mod download

ADD . .

RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH go build -o /bin/push-ratio-app ./cmd/ga4/

FROM alpine:3.16.0

WORKDIR app

COPY --from=build /bin/push-ratio-app /app/push-ratio-app

COPY push-ratio /etc/cron.d/push-ratio

RUN chmod 744 /etc/cron.d/push-ratio

COPY . /app

RUN chmod 755 /app/start_push_event.sh

RUN crontab /etc/cron.d/push-ratio

ENTRYPOINT ["crond", "-f"]