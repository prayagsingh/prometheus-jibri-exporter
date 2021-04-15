# Jibri Metrics Exporter

[![Integration](https://github.com/prayagsingh/prometheus-jibri-exporter/workflows/Integration/badge.svg?branch=master)](https://github.com/prayagsingh/prometheus-jibri-exporter/workflows/Integration/badge.svg?branch=master) [![Quality](https://github.com/prayagsingh/prometheus-jibri-exporter/workflows/Quality/badge.svg?branch=master)](https://github.com/prayagsingh/prometheus-jibri-exporter/workflows/Quality/badge.svg?branch=master) [![Docker Cloud Automated build](https://img.shields.io/docker/cloud/automated/prayagsingh/prometheus-jibri-exporter)](https://hub.docker.com/r/prayagsingh/prometheus-jibri-exporter) [![Docker Image Size (latest semver)](https://img.shields.io/docker/image-size/prayagsingh/prometheus-jibri-exporter)](https://hub.docker.com/r/prayagsingh/prometheus-jibri-exporter)

Prometheus Exporter for Jibri written in Go. Special thanks to `Systemli` team since this project is based on [Prometheus-jitsi-meet-exporter](https://github.com/systemli/prometheus-jitsi-meet-exporter) repository.

There's only one GET [endpoint to check the status of jibri](https://github.com/jitsi/jibri/blob/master/doc/http_api.md#url) (like /jibri/api/v1.0/health); you can configure the used URL with the `jibri-status-url`.
The exporter will handle it.

## Usage

```
go get github.com/prayagsingh/prometheus-jibri-exporter
go install github.com/prayagsingh/prometheus-jibri-exporter
$GOPATH/bin/prometheus-jibri-exporter
```

### Docker

```
docker run -p 9889:9889 prayagsingh/prometheus-jibri-exporter:latest -jibri-status-url http://localhost:2222/jibri/api/v1.0/health
```

## Metrics

```
# HELP jibri_busystatus It check the status of the jibri whether BUSY, IDLE.
# TYPE jibri_busystatus gauge
jibri_busystatus IDLE
# HELP jibri_healthstatus It check the health status of the jibri whether HEALTHY or not.
# TYPE jibri_healthstatus gauge
jibri_healthstatus HEALTHY`
```

## License

GPLv3
