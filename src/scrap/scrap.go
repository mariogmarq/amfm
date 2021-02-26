package scrap

import (
	"github.com/go-rod/rod"
	"github.com/mariogmarq/amfm/src/scrap/login"
)

//Login returns a browser with an opened google session with given credentials, allows google and
//go.ugr.es accounts only. It does not have any kind of error handling yet
func Login(email string, password string) (*rod.Browser, error) {
	return login.Login(email, password)
}
