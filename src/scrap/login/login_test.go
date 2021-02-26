package login

import (
	"testing"

	"github.com/go-rod/rod"
	"github.com/google/go-cmp/cmp"
)

type findHandlerTestCase struct {
	email    string
	expected func(email string, password string, browser *rod.Browser, page *rod.Page)
	err      bool
}

func testFindHandler(t *testing.T) {
	cases := []findHandlerTestCase{
		{"example@go.ugr.es", loginUGR, false},
		{"example@gmail.com", loginGoogle, false},
		{"notvalid@email.com", nil, true},
		{"notvalidemail", nil, true},
	}

	for i := range cases {
		if handler, err := findLoginHandler(cases[i].email); ((err != nil) != cases[i].err) || (!cmp.Equal(handler, cases[i].expected)) {
			t.Errorf("Invalid function in testcase number %d with email %s\n", i, cases[i].email)
		}
	}
}
