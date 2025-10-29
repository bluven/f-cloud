service {
  name = "user.api"
  id   = "user.api"
  address = "127.0.0.1"
  port = 8787
 
  check = {
    id = "uc-health"
    name = "Check User Center API health"
    http = "http://127.0.0.1:8787/health"
    method = "GET"
    header = {
      Content-Type = ["application/json"]
    }
    interval = "10s"
    timeout = "1s"
  }
}
