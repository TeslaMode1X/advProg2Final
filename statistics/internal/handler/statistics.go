package handler

import (
	"github.com/TeslaMode1X/advProg2Final/statistics/internal/handler/response"
	interfaces "github.com/TeslaMode1X/advProg2Final/statistics/internal/interface"
	"github.com/gin-gonic/gin"
	"net/http"
)

type StatisticsHandler struct {
	service interfaces.StatisticsService
}

func NewStatisticsHandler(service interfaces.StatisticsService) *StatisticsHandler {
	return &StatisticsHandler{service: service}
}

func (s *StatisticsHandler) GetUsersStatistics(c *gin.Context) {
	const op = "statistics.handler.GetUsersStatistics"

	userStats, err := s.service.GetUsersStatisticsService()
	if err != nil {
		response.Response(c, http.StatusInternalServerError, op, err)
		return
	}

	response.Response(c, http.StatusOK, op, userStats)
}

func (s *StatisticsHandler) GetRecipesStatistics(c *gin.Context) {
	const op = "statistics.handler.GetRecipesStatistics"

	recipeStats, err := s.service.GetRecipesStatisticsService()
	if err != nil {
		response.Response(c, http.StatusInternalServerError, op, err)
		return
	}

	response.Response(c, http.StatusOK, op, recipeStats)
}

func (s *StatisticsHandler) GetRecipeStatByID(c *gin.Context) {
	const op = "statistics.handler.GetRecipeStatByID"

	id, _ := c.Params.Get("id")

	recipeStat, err := s.service.GetRecipeStatByIDService(id)
	if err != nil {
		response.Response(c, http.StatusInternalServerError, op, err)
		return
	}

	response.Response(c, http.StatusOK, op, recipeStat)
}
