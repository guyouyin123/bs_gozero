package user

type Info struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Created int64  `json:"created"`
	Updated int64  `json:"updated"`
}

func (Info) TableName() string {
	return "user"
}
