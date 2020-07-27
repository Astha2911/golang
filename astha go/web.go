package main

import (
	"fmt"
	"net/http"
)

func index_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `<h1> hey there</h1>
		<p>go is fast</p>
		<p>go is an open source programming language that make it easy to built simple,reliable
		and efficient software. Go is procedural programming language and 
		static type it meanse compile type language</p>`)
}
func about_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "this is go web designing")

}
func main() {
	http.HandleFunc("/", index_handler)
	http.HandleFunc("/about/", index_handler)
	http.ListenAndServe(":8000", nil)
}
