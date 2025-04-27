package handler

import (
	"fmt"
	"github.com/TeslaMode1X/advProg2Final/recipe/internal/handler/dto"
	"github.com/TeslaMode1X/advProg2Final/recipe/internal/handler/response"
	interfaces "github.com/TeslaMode1X/advProg2Final/recipe/internal/interface"
	"github.com/TeslaMode1X/advProg2Final/recipe/internal/model"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"net/http"
	"os"
	"time"
)

type RecipeHandler struct {
	recipeService interfaces.RecipeService
}

func NewRecipeHandler(serv interfaces.RecipeService) *RecipeHandler {
	return &RecipeHandler{
		recipeService: serv,
	}
}

func (h *RecipeHandler) RecipeList(c *gin.Context) {
	const op = "handler.recipe.RecipeList"

	recipes, err := h.recipeService.RecipeListService()
	if err != nil {
		response.Response(c, http.StatusInternalServerError, op, err.Error())
		return
	}

	response.Response(c, http.StatusOK, op, recipes)
}

func (h *RecipeHandler) RecipeByID(c *gin.Context) {
	const op = "handler.recipe.RecipeByID"

	idStr := c.Param("id")

	recipeObject, err := h.recipeService.RecipeByIDService(idStr)
	if err != nil {
		response.Response(c, http.StatusInternalServerError, op, gin.H{"error": err})
		return
	}

	recipeDto := &dto.RecipeResponse{
		AuthorID:    recipeObject.AuthorID.String(),
		Title:       recipeObject.Title,
		Description: recipeObject.Description,
	}

	response.Response(c, http.StatusOK, op, gin.H{"recipeObject": recipeDto})
}

func (h *RecipeHandler) RecipeCreate(c *gin.Context) {
	const op = "handler.recipe.RecipeCreate"

	form, err := c.MultipartForm()
	if err != nil {
		response.Response(c, http.StatusBadRequest, op, gin.H{"error": "could not read multipart form"})
		return
	}

	var req dto.RecipeRequest
	req.Title = form.Value["title"][0]
	req.Description = form.Value["description"][0]
	req.AuthorID = form.Value["id"][0]

	if req.Title == "" || req.AuthorID == "" {
		response.Response(c, http.StatusBadRequest, op, gin.H{"error": "title and author_id are required"})
		return
	}

	idUUID, err := uuid.NewV4()
	if err != nil {
		response.Response(c, http.StatusInternalServerError, op, gin.H{"error": err.Error()})
		return
	}

	files := form.File["photos"]
	var photoPaths []string
	for _, file := range files {
		filename := fmt.Sprintf("%s-%s", idUUID.String(), file.Filename)

		dir := "../photo"
		filepath := fmt.Sprintf("%s/%s", dir, filename)

		if err = os.MkdirAll(dir, 0755); err != nil {
			response.Response(c, http.StatusInternalServerError, op, gin.H{"error": "directory creation error"})
			return
		}

		if err = c.SaveUploadedFile(file, filepath); err != nil {
			response.Response(c, http.StatusInternalServerError, op, gin.H{"error": "photo save error"})
			return
		}

		photoPaths = append(photoPaths, fmt.Sprintf("photo/%s", filename))
	}

	authorUUID := uuid.FromStringOrNil(req.AuthorID)

	recipe := model.Recipe{
		ID:            idUUID,
		AuthorID:      authorUUID,
		Title:         req.Title,
		Description:   req.Description,
		Photos:        photoPaths,
		AverageRating: 0.0,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	id, err := h.recipeService.RecipeCreateService(recipe)
	if err != nil {
		response.Response(c, http.StatusInternalServerError, op, gin.H{"error": err.Error()})
		return
	}

	response.Response(c, http.StatusOK, op, gin.H{"id": id})
}

func (h *RecipeHandler) RecipeUpdate(c *gin.Context) {
	const op = "handler.recipe.RecipeUpdate"

	form, err := c.MultipartForm()
	if err != nil {
		response.Response(c, http.StatusBadRequest, op, gin.H{"error": "could not read multipart form"})
		return
	}

	recipeID := form.Value["id"][0]
	if recipeID == "" {
		response.Response(c, http.StatusBadRequest, op, gin.H{"error": "recipe ID is required"})
		return
	}

	title := form.Value["title"][0]
	description := form.Value["description"][0]

	if title == "" {
		response.Response(c, http.StatusBadRequest, op, gin.H{"error": "title is required"})
		return
	}

	idUUID := uuid.FromStringOrNil(recipeID)
	if idUUID == uuid.Nil {
		response.Response(c, http.StatusBadRequest, op, gin.H{"error": "invalid recipe ID"})
		return
	}

	var photoPaths []string
	files := form.File["photos"]
	if len(files) > 0 {
		for _, file := range files {
			filename := fmt.Sprintf("%s-%s", idUUID.String(), file.Filename)

			dir := "../photo"
			filepath := fmt.Sprintf("%s/%s", dir, filename)

			if err = os.MkdirAll(dir, 0755); err != nil {
				response.Response(c, http.StatusInternalServerError, op, gin.H{"error": "directory creation error"})
				return
			}

			if err = c.SaveUploadedFile(file, filepath); err != nil {
				response.Response(c, http.StatusInternalServerError, op, gin.H{"error": "photo save error"})
				return
			}

			photoPaths = append(photoPaths, fmt.Sprintf("photo/%s", filename))
		}
	}

	recipe := model.Recipe{
		ID:          idUUID,
		Title:       title,
		Description: description,
	}

	if len(photoPaths) > 0 {
		recipe.Photos = photoPaths
	}

	err = h.recipeService.RecipeUpdateService(recipe)
	if err != nil {
		response.Response(c, http.StatusInternalServerError, op, gin.H{"error": err.Error()})
		return
	}

	response.Response(c, http.StatusOK, op, gin.H{"id": recipeID})
}

func (h *RecipeHandler) RecipeDelete(c *gin.Context) {
	const op = "handler.recipe.RecipeDelete"

	idStr := c.Param("id")

	err := h.recipeService.RecipeDeleteService(idStr)
	if err != nil {
		response.Response(c, http.StatusInternalServerError, op, gin.H{"error": err.Error()})
		return
	}

	response.Response(c, http.StatusOK, op, gin.H{"deleted": true})
}
