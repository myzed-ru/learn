package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	type TeamCityBuildType struct {
		Id string `json:"id"`
	}

	type TeamCityBuildProperty struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	}

	type TeamCityBuildProperties struct {
		TCBuildProperty []TeamCityBuildProperty `json:"property"`
	}

	type TeamCityAPIRequest struct {
		TCBuildType       TeamCityBuildType       `json:"buildType"`
		TCBuildProperties TeamCityBuildProperties `json:"properties"`
	}

	tcRequest := &TeamCityAPIRequest{
		TCBuildType: TeamCityBuildType{
			Id: "123",
		},
		TCBuildProperties: TeamCityBuildProperties{
			TCBuildProperty: []TeamCityBuildProperty{
				{
					Name:  "name1",
					Value: "value1",
				},
			},
		},
	}

	out, err := json.MarshalIndent(tcRequest, "", "  ")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	fmt.Println(string(out))
}
