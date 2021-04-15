package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/go-cmp/cmp"
)

type constHandler struct {
	s string
}

func (h constHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(h.s))
}

func TestGetMetrics(t *testing.T) {
	tcs := []struct {
		statsJson string
		expected  string
	}{
		{
			statsJson: `{"status": {"busyStatus": "IDLE", "health": {"healthStatus": "HEALTHY"}}}`,
			expected: `# HELP jibri_busystatus It check the status of the jibri whether BUSY, IDLE.
# TYPE jibri_busystatus gauge
jibri_busystatus IDLE
# HELP jibri_healthstatus It check the health status of the jibri whether HEALTHY or not.
# TYPE jibri_healthstatus gauge
jibri_healthstatus HEALTHY`,
		},
		{
			statsJson: `{"status": {"busyStatus": "BUSY", "health": {"healthStatus": "HEALTHY"}}}`,
			expected: `# HELP jibri_busystatus It check the status of the jibri whether BUSY, IDLE.
# TYPE jibri_busystatus gauge
jibri_busystatus BUSY
# HELP jibri_healthstatus It check the health status of the jibri whether HEALTHY or not.
# TYPE jibri_healthstatus gauge
jibri_healthstatus HEALTHY`,
		},
		{
			statsJson: `{"status": {"busyStatus": "EXPIRED", "health": {"healthStatus": "UNHEALTHY"}}}`,
			expected: `# HELP jibri_busystatus It check the status of the jibri whether BUSY, IDLE.
# TYPE jibri_busystatus gauge
jibri_busystatus EXPIRED
# HELP jibri_healthstatus It check the health status of the jibri whether HEALTHY or not.
# TYPE jibri_healthstatus gauge
jibri_healthstatus UNHEALTHY`,
		},
	}

	for _, tc := range tcs {
		srv := httptest.NewServer(constHandler{tc.statsJson})

		h := handler{
			sourceURL: srv.URL,
		}
		req, err := http.NewRequest("GET", "/metrics", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		h.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}

		if rr.Body.String() != tc.expected {
			t.Logf("\n\nValue of body String is: %s ", rr.Body.String())
			t.Logf("\n\nValue of expected String is: %s ", tc.expected)
			t.Log("\n\n")
			t.Errorf("Response does not match the expected string:\n%s", cmp.Diff(rr.Body.String(), tc.expected))
		}

		srv.Close()
	}
}
