package main

import (
  "flag"
  "fmt"
  "os"
  "github.com/carnegierobotics/greenhouse-client-go/greenhouse"
  "github.com/carnegierobotics/greenhouse-client-go/internal/http"
  "github.com/carnegierobotics/greenhouse-client-go/internal/utils"
)

func main() {
  var jobsUrl string
  var harvestUrl string
  var token string
  var onBehalfOf string
  flag.StringVar(&jobsUrl, "jobs-url", "https://boards-api.greenhouse.io", "Greenhouse Job Board API URL")
  flag.StringVar(&harvestUrl, "harvest-url", "https://harvest.greenhouse.io", "Greenhouse Harvest API URL")
  flag.StringVar(&onBehalfOf, "on-behalf-of", "", "On-Behalf-Of user")
  flag.StringVar(&token, "token", "", "Greenhouse API token")
  flag.Parse()
  if token == "" {
    fmt.Printf("Please provide a token.")
    os.Exit(1)
  }
  if onBehalfOf == "" {
    fmt.Printf("Please provide a On-Behalf-Of user.")
    os.Exit(1)
  }
  harvestClient := http.Client{BaseUrl: harvestUrl, Token: token, OnBehalfOf: onBehalfOf}
  harvestClient.BuildResty()
  // Get all users
  var allUsers []*greenhouse.User
  err := utils.GetAll(harvestClient, "users", &allUsers)
  if err != nil {
    fmt.Printf("An error occurred", err.Error())
    os.Exit(1)
  }
  for _, u := range allUsers {
    fmt.Printf("%s\n", u.Name)
  }
  // Get a single user by ID
  var user *greenhouse.User
  err = utils.GetById(harvestClient, "users", "4005448005", &user)
  if err != nil {
    fmt.Printf("An error occurred", err.Error())
    os.Exit(1)
  }
  fmt.Printf("%v", user.Name)
}
