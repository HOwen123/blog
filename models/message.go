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

func SaveMessage(message *Message) error {
	return db.Model(&Message{}).Save(message).Error
}

func QueryMessagesByNoteKey(noteKey string) (ms []*Message, err error) {
	return ms, db.Preload("User").Where("note_key = ?", noteKey).Order("updated_at desc").Find(&ms).Error
}

func QueryMessagesCountByNoteKey(noteKey string) (count int, err error) {
	return count, db.Model(&Message{}).Where("note_key = ?", noteKey).Count(&count).Error
}

func QueryPageMessagesByNoteKey(notekey string, pageno, pagesize int) (ms []*Message, err error) {
	return ms, db.Preload("User").Where("note_key = ?", notekey).Offset((pageno - 1) * pagesize).Limit(pagesize).Order("updated_at desc").Find(&ms).Error
}
