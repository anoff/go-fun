// https://hackernoon.com/how-to-create-a-web-server-in-go-a064277287c9
package main
import (
  "net/http"
  "strings"
  "fmt"
)
func sayHello(w http.ResponseWriter, r *http.Request) {
  message := r.URL.Path
  message = strings.TrimPrefix(message, "/")
  message = "Hello " + message
  w.Write([]byte(message))
}
func main() {
  http.HandleFunc("/", sayHello)
  fmt.Println("Server started on http://localhost:8080")
  if err := http.ListenAndServe(":8080", nil); err != nil {
    panic(err)
  }
}