# Docker StatsD

- [x] basic setup with Docker Compose
  - able to run the docker image with Graphite + StatsD
  - able to call StatsD server from the application
- [ ] basic setup with Kubernetes(Docker Desktop)
- [x] sample application with golang
  - sample code to send metrics to the server
- [ ] sample application with nodejd
- [ ] metrics naming convention best practices
- [ ] actual code in large codebase
- [x] exporting data as CSV
  - client render API to download the data as CSV from different chart
- [ ] how to periodically export data from Graphite
- [ ] how to create visualization (graphs) from the metrics gathered
- [ ] how to handle different type of metrics collection
- [ ] how to create dashboard for different services
- [ ] how to add alerting mechanism for graphite
- [ ] how to create actionable action from the metrics gathered

## Exporting as CSV

Just go the the graph you want, and change the url to:

```bash
/render?target=...&format=csv
```

## References

- https://sysdig.com/blog/monitoring-statsd-metrics/
- https://hub.docker.com/r/graphiteapp/docker-graphite-statsd
- https://github.com/quipo/statsd

