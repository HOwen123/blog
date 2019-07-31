package models

type Message struct {
	Model
	Key     string `gorm:"unique_index; not null" json:"key"`
	Content string `json:"content"`
	UserId  int    `json:"user_id"`
	User    User   `json:"user"`
	NoteKey string `json:"note_key"`
	Praise  int    `gorm:"default:0" json:"praise"` //点赞数量
}
