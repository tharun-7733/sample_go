package middleware

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/avukadin/goapi/internal/shared"
	"github.com/avukadin/goapi/internal/tools"

	log "github.com/sirupsen/logrus"
)

var UnAuthorizedError = errors.New(fmt.Sprintf("Invalid username or token."))

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var username string = r.URL.Query().Get("username")
		var token string = r.Header.Get("Authorization")
		var err error

		if username == "" {
			shared.RequestErrorHandler(w, UnAuthorizedError)
			return
		}

		var database *tools.DatabaseInterface

		database, err = tools.NewDatabase()

		if err != nil {
			shared.InternalErrorHandler(w)
			return
		}

		var loginDetails *tools.LoginDetails

		loginDetails = (*database).GetUserLoginDetails(username)

		if loginDetails == nil || token != (*loginDetails).AuthToken {
			log.Error(UnAuthorizedError)
			shared.RequestErrorHandler(w, UnAuthorizedError)
			return
		}

		next.ServeHTTP(w, r)
	})
}