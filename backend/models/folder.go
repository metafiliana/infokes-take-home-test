package models

import "time"

type Folder struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	ParentID  int       `json:"parent_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	// ParentFolder *Folder   `json:"parent_folder" gorm:"->"`
}

func (Folder) TableName() string {
	return "folder"
}
