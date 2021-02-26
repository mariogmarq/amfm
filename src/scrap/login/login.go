package login

import (
	"strings"

	"github.com/go-rod/rod"
)

var url = "https://accounts.google.com/signin/v2/identifier?ltmpl=meet&continue=https%3A%2F%2Fmeet.google.com%3Fhs%3D193&_ga=2.197449007.1812202053.1614183595-1863875012.1614183595&flowName=GlifWebSignIn&flowEntry=ServiceLogin"

//Login returns a browser with session already opened, this session can be opened with the following
//types of accounts: google, go.ugr.es. If you want to open add other type of account create an
//issue at https://github.com/mariogmarq/amfm
func Login(email string, password string) *rod.Browser {
	//Create browser and travel to login page
	browser := rod.New().MustConnect()
	page := browser.MustPage(url)

	//Enter email and wait for being redirected to ugr login
	page.MustSearch("input[id]").MustInput(email)
	page.MustSearch(`div[id="identifierNext"`).MustClick()

	//Search for handler
	handler := findLoginHandler(email)
	if handler == nil {
		panic("Invalid email")
	}

	handler(email, password, browser, page)

	//Wait for meet page(Here I use an image with an specific html tag)
	page.MustSearch(`img[role="img"]`)

	return browser
}

func loginUGR(email string, password string, browser *rod.Browser, page *rod.Page) {
	page.MustSearch(`img[src="https://idp.ugr.es/go/module.php/themeSURFnet/logo.png"`)

	//Login at ugr page
	page.MustSearch(`input[id="username"]`).MustInput(email)
	page.MustSearch(`input[id="password"]`).MustInput(password)
	page.MustSearch(`input[value="Login"]`).MustClick()
}

func loginGoogle(email string, password string, browser *rod.Browser, page *rod.Page) {
	page.MustSearch(`input[name="password"]`).MustInput(password)
	page.MustSearch(`div[id="passwordNext"]`).MustClick()
}

func findLoginHandler(email string) func(email string, password string, browser *rod.Browser, page *rod.Page) {
	words := strings.Split(email, "@")
	if len(words) != 2 {
		return nil
	}

	switch words[1] {
	case "go.ugr.es":
		return loginUGR
	case "gmail.com":
		return loginGoogle
	default:
		return nil
	}
}
