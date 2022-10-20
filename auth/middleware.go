package auth

import (
	"net/http"
	"strings"

	"github.com/adedaramola/golang-jwt-auth/utils"
	"github.com/julienschmidt/httprouter"
)

func Authenticate(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		authHeader := r.Header.Get("Authorization")
		headerInfo := strings.Split(authHeader, " ")

		if len(headerInfo) != 2 {
			utils.JSON(w, http.StatusUnauthorized, utils.H{"error": "Unauthorized"})
			return
		}

		// check if authorization type is bearer
		if headerInfo[0] != "Bearer" {
			utils.JSON(w, http.StatusUnauthorized, utils.H{"error": "Unauthorized"})
			return
		}

		token := headerInfo[1]
		// check token validity
		err := VerifyToken(token)
		if err != nil {
			utils.JSON(w, http.StatusUnauthorized, utils.H{"error": err.Error()})
			return
		}

		next(w, r, p)
	}
}
