package interfaces

import "github.com/TeslaMode1X/advProg2Final/review/internal/model"

type ReviewRepository interface {
	ReviewCreateRepo(model *model.Review) (string, error)
}
