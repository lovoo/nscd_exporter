package main

import (
	"flag"
	"net/http"

	"github.com/lovoo/nscd_exporter/collector"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/log"
)

var (
	listenAddress = flag.String("web.listen", ":9119", "Address on which to expose metrics and web interface.")
	nscdPath      = flag.String("nscd.path", "nscd", "Path to nscd.")
)

func main() {
	flag.Parse()
	prometheus.MustRegister(&collector.Exporter{NSCDPath: *nscdPath})

	handler := prometheus.Handler()
	http.Handle("/metrics", handler)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html>
			<head><title>NSCD Exporter</title></head>
			<body>
			<h1>NSCD Exporter</h1>
			<p><a href="/metrics">Metrics</a></p>
			</body>
			</html>`))
	})
	log.Infof("Starting HTTP Server on %v...", *listenAddress)
	err := http.ListenAndServe(*listenAddress, nil)
	if err != nil {
		log.Fatalf("Could not start http server: %v", err)
	}
}
