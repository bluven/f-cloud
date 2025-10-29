service {
  name = "host.api"
  id   = "host.api"
  address = "127.0.0.1"
  port = 8080
 
  check = {
    id = "host-health"
    name = "Check Host API health"
    http = "http://127.0.0.1:8080/health"
    method = "GET"
    header = {
      Content-Type = ["application/json"]
    }
    interval = "10s"
    timeout = "1s"
  }
}
