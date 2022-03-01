package greenhouse

import(
  "github.com/carnegierobotics/greenhouse-client-go/internal/http"
  "github.com/carnegierobotics/greenhouse-client-go/internal/utils"
)

type Office struct {
  Id int `json:"id"`
  Name string `json:"name"`
  Location map[string]string `json:"location"`
  PrimaryContactUserId int `json:"primary_contact_user_id"`
  ParentId int `json:"parent_id"`
  ChildIds []int `json:"child_ids"`
  /* Not in our product tier.
  ParentOfficeExternalId string `json:"parent_office_external_id"`
  ChildOfficeExternalIds []string `json:"child_office_external_ids"`
  ExternalId string `json:"external_id"`
  */
}

func GetOffice(c *http.Client, id int) (*Office, error) {
  var obj Office
  err := utils.GetById(c, "offices", id, obj)
  if err != nil {
    return nil, err
  }
  return &obj, nil
}

func CreateOffice(c *http.Client, obj *Office) error {
  err := utils.Create(c, "offices", obj)
  if err != nil {
    return err
  }
  return nil
}

func UpdateOffice(c *http.Client, obj *Office) error {
  err := utils.Update(c, "offices", obj.Id, obj)
  if err != nil {
    return err
  }
  return nil
}
