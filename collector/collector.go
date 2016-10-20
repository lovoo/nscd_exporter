package collector

import (
	"strings"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/log"
)

type Exporter struct {
	NSCDPath string
}

func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {
	_, err := Scrape(e.NSCDPath)
	if err != nil {
		log.Fatalf("could not collect data from nscd command: %v", err)
	}

	prometheus.NewGauge(prometheus.GaugeOpts{Help: "Dummy metric to make the prometheus library happy. It is not used anywhere.", Name: "nscd_dummy_metric"}).Describe(ch)
}

func (e *Exporter) Collect(ch chan<- prometheus.Metric) {
	data, err := Scrape(e.NSCDPath)
	if err != nil {
		log.Errorf("could not scrape from nscd: %v", err)
		return
	}
	for section, d := range data {
		var (
			m        *prometheus.Desc
			ok       bool
			pushFunc func()
		)

		for _, line := range d {
			if section == "nscd configuration" {
				m, ok = configMetrics[line.desc]
				pushFunc = func() {
					ch <- prometheus.MustNewConstMetric(m, prometheus.GaugeValue, line.val)
				}
			} else {
				m, ok = sectionMetrics[line.desc]
				pushFunc = func() {
					ch <- prometheus.MustNewConstMetric(m, prometheus.GaugeValue, line.val, strings.TrimSpace(strings.TrimSuffix(section, "cache")))
				}
			}
			if !ok {
				continue
			}
			pushFunc()
		}

	}
}
