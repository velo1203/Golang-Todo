package models

type Todo struct {
	ID        uint   `gorm:"primaryKey;autoIncrement"`
	Title     string `gorm:"not null"`
	Completed bool
	UserID    string
}

type CreateRequest struct {
	Title string `json:"title,omitempty"`
}
type ReadRequest struct {
	ID uint `json:"id"` // 조회할 Todo 항목의 ID
}
type UpdateRequest struct {
	ID        uint `json:"id"`        // 업데이트할 Todo 항목의 ID
	Completed bool `json:"completed"` // 변경할 완료 상태, 선택 사항
}
type DeleteRequest struct {
	ID uint `json:"id"` // 삭제할 Todo 항목의 ID
}
