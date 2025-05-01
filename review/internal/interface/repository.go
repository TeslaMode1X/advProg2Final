package interfaces

import (
	"github.com/TeslaMode1X/advProg2Final/review/internal/model"
	"github.com/TeslaMode1X/advProg2Final/review/internal/repository/dao"
)

type ReviewRepository interface {
	ReviewCreateRepo(model *model.Review) (string, error)
	ReviewListRepo() ([]*dao.ReviewEntity, error)
}
