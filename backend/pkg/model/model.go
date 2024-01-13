package model

type Task struct {
	ID          string `bson:"_id"`
	Owner       string `bson:"owner"`
	Title       string `bson:"title"`
	Description string `bson:"description"`
	Done        bool   `bson:"done"`
}

type ApiReturn struct {
	Message string `json:"message,omitempty"`
	Id      string `json:"id,omitempty"`
}
