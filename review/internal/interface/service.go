package interfaces

import (
	"github.com/TeslaMode1X/advProg2Final/review/internal/model"
	"github.com/TeslaMode1X/advProg2Final/review/internal/repository/dao"
)

type ReviewService interface {
	ReviewCreateService(model *model.Review) (string, error)
	ReviewListService() ([]*dao.ReviewEntity, error)
}
