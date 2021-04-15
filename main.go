package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"text/template"
)

var (
	addr     = flag.String("web.listen-address", ":9889", "Address on which to expose metrics")
	jibriURL = flag.String("jibri-status-url", "http://localhost:2222/jibri/api/v1.0/health", "Jitsi jibri URL to scrape")
)

type jibriStats struct {
	Status struct {
		BusyStatus string `json:"busyStatus"`
		Health     struct {
			HealthStatus string   `json:"healthStatus"`
			Details      struct{} `json:"details"`
		} `json:"health"`
	} `json:"status"`
}

var tpl = template.Must(template.New("stats").Parse(`# HELP jibri_busystatus It check the status of the jibri whether BUSY, IDLE.
# TYPE jibri_busystatus gauge
jibri_busystatus {{.Status.BusyStatus}}
# HELP jibri_healthstatus It check the health status of the jibri whether HEALTHY or not.
# TYPE jibri_healthstatus gauge
jibri_healthstatus {{.Status.Health.HealthStatus}}`))

type handler struct {
	sourceURL string
}

func (h handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	resp, err := http.Get(h.sourceURL)
	if err != nil {
		log.Printf("scrape error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var stats jibriStats
	if err := json.NewDecoder(resp.Body).Decode(&stats); err != nil {
		log.Printf("json decoding error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/plain")
	_ = tpl.Execute(w, &stats)
}

func main() {
	log.SetFlags(0)
	flag.Parse()

	http.Handle("/metrics", handler{sourceURL: *jibriURL})
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal(err)
	}

	log.Println("Started Jitsi Jibri Metrics Exporter")
}
