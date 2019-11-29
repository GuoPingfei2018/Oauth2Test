package main

import (
	"github.com/go-redis/redis"
	"gopkg.in/oauth2.v3/errors"
	"gopkg.in/oauth2.v3/manage"

	"gopkg.in/oauth2.v3/server"

	"log"
	"net/http"

	oredis "gopkg.in/go-oauth2/redis.v3"
)

func main() {
	manager := manage.NewDefaultManager()


	// use redis token store
	tokenStore := oredis.NewRedisStore(&redis.Options{
		Addr: "127.0.0.1:6379",
		DB: 5,
	})

	manager.MapTokenStorage(tokenStore)



	srv := server.NewDefaultServer(manager)
	srv.SetAllowGetAccessRequest(true)

	srv.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		log.Println("Internal Error:", err.Error())
		return
	})

	srv.SetResponseErrorHandler(func(re *errors.Response) {
		log.Println("Response Error:", re.Error.Error())
	})

	http.HandleFunc("/authorize", func(w http.ResponseWriter, r *http.Request) {
		err := srv.HandleAuthorizeRequest(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	})

	http.HandleFunc("/token", func(w http.ResponseWriter, r *http.Request) {
		srv.HandleTokenRequest(w, r)
	})

	log.Fatal(http.ListenAndServe(":9096", nil))
}