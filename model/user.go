package model

type User struct {
	Uid      int    `json:"uid" form:"uid"`
	NickName string `json:"nick" form:"nick"`
}
