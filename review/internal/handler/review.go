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

	modelObj := dto.FromDTO(req)

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

	returnList := dto.ToDTO(list)

	response.Response(c, http.StatusOK, op, gin.H{"reviews": returnList})
}
