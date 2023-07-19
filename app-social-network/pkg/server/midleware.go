package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/TonyQuedeville/social-network/app-social-network/pkg/api"
	"github.com/TonyQuedeville/social-network/database-manager/structs/user"
)

// func exampleMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		// Our middleware logic goes here...
// 		next.ServeHTTP(w, r)
// 	})
// }

func ApplyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next = AlowCorse(next)
		next = CheckSession(next)
		next.ServeHTTP(w, r)
	})
}

func AlowCorse(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*") // A CHANGER POUR UN POTENTIEL DEPLOIMENT
		next.ServeHTTP(w, r)
	})
}

func CheckSession(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ck, err := r.Cookie("session")
		u_id := uint64(0)
		if err == nil {
			u_id = user.GetUserIdByUuid(ck.Value)
		}
		fmt.Println("Middleware Session idfind: ", u_id)
		ctx := context.WithValue(r.Context(), api.USER_ID, u_id)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}