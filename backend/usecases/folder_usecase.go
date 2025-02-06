package usecases

import (
	"github.com/infokes-take-home-test/backend/models"
	"github.com/infokes-take-home-test/backend/repositories"
)

type FolderUsecase interface {
	GetFolders() ([]models.Folder, error)
	GetSubFolders(folderId int) ([]models.Folder, error)
}

type folderUsecase struct {
	folderRepo repositories.FolderRepository
}

func NewFolderUsecase(folderRepo repositories.FolderRepository) FolderUsecase {
	return &folderUsecase{folderRepo}
}

func (u *folderUsecase) GetFolders() ([]models.Folder, error) {
	return u.folderRepo.GetFolders()
}

func (u *folderUsecase) GetSubFolders(folderId int) ([]models.Folder, error) {
	return u.folderRepo.GetSubFolders(folderId)
}
