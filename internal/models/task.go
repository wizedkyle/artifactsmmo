package models

const (
	TaskStatusCancelled = "cancelled"
	TaskStatusPending   = "pending"
	TaskStatusRunning   = "running"
	TaskStatusSuccess   = "success"
	TaskStatusError     = "error"
)

type CancelTask struct {
	Reason string `json:"reason"`
}

type CreateTask struct {
	Action         string `json:"action,omitempty"`
	ActionCategory string `json:"actionCategory"`
	Monster        string `json:"monster,omitempty"`
	Item           string `json:"item,omitempty"`
	Quantity       int    `json:"quantity"`
	Character      string `json:"character"`
}

type Task struct {
	Id                string `bson:"_id" json:"id"`
	Action            string `bson:"action,omitempty" json:"action,omitempty"`
	ActionCategory    string `bson:"actionCategory" json:"actionCategory"`
	Item              string `bson:"item,omitempty" json:"item,omitempty"`
	Monster           string `bson:"monster,omitempty" json:"monster,omitempty"`
	Quantity          int    `bson:"quantity" json:"quantity"`
	QuantityCompleted int    `bson:"quantityCompleted" json:"quantityCompleted"`
	Status            string `bson:"status" json:"status"`
	ErrorReason       string `bson:"errorReason" json:"errorReason"`
	Character         string `bson:"character" json:"character"`
}
