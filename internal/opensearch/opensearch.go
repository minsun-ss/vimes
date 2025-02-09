package opensearch

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"strings"
	"time"
)

// fetches opensearch to grab logs. Configuration of which logs to fetch and dump to
// clickhouse determined entirely by the configmap
func FetchOpensearch() map[string]interface{} {
	req, err := http.NewRequest("GET", "https://opensearch-cluster-master.mon.tristan:9200/fluent-bit/_search", strings.NewReader(
		`{
		"query": {"match_all": {}},
		"sort": [{ "@timestamp": "desc" }],
		"size": 1
        }`))

	user := os.Getenv("VIMES_OPENSEARCH_USER")
	pw := os.Getenv("VIMES_OPENSEARCH_PASSWORD")
	req.SetBasicAuth(user, pw)
	req.Header.Set("Content-Type", "application/json")
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	if err != nil {
		slog.Error("Shit")
	}
	resp, err := client.Do(req)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		slog.Error("shit")
	}

	var m map[string]interface{}
	// var m2 map[string]string
	json.Unmarshal(body, &m)

	for k, v := range m {
		slog.Error("kv", "key", k, "value", v)
	}

	m2 := m["hits"].(map[string]interface{})["hits"].([]interface{})[0].(map[string]interface{})
	for k, v := range m2 {
		slog.Error("2kv", "key", k, "value", v)
	}
	if v, exists := m2["kubernetes"]; exists {
		slog.Error("whatever", "value", v)
		return m["hits"].(map[string]interface{})

	} else {
		return m["hits"].(map[string]interface{})
	}

	// slog.Error("Hmm what is happening here", "log", m["hits"])
}

func StreamLogs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Create channel for client disconnect detection
	notify := w.(http.CloseNotifier).CloseNotify()

	// Send update every 2 seconds
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-notify:
			// Client disconnected
			return
		case <-ticker.C:
			// Send update

			resp := FetchOpensearch()
			msg := fmt.Sprintf("Update #%s: %s", resp, time.Now().Format(time.RFC3339))
			fmt.Fprintf(w, "data: %s\n\n", msg)
			w.(http.Flusher).Flush()
		}
	}
}
