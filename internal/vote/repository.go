package vote

import "TechnoParkDBProject/internal/vote/models"

type Repository interface {
	CreateNewVote(vote *models.Vote) error
	UpdateVote(vote *models.Vote) (int, error)
}
