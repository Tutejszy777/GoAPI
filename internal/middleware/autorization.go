package middleware

import (
	"errors"
	"net/http"

	"github.com/Tutejszy777/GoAPI/api"
	"github.com/Tutejszy777/GoApi/api"
	"github.com/Tutejszy777/GoApi/internal/tools"
	log "github.com/sirupsen/logrus"
)

var UnAuthorizedError = errors.New("Invalid userName or Token.")
func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request)) {
		var username string = r.URL.QUERY().Get("username")
		var token = r.Header.Get("Authorization")
		var err error

		if username == "" || token = ""{
			log.Error(UnAuthorizedError)
			api.RequestErrorHandler(w, UnAuthorizedError)
			return 
		}

		var database *tools.DatabaseInterface
		database, err = tools.NewDatabase()
		if err != nil {
			api.InternalErrorHabdler(w)
			return
		}

		var loginDetails *tools.loginDetails
		loginDetails = (*database).GetUserLoginDetails(username)

		if (loginDetails == nil || (toke != (*loginDetails).AuthToken)) {
			log.Error(UnAuthorizedError)
			api.RequestErrorHandler(w, UnAuthorizedError)
			return
		}

		next.ServeHTTP(w, r)

	}
}