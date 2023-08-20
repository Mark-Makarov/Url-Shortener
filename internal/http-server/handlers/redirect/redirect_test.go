package redirect_test

import (
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"Url-Shortener/internal/http-server/handlers/redirect"
	"Url-Shortener/internal/http-server/handlers/redirect/mocks"
	"Url-Shortener/internal/lib/api"
	"Url-Shortener/internal/lib/logger/slogdiscard"
)

func TestSaveHandler(t *testing.T) {
	cases := []struct {
		name      string
		alias     string
		url       string
		respError string
		mockError error
	}{
		{
			name:  "Success",
			alias: "test_alias",
			url:   "https://www.google.com/",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			urlGet := mocks.NewURLGet(t)

			if tc.respError == "" || tc.mockError != nil {
				urlGet.On("GetURL", tc.alias).
					Return(tc.url, tc.mockError).Once()
			}

			r := chi.NewRouter()
			r.Get("/{alias}", redirect.New(slogdiscard.NewDiscardLogger(), urlGet))

			ts := httptest.NewServer(r)
			defer ts.Close()

			redirectedToURL, err := api.GetRedirect(ts.URL + "/" + tc.alias)
			require.NoError(t, err)

			// Check the final URL after redirection.
			assert.Equal(t, tc.url, redirectedToURL)
		})
	}
}
