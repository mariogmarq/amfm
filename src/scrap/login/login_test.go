package login

import (
	"testing"

	"github.com/go-rod/rod"
	"github.com/google/go-cmp/cmp"
)

type findHandlerTestCase struct {
	email    string
	expected func(email string, password string, browser *rod.Browser, page *rod.Page)
}

func testFindHandler(t *testing.T) {
	cases := []findHandlerTestCase{
		{"example@go.ugr.es", loginUGR},
		{"example@gmail.com", loginGoogle},
		{"notvalid@email.com", nil},
		{"notvalidemail", nil},
	}

	for i := range cases {
		if handler := findLoginHandler(cases[i].email); !cmp.Equal(handler, cases[i].expected) {
			t.Errorf("Invalid function in testcase number %d with email %s\n", i, cases[i].email)
		}
	}
}
