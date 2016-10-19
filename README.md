# NSCD Exporter for Prometheus

[![GoDoc](https://godoc.org/github.com/lovoo/nscd_exporter?status.svg)](https://godoc.org/github.com/lovoo/nscd_exporter) [![Build Status](https://travis-ci.org/lovoo/nscd_exporter.svg?branch=master)](https://travis-ci.org/lovoo/nscd_exporter)

Exports statistics from NSCD (Name service caching daemon) and publishes them for scraping by Prometheus.

## Requirements

* ncsd

## Docker Usage

    docker run --privileged -d --name nscd_exporter -p 9289:9289 lovoo/nscd_exporter:latest

## Building

    make build

The resulting binary file will be placed under `build/nscd_exporter`.

## Running options

	Usage of build/nscd_exporter:
	  -log.format value
	    	If set use a syslog logger or JSON logging. Example: logger:syslog?appname=bob&local=7 or logger:stdout?json=true. Defaults to stderr.
	  -log.level value
	    	Only log messages with the given severity or above. Valid levels: [debug, info, warn, error, fatal].
	  -nscd.path string
	    	Path to nscd. (default "nscd")
	  -web.listen string
	    	Address on which to expose metrics and web interface. (default ":9119")
