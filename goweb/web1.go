//basics of web application with go
package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html") //setting the content-type header
	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1>welcome to my awesome site</h1>")
	} else if r.URL.Path == "/contact" {
		fmt.Fprint(w, "To get in touch,please send an email to <a href=\"mailto:support@lenslocked.com\">support@lenslocked.com</a>.")
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "<h1> we could not find the page you were looking for:(</h1><p>please email us if you keep being sent to an invalid page.</p>")
	}
}
func main() {
	mux := &http.ServeMux{}
	mux.HandleFunc("/contact", handlerFunc)
	//http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":3000", mux)
}
