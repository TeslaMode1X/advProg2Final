package interfaces

import "github.com/TeslaMode1X/advProg2Final/review/internal/model"

type ReviewService interface {
	ReviewCreateService(model *model.Review) (string, error)
}
