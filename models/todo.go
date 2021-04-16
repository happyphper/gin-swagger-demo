package models

type Todo struct {
	ID        int    `json:"id" example:"11012341234"`
	Title     string `json:"title" example:"完成TodoList布局"`
	Status    int    `json:"status" example:"1"`
	CreatedAt string `json:"created_at" example:"2006-01-02 15:04:05"`
	UpdatedAt string `json:"updated_at" example:"2006-01-02 15:04:05"`
}
