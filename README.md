# Jibri Metrics Exporter

[![Release](https://img.shields.io/github/v/release/prayagsingh/prometheus-jibri-exporter?color=dark-green)](https://github.com/prayagsingh/prometheus-jibri-exporter/releases/)
[![Integration](https://github.com/prayagsingh/prometheus-jibri-exporter/workflows/Integration/badge.svg?branch=main)](https://github.com/prayagsingh/prometheus-jibri-exporter/workflows/Integration/badge.svg?branch=main) [![Quality](https://github.com/prayagsingh/prometheus-jibri-exporter/workflows/Quality/badge.svg?branch=main)](https://github.com/prayagsingh/prometheus-jibri-exporter/workflows/Quality/badge.svg?branch=main) [![Docker Image Version](https://img.shields.io/docker/v/prayagsingh/prometheus-jibri-exporter/latest)](https://hub.docker.com/r/prayagsingh/prometheus-jibri-exporter) [![Docker Image Size (latest semver)](https://img.shields.io/docker/image-size/prayagsingh/prometheus-jibri-exporter)](https://hub.docker.com/r/prayagsingh/prometheus-jibri-exporter)

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
docker run -p 9889:9889 --network <jitsi_network> prayagsingh/prometheus-jibri-exporter:latest -jibri-status-url http://<jibri_container_name OR jibri_container_ip>:2222/jibri/api/v1.0/health
```

## Metrics

**Please select the tag as per your requirement. The output of v1.2.0 and v1.3.0 is different**

**v1.2.0**

```
# HELP jibri_busystatus It check the status of the jibri whether BUSY, IDLE.
# TYPE jibri_busystatus gauge
jibri_busystatus IDLE
# HELP jibri_healthstatus It check the health status of the jibri whether HEALTHY or not.
# TYPE jibri_healthstatus gauge
jibri_healthstatus HEALTHY
```

**v1.3.0** Here IDLE is 0, HEALTHY is 1, BUSY is 1 and UNHEALTHY is 0

```
# HELP jibri_busystatus It check the status of the jibri whether BUSY, IDLE.
# TYPE jibri_busystatus gauge
jibri_busystatus 0
# HELP jibri_healthstatus It check the health status of the jibri whether HEALTHY or not.
# TYPE jibri_healthstatus gauge
jibri_healthstatus 1
```


## License

GPLv3
