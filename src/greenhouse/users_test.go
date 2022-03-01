package greenhouse

import (
  "testing"
  "github.com/carnegierobotics/greenhouse-client-go/internal/http"
  "github.com/carnegierobotics/greenhouse-client-go/internal/utils"
)

func TestMain(m *testing.M) {
  harvestClient := http.Client{BaseUrl: harvestUrl, Token: token, OnBehalfOf: onBehalfOf}
  harvestClient.BuildResty()
}

func TestGetAllUsers(*testing.T) {
  var allUsers []*User
  err := utils.GetAll(&harvestClient, "users", &allUsers)
  if err != nil {
    t.Errorf(err.Error())
  }
  for _, u := range allUsers {
    fmt.Printf("%s\n", u.Name)
  }
}

func TestGetUserById(*testing.T) {
  var user *User
  err = utils.GetById(&harvestClient, "users", 4005448005, &user)
  if err != nil {
    t.Errorf(err.Error())
  }
  fmt.Printf("%v", user.Name)
}
