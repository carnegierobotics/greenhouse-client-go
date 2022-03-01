package main

import (
  "flag"
  "fmt"
  "os"
  "github.com/carnegierobotics/greenhouse-client-go/greenhouse"
)

func main() {
  var jobsUrl string
  var harvestUrl string
  var token string
  var onBehalfOf int
  flag.StringVar(&jobsUrl, "jobs-url", "https://boards-api.greenhouse.io", "Greenhouse Job Board API URL")
  flag.StringVar(&harvestUrl, "harvest-url", "https://harvest.greenhouse.io", "Greenhouse Harvest API URL")
  flag.IntVar(&onBehalfOf, "on-behalf-of", 0, "On-Behalf-Of user ID")
  flag.StringVar(&token, "token", "", "Greenhouse API token")
  flag.Parse()
  if token == "" {
    fmt.Printf("Please provide a token.")
    os.Exit(1)
  }
  if onBehalfOf == 0 {
    fmt.Printf("Please provide a On-Behalf-Of user.")
    os.Exit(1)
  }
  harvestClient := greenhouse.Client{BaseUrl: harvestUrl, Token: token, OnBehalfOf: onBehalfOf}
  harvestClient.BuildResty()
}
