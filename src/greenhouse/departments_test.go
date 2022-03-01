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

func TestGetAllDepartments(*testing.T) {
  var allDepartments []*Department
  err := utils.GetAll(&harvestClient, "departments", &allDepartments)
  if err != nil {
    t.Errorf(err.Error())
  }
  for _, d := range allDepartments {
    fmt.Printf("%s\n", d.Name)
  }
}

func TestDepartmentCreate(*testing.T) {
  testDept := Department{Name: "test"}
  err := CreateDepartment(&harvestClient, &testDept)
  if err != nil {
    t.Fail()
  }
}
