FROM golang:1.19.5-alpine as builder

WORKDIR /go/src/github.com/systemli/prometheus-jibri-exporter

ENV USER=appuser
ENV UID=10001

RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    "${USER}"

ADD . /go/src/github.com/systemli/prometheus-jibri-exporter
RUN go get -d -v && \
    go mod download && \
    go mod verify && \
    CGO_ENABLED=0 go build -ldflags="-w -s" -o /prometheus-jibri-exporter


FROM scratch

COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group
COPY --from=builder /prometheus-jibri-exporter /prometheus-jibri-exporter

USER appuser:appuser

EXPOSE 9888

ENTRYPOINT ["/prometheus-jibri-exporter"]
