// REFERENCE (Simple Static Web Server in Go): https://stackoverflow.com/questions/26559557/how-do-you-serve-a-static-html-file-using-a-go-web-server
package main

import (
        "net/http"
)

func main() {
        http.Handle("/", http.FileServer(http.Dir("./")))
        // http.Handle("/myroute", http.FileServer(http.Dir("./myroute")))
        // http.Handle("/my-other-route", http.FileServer(http.Dir("./my-other-route")))
        http.ListenAndServe(":3000", nil)
}