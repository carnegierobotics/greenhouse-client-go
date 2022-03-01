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
  harvestClient := http.Client{BaseUrl: harvestUrl, Token: token, OnBehalfOf: onBehalfOf}
  harvestClient.BuildResty()
  // Get all users
  var allUsers []*greenhouse.User
  err := utils.GetAll(&harvestClient, "users", &allUsers)
  if err != nil {
    fmt.Printf("An error occurred", err.Error())
    os.Exit(1)
  }
  /*
  for _, u := range allUsers {
    fmt.Printf("%s\n", u.Name)
  }
  */
  // Get a single user by ID
  var user *greenhouse.User
  err = utils.GetById(&harvestClient, "users", 4005448005, &user)
  if err != nil {
    fmt.Printf("An error occurred", err.Error())
    os.Exit(1)
  }
  // fmt.Printf("%v", user.Name)
  // Get all departments
  var allDepartments []*greenhouse.Department
  err = utils.GetAll(&harvestClient, "departments", &allDepartments)
  if err != nil {
    fmt.Printf(err.Error())
    os.Exit(1)
  }
  /*
  for _, d := range allDepartments {
    fmt.Printf("%s\n", d.Name)
  }
  */
  // Create a department
  /*
  testDept := greenhouse.DepartmentBasics{Name: "test"}
  err = greenhouse.CreateDepartment(&harvestClient, &testDept)
  if err != nil {
    fmt.Printf(err.Error())
  }
  */
}
