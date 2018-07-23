package main

import (
	"net/http"

	"github.com/pankona/trigger"
	"google.golang.org/appengine"
)

func main() {
	http.HandleFunc("/trigger", trigger.Handler)
	appengine.Main()
}
