package models

const (
	TaskActionCrafting = "crafting"
	TaskStatusPending  = "pending"
	TaskStatusRunning  = "running"
	TaskStatusSuccess  = "success"
	TaskStatusError    = "error"
)

type CreateTask struct {
	Action   string `json:"action"`
	Item     string `json:"item"`
	Quantity int    `json:"quantity"`
	Status   string `json:"status"`
}

type Task struct {
	Id          string `bson:"_id" json:"id"`
	Action      string `bson:"action" json:"action"`
	Item        string `bson:"item" json:"item"`
	Quantity    int    `bson:"quantity" json:"quantity"`
	Status      string `bson:"status" json:"status"`
	ErrorReason string `bson:"errorReason" json:"errorReason"`
}
