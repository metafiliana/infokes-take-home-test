package repositories

import (
	"github.com/infokes-take-home-test/backend/models"
	"gorm.io/gorm"
)

type folderRepository struct {
	db *gorm.DB
}

func NewFolderRepository(db *gorm.DB) FolderRepository {
	return &folderRepository{db: db}
}

type FolderRepository interface {
	GetFolders() ([]models.Folder, error)
	GetSubFolders(folderId int) ([]models.Folder, error)
}

func (r *folderRepository) GetFolders() ([]models.Folder, error) {
	var folders []models.Folder
	if err := r.db.Preload("Files").Where("parent_id IS NULL").Find(&folders).Error; err != nil {
		return nil, err
	}
	return folders, nil
}

func (r *folderRepository) GetSubFolders(folderId int) ([]models.Folder, error) {
	var folders []models.Folder
	if err := r.db.Preload("Files").Where("parent_id = ?", folderId).Find(&folders).Error; err != nil {
		return nil, err
	}
	return folders, nil
}
