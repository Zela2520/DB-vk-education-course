package forum

import (
	"TechnoParkDBProject/internal/forum/models"
	usersModels "TechnoParkDBProject/internal/user/models"
)

type Usecase interface {
	CreateForum(forum *models.Forum) error
	GetForumBySlug(slug string) (*models.Forum, error)
	GetUsersByForum(formSlug, since string, limit int, desc bool) ([]*usersModels.User, error)
}
