package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/infokes-take-home-test/backend/usecases"
)

type FolderHandler interface {
	GetFolders(ctx *gin.Context)
	GetSubFolders(ctx *gin.Context)
}

type folderHandler struct {
	folderUsecase usecases.FolderUsecase
}

func NewFolderHandler(folderUsecase usecases.FolderUsecase) FolderHandler {
	return &folderHandler{folderUsecase}
}

func (h *folderHandler) GetFolders(c *gin.Context) {
	out, err := h.folderUsecase.GetFolders()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "success",
		"data":    out,
	})
}

func (h *folderHandler) GetSubFolders(c *gin.Context) {
	folderId := c.Params.ByName("id")
	if folderId == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "folder id is required",
		})
		return
	}

	folderIdInt, err := strconv.Atoi(folderId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "folder id must be a number",
		})
		return
	}

	out, err := h.folderUsecase.GetSubFolders(folderIdInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "success",
		"data":    out,
	})
}
