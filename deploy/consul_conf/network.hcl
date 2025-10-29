service {
  name = "network.api"
  id   = "network.api"
  address = "127.0.0.1"
  port = 9090
 
  check = {
    id = "network-health"
    name = "Check Network API health"
    http = "http://127.0.0.1:9090/health"
    method = "GET"
    header = {
      Content-Type = ["application/json"]
    }
    interval = "10s"
    timeout = "1s"
  }
}