package Hash

import (
	"errors"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func GeneratePass(password *[]byte) error {

	var err1 error
	*password, err1 = bcrypt.GenerateFromPassword(*password, 8)

	if err1 != nil {
		// Err("500 Internal Server Error", http.StatusInternalServerError, w, r)
		return errors.New("Internal Server Error " + string(http.StatusInternalServerError))
	}
	return nil
}
