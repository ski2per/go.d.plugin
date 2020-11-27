<!--
title: "NGINX Vts monitoring"
custom_edit_url: https://github.com/netdata/go.d.plugin/edit/master/modules/nginx/README.md
sidebar_label: "NGINX VTS"
-->

# NGINX VTS monitoring with Netdata

[`NGINX`](https://www.nginx.com/) is a web server which can also be used as a reverse proxy, load balancer, mail proxy and HTTP cache. 

This module will monitor one or more `NGINX` servers, depending on your configuration.

## Requirements

 -   `NGINX` with configured [`nginx-module-vts`](https://github.com/vozlt/nginx-module-vts).

## Charts
- Main charts:
  - Nginx running times: 
    - Starting time (`loadMsec`)
    - Up time (`nowMsec`)
  - Nginx connections
    - Active connections (`active`)
		-	Reading connections, Name: "Reading"},
			writing", Name: "Writing"},
			waiting", Name: "Waiting"},
			accepted", Name: "Accepted"},
			handled", Name: "Handled"},
			requests", Name: "Requests"},

It produces following charts:
- TBC


## Configuration

Edit the `go.d/nginxvts.conf` configuration file using `edit-config` from the your agent's [config
directory](/docs/step-by-step/step-04.md#find-your-netdataconf-file), which is typically at `/etc/netdata`.

```bash
cd /etc/netdata # Replace this path with your Netdata config directory
sudo ./edit-config go.d/nginxvts.conf
```

Needs only `url` to server's `stub_status`. Here is an example for local and remote servers:

```yaml
jobs:
  - name: local
    url: http://192.168.66.6/status/format/json
  - name: remote
    url: http://8.8.8.8/status/format/json
```

For all available options please see module [configuration file](https://github.com/netdata/go.d.plugin/blob/master/config/go.d/nginxvts.conf).


## Troubleshooting

Check the module debug output. Run the following command as `netdata` user:

> ./go.d.plugin -d -m nginxvts
