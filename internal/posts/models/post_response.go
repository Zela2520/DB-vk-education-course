package models

import (
	modelsForum "TechnoParkDBProject/internal/forum/models"
	modelsThread "TechnoParkDBProject/internal/thread/models"
	modelsUser "TechnoParkDBProject/internal/user/models"
)

type PostResponse struct {
	Post   *Post                `json:"post"`
	User   *modelsUser.User     `json:"author,omitempty"`
	Forum  *modelsForum.Forum   `json:"forum,omitempty"`
	Thread *modelsThread.Thread `json:"thread,omitempty"`
}
