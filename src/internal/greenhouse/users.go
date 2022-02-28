package greenhouse

import ()

type UserBasics struct {
  Id int `json:"id"`
  Name string `json:"name"`
  FirstName string `json:"first_name"`
  LastName string `json:"last_name"`
  EmployeeId string `json:"employee_id"`
}

type User struct {
  UserBasics
  PrimaryEmail string `json:"primary_email_address"`
  UpdatedAt string `json:"updated_at"`
  CreatedAt string `json:"created_at"`
  Disabled bool `json:"disabled"`
  SiteAdmin bool `json:"site_admin"`
  Emails []string `json:"emails"`
  LinkedCandidateIds []int `json:"linked_candidate_ids"`
}
