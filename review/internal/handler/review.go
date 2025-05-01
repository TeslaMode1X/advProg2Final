package handler

import (
	"github.com/TeslaMode1X/advProg2Final/review/internal/handler/dto"
	"github.com/TeslaMode1X/advProg2Final/review/internal/handler/response"
	interfaces "github.com/TeslaMode1X/advProg2Final/review/internal/interface"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ReviewHandler struct {
	userService interfaces.ReviewService
}

func NewReviewHandler(userService interfaces.ReviewService) *ReviewHandler {
	return &ReviewHandler{
		userService: userService,
	}
}

func (h *ReviewHandler) ReviewCreateHandler(c *gin.Context) {
	const op = "handler.review.ReviewCreateHandler"

	var req dto.ReviewCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Response(c, http.StatusInternalServerError, op, err.Error())
		return
	}

	err := req.Validate()
	if err != nil {
		response.Response(c, http.StatusBadRequest, op, err.Error())
		return
	}

	modelObj := dto.ConvertCreateRequestToModel(req)

	id, err := h.userService.ReviewCreateService(modelObj)
	if err != nil {
		response.Response(c, http.StatusInternalServerError, op, err.Error())
		return
	}

	response.Response(c, http.StatusOK, op, gin.H{"id": id})
}

func (h *ReviewHandler) ReviewListHandler(c *gin.Context) {
	const op = "handler.review.ReviewListHandler"

	list, err := h.userService.ReviewListService()
	if err != nil {
		response.Response(c, http.StatusInternalServerError, op, err.Error())
		return
	}

	returnList := dto.ConvertEntitiesToDTOs(list)

	response.Response(c, http.StatusOK, op, gin.H{"reviews": returnList})
}

func (h *ReviewHandler) ReviewByIDHandler(c *gin.Context) {
	const op = "handler.review.ReviewByIDHandler"

	id, _ := c.Params.Get("id")

	object, err := h.userService.ReviewByIDService(id)
	if err != nil {
		response.Response(c, http.StatusInternalServerError, op, err.Error())
		return
	}

	returnObject := dto.ConvertEntityToDTO(*object)

	response.Response(c, http.StatusOK, op, returnObject)
}

func (h *ReviewHandler) ReviewUpdateHandler(c *gin.Context) {
	const op = "handler.review.ReviewUpdateHandler"

	var req dto.ReviewUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Response(c, http.StatusInternalServerError, op, err.Error())
		return
	}

	object := dto.ConvertUpdateRequestToModel(req)

	err := h.userService.ReviewUpdateService(object)
	if err != nil {
		response.Response(c, http.StatusInternalServerError, op, err.Error())
		return
	}

	response.Response(c, http.StatusOK, op, object)
}

func (h *ReviewHandler) ReviewDeleteHandler(c *gin.Context) {
	const op = "handler.review.ReviewDeleteHandler"

	id, _ := c.Params.Get("id")

	err := h.userService.ReviewDeleteService(id)
	if err != nil {
		response.Response(c, http.StatusInternalServerError, op, err.Error())
		return
	}

	response.Response(c, http.StatusOK, op, gin.H{"status": "successfully deleted"})
}
