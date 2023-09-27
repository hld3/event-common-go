package events

type UserDataEvent struct {
	NodeId   string `json:"nodeId"`
	UserId   string `json:"userId"`
	Username string `json:"username"`
}
