package main

import (
	"flag"
	"fmt"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"os"
)

func main() {
	var jobsUrl string
  var jobsToken string
	var harvestUrl string
	var harvestToken string
	var onBehalfOf int
	flag.StringVar(&jobsUrl, "jobs-url", "https://boards-api.greenhouse.io", "Greenhouse Job Board API URL")
	flag.StringVar(&harvestUrl, "harvest-url", "https://harvest.greenhouse.io", "Greenhouse Harvest API URL")
	flag.IntVar(&onBehalfOf, "on-behalf-of", 0, "On-Behalf-Of user ID")
	flag.StringVar(&harvestToken, "harvest-token", "", "Greenhouse Harvest API token")
  flag.StringVar(&jobsToken, "jobs-token", "", "Greenhouse Job Board API token")
	flag.Parse()
	if harvestToken == "" {
		fmt.Printf("Please provide a token.")
		os.Exit(1)
	}
	if onBehalfOf == 0 {
		fmt.Printf("Please provide a On-Behalf-Of user.")
		os.Exit(1)
	}
	harvestClient := greenhouse.Client{BaseUrl: harvestUrl, Token: harvestToken, OnBehalfOf: onBehalfOf}
	harvestClient.BuildResty()
  /*
	obj, err := greenhouse.GetJob(&harvestClient, 4003423005)
  if err != nil {
    fmt.Printf(err.Error())
  }
	fmt.Printf("%+v\n", obj)
  */
  body, err := greenhouse.GetAllRejectionReasons(&harvestClient, true, 20)
  if err != nil {
    fmt.Printf(err.Error())
    os.Exit(1)
  }
  fmt.Printf("%+v", body)
}
