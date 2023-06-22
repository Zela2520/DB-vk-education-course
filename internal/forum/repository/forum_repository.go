package repository

import (
	"TechnoParkDBProject/internal/forum/models"
	usersModels "TechnoParkDBProject/internal/user/models"
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
)

type ForumRepository struct {
	Conn *pgxpool.Pool
}

func NewUserRepository(con *pgxpool.Pool) *ForumRepository {
	return &ForumRepository{
		Conn: con,
	}
}

func (forumRep *ForumRepository) CreateForum(forum *models.Forum) error {
	query := `INSERT INTO forum (title, user_nickname, slug, posts, threads)
			VALUES ($1, $2, $3, $4, $5)`

	_, err := forumRep.Conn.Exec(context.Background(), query, forum.Tittle,
		forum.UserNickname, forum.Slug, forum.Posts, forum.Threads,
	)
	return err
}

func (forumRep *ForumRepository) GetForumBySlug(slug string) (*models.Forum, error) {
	query := `SELECT title, user_nickname, slug, posts, threads from forum
			where slug = $1`
	forum := &models.Forum{}

	err := forumRep.Conn.QueryRow(context.Background(), query,
		slug).Scan(&forum.Tittle, &forum.UserNickname, &forum.Slug,
		&forum.Posts, &forum.Threads)

	if err != nil {
		return nil, err
	}
	return forum, nil
}

func (forumRep *ForumRepository) GetUsersByForum(forumSlug, since string, limit int, desc bool) ([]*usersModels.User, error) {
	query := fmt.Sprintf(`select uf.nickname, uf.fullname, uf.about, uf.email from users_to_forums as uf
			where uf.forum = '%s'`, forumSlug)
	if desc && since != "" {
		query += fmt.Sprintf(` and uf.nickname < '%s'`, since)
	} else if since != "" {
		query += fmt.Sprintf(` and uf.nickname > '%s'`, since)
	}
	query += ` order by uf.nickname `
	if desc {
		query += "desc"
	}
	query += fmt.Sprintf(` limit %d`, limit)
	rows, err := forumRep.Conn.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	users := make([]*usersModels.User, 0)

	for rows.Next() {
		user := &usersModels.User{}
		err := rows.Scan(&user.Nickname, &user.FullName, &user.About, &user.Email)
		if err != nil {
			fmt.Println(err)
		}
		users = append(users, user)
	}
	return users, nil
}
