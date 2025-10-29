service {
  name = "storage.api"
  id   = "storage.api"
  address = "127.0.0.1"
  port = 3030
 
  check = {
    id = "storage-health"
    name = "Check Storage API health"
    http = "http://127.0.0.1:3030/health"
    method = "GET"
    header = {
      Content-Type = ["application/json"]
    }
    interval = "10s"
    timeout = "1s"
  }
}