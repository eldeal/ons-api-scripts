package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// List represents an object containing a list of datasets
type DatasetList struct {
	Items      []Dataset `json:"items"`
	Count      int       `json:"count"`
	Offset     int       `json:"offset"`
	Limit      int       `json:"limit"`
	TotalCount int       `json:"total_count"`
}

type DatasetDetails struct {
	ID    string `json:"id,omitempty"`
	Links Links  `json:"links,omitempty"`
	State string `json:"state,omitempty"`
	Title string `json:"title,omitempty"`
}

type Dataset struct {
	Title string `json:"title"`
	Links Links  `json:"links"`
}

type Links struct {
	Dimensions    Link `json:"dimensions,omitempty"`
	CodeList      Link `json:"code_list,omitempty"`
	LatestVersion Link `json:"latest_version,omitempty"`
	Self          Link `json:"self,omitempty"`
}

type Link struct {
	Href string `json:"href,omitempty"`
	ID   string `json:"id,omitempty"`
}

type DimensionList struct {
	Items []Dimension `json:"items"`
}

type Dimension struct {
	Label string `json:"label"`
	Name  string `json:"name"`
	Links Links  `json:"links"`
}

// curl https://api.beta.ons.gov.uk/v1/datasets\?limit\=200 | jq '.items | .[] | .title'

func main() {
	list := getDatasets()

	if list == nil {
		log.Fatal("no list")
	}

	if len(list.Items) == 0 {
		log.Fatal("no items in list")
	}

	result := make(map[string][]Dimension) //title to list of dimension names
	for _, item := range list.Items {
		result[item.Title] = getDimensions(item.Links.LatestVersion.Href)
	}

	for title, dims := range result {
		fmt.Println(title)
		fmt.Println("  Contains: ")
		for _, d := range dims {
			fmt.Println(fmt.Sprintf("  - Label: \"%s\" CodeList: \"%s\"", d.Label, d.Links.CodeList.ID))
		}
	}
}

func getDimensions(latest string) []Dimension {
	res, err := http.Get(latest + "/dimensions?limit=200")
	if err != nil {
		log.Fatalf("cannot list datasets: [%s]", err.Error())
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var list *DimensionList
	json.Unmarshal(body, &list)

	return list.Items
}

func getDatasets() *DatasetList {
	resp, err := http.Get("https://api.beta.ons.gov.uk/v1/datasets?limit=200")
	if err != nil {
		log.Fatalf("cannot list datasets: [%s]", err.Error())
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body) // response body is []byte
	if err != nil {
		log.Fatal(err)
	}

	var l *DatasetList
	if err := json.Unmarshal(body, &l); err != nil {
		log.Fatal(err)
	}

	return l
}
