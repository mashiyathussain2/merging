package main

import "net/http"

// Write an http web server that returns an html header that shows the user's IP address and the time.
// Usage:
// go run examples/web/httpserverwithhtml.go
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<h1>Hello, world!</h1>"))
	})
	http.ListenAndServe(":8080", nil)
}
