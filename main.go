package main
import (
    "fmt"
    "net/http"
)
func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello Hikari")
}
func main() {
    http.HandleFunc("/World", handler)
    http.ListenAndServe(":8080", nil)
}