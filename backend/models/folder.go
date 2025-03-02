package models

import "time"

type Folder struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	ParentID  int       `json:"parent_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Files     []File    `json:"files" gorm:"foreignKey:FolderID"`
}

func (Folder) TableName() string {
	return "folder"
}
