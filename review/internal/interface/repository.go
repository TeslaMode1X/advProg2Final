package interfaces

import (
	"github.com/TeslaMode1X/advProg2Final/review/internal/model"
	"github.com/TeslaMode1X/advProg2Final/review/internal/repository/dao"
)

type ReviewRepository interface {
	ReviewCreateRepo(model *model.Review) (string, error)
	ReviewListRepo() ([]*dao.ReviewEntity, error)
	ReviewByIDRepo(id string) (*dao.ReviewEntity, error)
	ReviewUpdateRepo(model *model.Review) error
	ReviewDeleteRepo(id string) error
}
