package bike

type Info struct {
	Id       int64  `json:"id"`
	BikeName string `json:"bikeName"`
	Created  int64  `json:"created"`
	Updated  int64  `json:"updated"`
}

func (Info) TableName() string {
	return "bike"
}
