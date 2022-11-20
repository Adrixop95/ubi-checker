# UBI checker
A very simple application whose purpose is to check when a container located on the RedHat Ecosystem Catalog (catalog.redhat.com) was last updated.

The output of the application allows it to be consumed by [Telegraf](https://github.com/influxdata/telegraf) within the [exec Input Plugin](https://github.com/influxdata/telegraf/blob/master/plugins/inputs/exec/README.md) and then be integrated into the Grafana/Prometheus stack.

The latest build can be downloaded from the GitHub Actions artifacts.