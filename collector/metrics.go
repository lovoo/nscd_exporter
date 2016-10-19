package collector

import "github.com/prometheus/client_golang/prometheus"

var (
	configMetrics = map[string]*prometheus.Desc{
		"server runtime": prometheus.NewDesc(
			"nscd_config_server_runtime",
			"server runtime",
			nil, nil,
		),
		"current number of threads": prometheus.NewDesc(
			"nscd_cur_threads",
			"current number of threads",
			nil, nil,
		),
		"maximum number of threads": prometheus.NewDesc(
			"nscd_max_threads",
			"server runtime",
			nil, nil,
		),
		"number of times clients had to wait": prometheus.NewDesc(
			"nscd_clients_wait_count",
			"number of times clients had to wait",
			nil, nil,
		),
		"paranoia mode enabled": prometheus.NewDesc(
			"nscd_paranoia_mode",
			"paranoia mode enabled",
			nil, nil,
		),
		"reload count": prometheus.NewDesc(
			"nscd_reload_count",
			"reload count",
			nil, nil,
		),
	}

	sectionMetrics = map[string]*prometheus.Desc{
		"cache is enabled": prometheus.NewDesc(
			"nscd_cache_enabled",
			"cache is enabled",
			[]string{"section"},
			nil,
		),
		"cache is persistent": prometheus.NewDesc(
			"nscd_cache_persistent",
			"cache is persistent",
			[]string{"section"},
			nil,
		),
		"cache is shared": prometheus.NewDesc(
			"nscd_cache_shared",
			"cache is shared",
			[]string{"section"},
			nil,
		),
		"suggested size": prometheus.NewDesc(
			"nscd_suggested_size",
			"suggested size",
			[]string{"section"},
			nil,
		),
		"total data pool size": prometheus.NewDesc(
			"nscd_total_data_pool_size",
			"total data pool size",
			[]string{"section"},
			nil,
		),
		"used data pool size": prometheus.NewDesc(
			"nscd_used_data_pool_size",
			"used data pool size",
			[]string{"section"},
			nil,
		),
		"seconds time to live for positive entries": prometheus.NewDesc(
			"nscd_time_ttl_positive_seconds",
			"seconds time to live for positive entries",
			[]string{"section"},
			nil,
		),
		"seconds time to live for negative entries": prometheus.NewDesc(
			"nscd_time_ttl_negative_seconds",
			"seconds time to live for negative entries",
			[]string{"section"},
			nil,
		),
		"cache hits on positive entries": prometheus.NewDesc(
			"nscd_cache_hits_positive",
			"cache hits on positive entries",
			[]string{"section"},
			nil,
		),
		"cache hits on negative entries": prometheus.NewDesc(
			"nscd_cache_hits_negative",
			"cache hits on negative entries",
			[]string{"section"},
			nil,
		),
		"cache misses on positive entries": prometheus.NewDesc(
			"nscd_cache_misses_positive",
			"cache misses on positive entries",
			[]string{"section"},
			nil,
		),
		"cache misses on negative entries": prometheus.NewDesc(
			"nscd_cache_misses_negative",
			"cache misses on negative entries",
			[]string{"section"},
			nil,
		),
		"cache hit rate": prometheus.NewDesc(
			"nscd_cache_hit_rate",
			"cache hit rate",
			[]string{"section"},
			nil,
		),
		"current number of cached values": prometheus.NewDesc(
			"nscd_cur_cached_values",
			"current number of cached values",
			[]string{"section"},
			nil,
		),
		"maximum number of cached values": prometheus.NewDesc(
			"nscd_max_cached_values",
			"maximum number of cached values",
			[]string{"section"},
			nil,
		),
		"maximum chain length searched": prometheus.NewDesc(
			"nscd_max_chain_length_searched",
			"maximum chain length searched",
			[]string{"section"},
			nil,
		),
		"number of delays on rdlock": prometheus.NewDesc(
			"nscd_delays_rdlock",
			"number of delays on rdlock",
			[]string{"section"},
			nil,
		),
		"number of delays on wrlock": prometheus.NewDesc(
			"nscd_delays_wrlock",
			"number of delays on wrlock",
			[]string{"section"},
			nil,
		),
		"memory allocations failed": prometheus.NewDesc(
			"nscd_mem_allocs_failed",
			"memory allocations failed",
			[]string{"section"},
			nil,
		),
		"check /etc/passwd for changes": prometheus.NewDesc(
			"nscd_check_etc_passwd_for_changes",
			"check /etc/passwd for changes",
			[]string{"section"},
			nil,
		),
	}
)
