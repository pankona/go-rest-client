package main

import (
	"net/http"

	"github.com/pankona/gomo-simra-daily-build-trigger/trigger"
	"google.golang.org/appengine"
)

func main() {
	http.HandleFunc("/trigger", trigger.Handler)
	appengine.Main()
}
