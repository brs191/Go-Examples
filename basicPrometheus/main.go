package main

import "basicPrometheus/prometheus"

/*
1. Install Prometheus server in the local system
https://computingforgeeks.com/install-prometheus-server-on-debian-ubuntu-linux/

2. Create a server/service to listen at a port ex: 2112;; This is a prometheus target/endpoint

3. Create a Scrape configuration for the new server/service with the port ex: 2112

4. Install Grafana as https://prometheus.io/docs/visualization/grafana/

5. Link prometheus datastore @default prometheus 9090

6. Create a dashboard and look for the metrics in either default prometheus or our endpoint/service
 */
func main() {
	prometheus.Init()
	println("basicPrometheus done!!")
}
