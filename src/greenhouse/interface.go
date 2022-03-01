package greenhouse

type GreenhouseObject interface {
  Create()
  Read()
  Update()
  Delete()
  GetId()
}
