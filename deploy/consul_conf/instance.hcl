service {
  name = "instance.api"
  id   = "instance.api"
  address = "127.0.0.1"
  port = 5050
 
  check = {
    id = "instance-health"
    name = "Check Instance API health"
    http = "http://127.0.0.1:5050/health"
    method = "GET"
    header = {
      Content-Type = ["application/json"]
    }
    interval = "10s"
    timeout = "1s"
  }
}