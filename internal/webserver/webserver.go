package webserver

import (
	"crypto/tls"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"strings"
	"text/template"
	"time"
)

func Webserver(port string) {
	mux := http.NewServeMux()

	mux.HandleFunc("/", serveHome)
	mux.HandleFunc("/events", streamLogs)

	http.ListenAndServe(port, mux)
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	var tmpl = template.Must(template.ParseFiles("./cmd/assets/index.html"))
	tmpl.Execute(w, nil)
}

func FetchOpensearch() string {
	req, err := http.NewRequest("GET", "https://opensearch-cluster-master.mon.tristan:9200/fluent-bit/_search", strings.NewReader(
		`{
		"query": {"match_all": {}},
		"sort": [{ "@timestamp": "desc" }],
		"size": 1
        }`))

	user := os.Getenv("VIMES_OPENSEARCH_USER")
	pw := os.Getenv("VIMES_OPENSEARCH_PASSWORD")
	slog.Error("check", "user", user, "pw", pw)
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

	slog.Error("Hmm what is happening here", "log", string(body))
	return string(body)
}

func streamLogs(w http.ResponseWriter, r *http.Request) {
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
