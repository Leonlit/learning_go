package main

import (
	"fmt"
	"gulnManagement/webScanner/headers"
	"gulnManagement/webScanner/internal/loader"
	"gulnManagement/webScanner/internal/models"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func initLog() {
	file, err := os.OpenFile("scanner.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}

	log.SetOutput(file)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func main() {
	USER_AGENT := "Mozilla/5.0 (Linux; Android 16; Pixel 9) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/143.0.12.45 Mobile Safari/537.36"
	URL := "https://en.wikipedia.org/wiki/Body_armor#Middle_Ages"

	initLog()

	c := &http.Client{
		Timeout: 10 * time.Second,
	}

	req, err := http.NewRequest("GET", URL, nil)

	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header.Set("User-Agent", USER_AGENT)

	res, err := c.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	if res.StatusCode != http.StatusOK {
		log.Printf("Error: Got %d status code for %s", res.StatusCode, URL)
	}

	//resHeaders := res.Header
	if err != nil {
		fmt.Println(err)
		return
	}

	//TODO: need to support multiple folders
	folderPath := "./rules/headers/"
	rules, err := loader.LoadRulesFromYAML(folderPath)

	//fmt.Println(rules)

	if err != nil {
		log.Println("Error loading rules, ", err)
		return
	}

	rulesByHeader := map[string]models.Rule{}

	for _, r := range rules {
		rulesByHeader[strings.ToLower(r.Header)] = r
	}

	//fmt.Println(rulesByHeader)

	/* 	for key, values := range resHeaders {
	//fmt.Println(strings.ToLower(key), values)
	*/
	rule, hasRule := rulesByHeader[strings.ToLower("strict-transport-security")]
	//fmt.Println(rule, hasRule)
	if !hasRule {
		return
	}

	values := []string{"max-age=384710; includeSubDomains; preload"}

	findings := headers.CheckHeader(rule, values)
	for _, f := range findings {
		fmt.Println(f)
	}
	/* 	} */
}
