// REFERENCE (Learn Golang in 30 Minutes): https://topherpedersen.blog/2019/12/23/learn-golang-in-15-minutes/
// REFERENCE (html/template package): https://golang.org/pkg/html/template/
// REFERENCE (Writing Web Applications in Go): https://golang.org/doc/articles/wiki/

package main

/*
import(
	"fmt"
	"io/ioutil"
	"net/http"
	"html/template"
)
*/

/*
type WebPage struct {
	H
}

func main() {
	fmt.Printf("hello, golang\n")
}
*/

import (
    "fmt"
    "log"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}