# Notification Hub
A Go application server to receive health check information from the status monitoring [deamon](https://github.com/dumindux/application-status-monitoring-deamon) using websockets and saves them in InfluxDB
<br><br>
This application acting as a websockets server has the ability to connect to multiple status monitoring [deamons](https://github.com/dumindux/application-status-monitoring-deamon) and collect metrics over time.<br>
Collected metrics are written to the InfluxDB instance specified in the configuration.<br>
This status data can be visualized using [Grafana](https://grafana.com/) or [Chronograf](https://www.influxdata.com/time-series-platform/chronograf/)

