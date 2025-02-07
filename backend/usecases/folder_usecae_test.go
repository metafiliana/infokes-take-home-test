package usecases

import (
	"testing"

	"github.com/infokes-take-home-test/backend/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockFolderRepository is a mock type for the FolderRepository
type MockFolderRepository struct {
	mock.Mock
}

func (m *MockFolderRepository) GetFolders() ([]models.Folder, error) {
	args := m.Called()
	return args.Get(0).([]models.Folder), args.Error(1)
}

func (m *MockFolderRepository) GetSubFolders(folderId int) ([]models.Folder, error) {
	args := m.Called(folderId)
	return args.Get(0).([]models.Folder), args.Error(1)
}

func TestGetSubFolders(t *testing.T) {
	testCases := []struct {
		name           string
		folderId       int
		mockSubfolders []models.Folder
		mockError      error
		expectedResult []models.Folder
		expectedError  bool
	}{
		{
			name:     "Successful retrieval of subfolders",
			folderId: 1,
			mockSubfolders: []models.Folder{
				{ID: 2, Name: "Subfolder 1", ParentID: 1},
				{ID: 3, Name: "Subfolder 2", ParentID: 1},
			},
			mockError: nil,
			expectedResult: []models.Folder{
				{ID: 2, Name: "Subfolder 1", ParentID: 1},
				{ID: 3, Name: "Subfolder 2", ParentID: 1},
			},
			expectedError: false,
		},
		{
			name:           "No subfolders found",
			folderId:       99,
			mockSubfolders: []models.Folder{},
			mockError:      nil,
			expectedResult: []models.Folder{},
			expectedError:  false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create a new mock repository
			mockRepo := new(MockFolderRepository)

			// Setup expectations
			mockRepo.On("GetSubFolders", tc.folderId).Return(tc.mockSubfolders, tc.mockError)

			// Create usecase with mock repository
			usecase := NewFolderUsecase(mockRepo)

			// Call the method
			result, err := usecase.GetSubFolders(tc.folderId)

			// Assertions
			if tc.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedResult, result)
			}

			// Verify that the mock was called with the correct arguments
			mockRepo.AssertExpectations(t)
		})
	}
}
