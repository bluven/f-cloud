#!/bin/bash
export admin_key=edd1c9f034335f136f87ad84b625c8f1

function add_uc_route() {
  curl http://127.0.0.1:9180/apisix/admin/routes/uc-api -H "X-API-KEY: $admin_key" -X PUT -i -d '
  {
      "uri": "/uc/*",
      "upstream": {
          "service_name": "user.api",
          "type": "roundrobin",
          "discovery_type": "consul"
      },
      "plugins": {
        "proxy-rewrite": {
          "regex_uri": ["^/uc/(.*)", "/$1"]
        }
      }
  }'
}

function add_host_route() {
  curl http://127.0.0.1:9180/apisix/admin/routes/host-api -H "X-API-KEY: $admin_key" -X PUT -i -d '
  {
      "uri": "/host/*",
      "upstream": {
          "service_name": "host.api",
          "type": "roundrobin",
          "discovery_type": "consul"
      },
      "plugins": {
        "proxy-rewrite": {
          "regex_uri": ["^/host/(.*)", "/$1"]
        }
      }
  }'
}

function add_network_route() {
  curl http://127.0.0.1:9180/apisix/admin/routes/network-api -H "X-API-KEY: $admin_key" -X PUT -i -d '
  {
      "uri": "/network/*",
      "upstream": {
          "service_name": "network.api",
          "type": "roundrobin",
          "discovery_type": "consul"
      },
      "plugins": {
        "proxy-rewrite": {
          "regex_uri": ["^/network/(.*)", "/$1"]
        }
      }
  }'
}

function add_storage_route() {
  curl http://127.0.0.1:9180/apisix/admin/routes/storage-api -H "X-API-KEY: $admin_key" -X PUT -i -d '
  {
      "uri": "/storage/*",
      "upstream": {
          "service_name": "storage.api",
          "type": "roundrobin",
          "discovery_type": "consul"
      },
      "plugins": {
        "proxy-rewrite": {
          "regex_uri": ["^/storage/(.*)", "/$1"]
        }
      }
  }'
}

function add_instance_route() {
  curl http://127.0.0.1:9180/apisix/admin/routes/instance-api -H "X-API-KEY: $admin_key" -X PUT -i -d '
  {
      "uri": "/instance/*",
      "upstream": {
          "service_name": "instance.api",
          "type": "roundrobin",
          "discovery_type": "consul"
      },
      "plugins": {
        "proxy-rewrite": {
          "regex_uri": ["^/instance/(.*)", "/$1"]
        }
      }
  }'
}

add_uc_route
add_host_route
add_network_route
add_storage_route
add_instance_route
