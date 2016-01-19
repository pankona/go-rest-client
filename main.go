package main

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func getEnvVar(varName string) (result string) {
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		if pair[0] == varName {
			return pair[1]
		}
	}
	return ""
}

func main() {

	circleci_url := "https://circleci.com/api/v1/project/pankona/gomo-simra/tree/master"
	client := &http.Client{}
	query := url.Values{"circle-token": {getEnvVar("CIRCLECI_API_KEY")}}

	req, _ := http.NewRequest(
		"POST",
		circleci_url+"?"+query.Encode(),
		nil,
	)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	println(string(body))

}
