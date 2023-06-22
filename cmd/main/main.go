package main

import (
	forumHTTP "TechnoParkDBProject/internal/forum/delivery/http"
	forumRepository "TechnoParkDBProject/internal/forum/repository"
	forumUsecase "TechnoParkDBProject/internal/forum/usecase"
	postHTTP "TechnoParkDBProject/internal/posts/delivery/http"
	postRepository "TechnoParkDBProject/internal/posts/repository"
	postUsecase "TechnoParkDBProject/internal/posts/usecase"
	threadHTTP "TechnoParkDBProject/internal/thread/delivery/http"
	threadRepositoru "TechnoParkDBProject/internal/thread/repository"
	threadUsecase "TechnoParkDBProject/internal/thread/usecase"
	userHTTP "TechnoParkDBProject/internal/user/delivery/http"
	"TechnoParkDBProject/internal/user/repository"
	"TechnoParkDBProject/internal/user/usecase"
	voteHTTP "TechnoParkDBProject/internal/vote/delivery/http"
	voteRepository "TechnoParkDBProject/internal/vote/repository"
	voteUsecase "TechnoParkDBProject/internal/vote/usecase"
	"context"
	"fmt"
	"os"

	"github.com/fasthttp/router"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/valyala/fasthttp"
)

func main() {
	router := router.New()
	dbpool, err := pgxpool.Connect(context.Background(),
		"host=localhost port=5432 user=myuser dbname=db_forum password=password sslmode=disable",
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer dbpool.Close()

	userRep := repository.NewUserRepository(dbpool)
	forumRep := forumRepository.NewUserRepository(dbpool)
	thredRep := threadRepositoru.NewThreadRepository(dbpool)
	postsRep := postRepository.NewPostsRepository(dbpool)
	voteRep := voteRepository.NewVoteRepository(dbpool)

	userUsecase := usecase.NewUserUsecase(userRep)
	forumUsec := forumUsecase.NewForumUsecase(forumRep)
	thredUsec := threadUsecase.NewThreadUsecase(thredRep)
	postUse := postUsecase.NewPostsUsecase(postsRep, thredRep, forumRep, userRep)
	voteUse := voteUsecase.NewVoteUsecase(voteRep, thredRep)

	userHTTP.NewUserHandler(router, userUsecase)
	forumHTTP.NewForumHandler(router, forumUsec, userUsecase)
	threadHTTP.NewThreadHandler(router, thredUsec, forumUsec)
	postHTTP.NewPostsHandler(router, postUse)
	voteHTTP.NewVoteHandler(router, voteUse)

	err = fasthttp.ListenAndServe(":5000", router.Handler)
	fmt.Println(err)
}
