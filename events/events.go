package events

type UserDataEvent struct {
	NodeId         string `json:"nodeId"`
	UserId         string `json:"userId"`
	Username       string `json:"username"`
	Status         string `json:"status"`
	Comment        string `json:"comment"`
	ReceiveUpdates bool   `json:"receiveUpdates"`
}

type GroupDataEvent struct {
	Name             string `json:"name"`
	Code             string `json:"code"`
	GroupId          string `json:"groupId"`
	OwnerId          string `json:"ownerId"`
	KnownLanguage    string `json:"knownLanguage"`
	LearningLanguage string `json:"learningLanguage"`
}
