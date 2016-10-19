package collector

import (
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

	g := prometheus.NewGauge(prometheus.GaugeOpts{Help: "Dummy metric to make the prometheus library happy. It is not used anywhere.", Name: "nscd_dummy_metric"})
	g.Describe(ch)
}

func (e *Exporter) Collect(ch chan<- prometheus.Metric) {
	data, err := Scrape(e.NSCDPath)
	if err != nil {
		log.Errorf("could not scrape from nscd: %v", err)
		return
	}
	for section, d := range data {
		if section == "nscd configuration" {
			for _, line := range d {
				m, ok := configMetrics[line.desc]
				if !ok {
					continue
				}
				ch <- prometheus.MustNewConstMetric(m, prometheus.GaugeValue, line.val)
			}
			continue
		}
		for _, line := range d {
			m, ok := sectionMetrics[line.desc]
			if !ok {
				continue
			}
			ch <- prometheus.MustNewConstMetric(m, prometheus.GaugeValue, line.val, section)
		}
	}
}
