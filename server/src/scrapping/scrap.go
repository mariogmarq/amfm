package scrap

import (
	"github.com/go-rod/rod"
)

var url = "https://accounts.google.com/signin/v2/identifier?ltmpl=meet&continue=https%3A%2F%2Fmeet.google.com%3Fhs%3D193&_ga=2.197449007.1812202053.1614183595-1863875012.1614183595&flowName=GlifWebSignIn&flowEntry=ServiceLogin"

//Login login with your go.ugr.es account(University of Granada) and returns a browser with the
//session opened
func Login(email string, password string) *rod.Browser {
	//Create browser and travel to login page
	browser := rod.New().MustConnect()
	page := browser.MustPage(url)

	//Enter email and wait for being redirected to ugr login
	page.MustSearch("input[id]").MustInput(email)
	page.MustSearch(`div[id="identifierNext"`).MustClick()
	page.MustSearch(`img[src="https://idp.ugr.es/go/module.php/themeSURFnet/logo.png"`)

	//Login at ugr page
	page.MustSearch(`input[id="username"]`).MustInput(email)
	page.MustSearch(`input[id="password"]`).MustInput(password)
	page.MustSearch(`input[value="Login"]`).MustClick()

	//Wait for meet page(Here I use an image with an specific html tag)
	page.MustSearch(`img[role="img"]`)

	return browser
}
