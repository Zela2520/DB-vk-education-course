package vote

import (
	threadModels "TechnoParkDBProject/internal/thread/models"
	voteModels "TechnoParkDBProject/internal/vote/models"
)

type Usecase interface {
	CreateNewVote(vote *voteModels.Vote, slugOrID string) (*threadModels.Thread, error)
}
