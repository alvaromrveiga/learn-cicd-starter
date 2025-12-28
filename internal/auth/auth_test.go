package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	type test struct {
		name         string
		expect       string
		httpHeader   http.Header
		errorMessage string
	}

	tests := [3]test{
		{
			name:         "Return API Key",
			expect:       "TEST_API_KEY",
			httpHeader:   http.Header{"Authorization": []string{"ApiKey TEST_API_KEY"}},
			errorMessage: "",
		},
		{
			name:         "Handle no Authorization header",
			expect:       "",
			httpHeader:   http.Header{},
			errorMessage: "no authorization header included",
		},
		{
			name:         "Handle malformed Authorization header",
			expect:       "",
			httpHeader:   http.Header{"Authorization": []string{"TEST_API_KEY"}},
			errorMessage: "malformed authorization header",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res, err := GetAPIKey(test.httpHeader)

			if res != test.expect {
				t.Errorf("Response error, expected %s, got %s", test.expect, res)
			}

			if err != nil && err.Error() != test.errorMessage {
				t.Errorf("Error, expected error %s, got error %s", test.errorMessage, err.Error())
			}
		})
	}

}
