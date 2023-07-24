package server

import (
	"net/http"

	"github.com/TonyQuedeville/social-network/app-social-network/pkg/api"
	"github.com/TonyQuedeville/social-network/database-manager/database"
)

func InitServer() {
	database.OpenDatabase()
	defer database.CloseDatabase()

	mux := http.NewServeMux()

	// user
	mux.Handle("/user/register", ApplyMiddleware(http.HandlerFunc(api.UserRegister)))
	mux.Handle("/user/login", ApplyMiddleware(http.HandlerFunc(api.UserLogin)))
	mux.Handle("/users", ApplyMiddleware(http.HandlerFunc(api.GetUsers)))
	mux.Handle("/user/", ApplyMiddleware(http.HandlerFunc(api.GetUserById)))
	mux.Handle("/addfollower/", ApplyMiddleware(http.HandlerFunc(api.AddFollower)))
	mux.Handle("/acceptfollower/", ApplyMiddleware(http.HandlerFunc(api.AcceptFollower)))
	mux.Handle("/refusefollower/", ApplyMiddleware(http.HandlerFunc(api.RefuseFollower)))
	mux.Handle("/supfollower/", ApplyMiddleware(http.HandlerFunc(api.SupFollower)))
	// mux.Handle("/follower/", ApplyMiddleware(http.HandlerFunc(api.GetFollowerUserById)))

	// post
	mux.Handle("/newpost", ApplyMiddleware(http.HandlerFunc(api.NewPost)))
	mux.Handle("/userposts/", ApplyMiddleware(http.HandlerFunc(api.PostByUserId)))
	//mux.Handle("/groupposts/", ApplyMiddleware(http.HandlerFunc(api.PostByGrouypId)))

	http.ListenAndServe(ADRESS, mux)
}
