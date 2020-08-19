package models

import "time"

type Memo struct {
	ID          uint32
	Title       string
	Body        string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Sender      User
	To          []*User `gorm:"many2many:users_id;"`
	ReadReceipt []*User
	Note        []Note
}

type Note struct {
	ID        uint32
	Body      string
	CreatedAt time.Time
}
