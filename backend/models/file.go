package models

import "time"

type File struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	FolderID  int       `json:"folder_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (File) TableName() string {
	return "file"
}
