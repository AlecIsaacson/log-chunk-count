package main

import (
	//"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
)

type nrInsightsJSONStruct struct {
	Results []struct {
		Count int `json:"count"`
	} `json:"results"`
	// PerformanceStats struct {
	// 	InspectedCount int `json:"inspectedCount"`
	// 	OmittedCount   int `json:"omittedCount"`
	// 	MatchCount     int `json:"matchCount"`
	// 	WallClockTime  int `json:"wallClockTime"`
	// } `json:"performanceStats"`
	// Metadata struct {
	// 	EventTypes      []string      `json:"eventTypes"`
	// 	EventType       string        `json:"eventType"`
	// 	OpenEnded       bool          `json:"openEnded"`
	// 	BeginTime       time.Time     `json:"beginTime"`
	// 	EndTime         time.Time     `json:"endTime"`
	// 	BeginTimeMillis int64         `json:"beginTimeMillis"`
	// 	EndTimeMillis   int64         `json:"endTimeMillis"`
	// 	RawSince        string        `json:"rawSince"`
	// 	RawUntil        string        `json:"rawUntil"`
	// 	RawCompareWith  string        `json:"rawCompareWith"`
	// 	GUID            string        `json:"guid"`
	// 	RouterGUID      string        `json:"routerGuid"`
	// 	Messages        []interface{} `json:"messages"`
	// 	Contents        []struct {
	// 		Function string `json:"function"`
	// 		Simple   bool   `json:"simple"`
	// 	} `json:"contents"`
	// } `json:"metadata"`
}

func runInsightsQuery(urlToGet string, nrAPI string, logVerbose bool) ([]byte) {
	if logVerbose {
		fmt.Println("In runInsightsQuery")
		fmt.Println(urlToGet, nrAPI)
	}
	client := &http.Client{}
	req, err := http.NewRequest("GET", urlToGet, nil)
	req.Header.Set("X-Query-Key", nrAPI)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)

	if err != nil || resp.StatusCode != 200 {
		fmt.Println("New Relic error")
		fmt.Println(resp)
	}
	defer resp.Body.Close()
	response, err := ioutil.ReadAll(resp.Body)

	if logVerbose {
		fmt.Println("New Relic Response:", resp.Status)
		fmt.Println("End of runInsightsQuery")
	}
	return response
}

func main() {
	//Get the command line arguments
	fmt.Println("Log Chunk Count v1.0")
	nrAPI := flag.String("apikey", "", "New Relic admin user API Key")
	nrAccount := flag.String("account", "", "New Relic account number")
	chunkCount := flag.Int("chunks", 20, "Maximum number of chunks to query for")
	logVerbose := flag.Bool("verbose", false, "Writes verbose logs for debugging")
	flag.Parse()

	if *logVerbose {
		fmt.Println("Verbose logging enabled.")
	}

	fmt.Println("Account:",*nrAccount)
	fmt.Println("Max chunks:",*chunkCount)
	fmt.Println("Since 24 hours ago")
	fmt.Println("Chunk Number,Count")

	for chunk := 1; chunk < *chunkCount+1; chunk++ {
		nrBaseURL := fmt.Sprintf("https://insights-api.newrelic.com/v1/accounts/%v/query?nrql=SELECT%%20count(*)%%20FROM%%20Log%%20WHERE%%20%%60message-%.2d%%60%%20IS%%20NOT%%20NULL%%20SINCE%%2024%%20HOURS%%20AGO", *nrAccount, chunk)
		//fmt.Println("URL:", nrBaseURL)
		//fmt.Println("API:", nrAPI)
		nrQueryJSON := runInsightsQuery(nrBaseURL, *nrAPI, *logVerbose)

		if *logVerbose {
			fmt.Printf("Monitor list JSON:\n%s\n", nrQueryJSON)
		}

		//Unmarshal the monitors list into a struct
		if *logVerbose {
			fmt.Println("Unmarshalling monitors into struct")
		}
		var nrQueryResult nrInsightsJSONStruct
		if err := json.Unmarshal(nrQueryJSON, &nrQueryResult); err != nil {
			panic(err)
		}

		for _, result := range nrQueryResult.Results {
			fmt.Printf("%v,%v\n", chunk, result.Count)
		}
	}
}
