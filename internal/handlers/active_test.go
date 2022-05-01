package handlers_test

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/MalukiMuthusi/pulseid/internal/handlers"
	"github.com/MalukiMuthusi/pulseid/internal/models"
	"github.com/MalukiMuthusi/pulseid/internal/store"
	"github.com/stretchr/testify/assert"
)

func TestActive(t *testing.T) {

	store := store.NewMockStore()

	type test struct {
		Name     string
		EndPoint string
		Status   int
	}

	tests := []test{
		{
			Name:     "happy case",
			EndPoint: "/active",
			Status:   http.StatusOK,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(tt *testing.T) {

			req, err := http.NewRequest(http.MethodGet, test.EndPoint, nil)
			if err != nil {
				assert.Fail(t, "failed to create a new request in test")
			}

			w := httptest.NewRecorder()

			// Add Authorization Header

			credentials := base64.StdEncoding.EncodeToString([]byte("username:password"))

			req.Header.Add("Authorization", fmt.Sprintf("Basic %s", credentials))

			router := handlers.SetUpRouter(store)

			router.ServeHTTP(w, req)

			assert.Equal(tt, test.Status, w.Code)

			var res []*models.Token

			b, err := ioutil.ReadAll(w.Body)
			if err != nil {
				assert.Fail(t, "failed to read response")
			}

			err = json.Unmarshal(b, &res)
			if err != nil {
				assert.Fail(t, "failed to unMarshal response")
			}

			assert.Equal(tt, true, len(res) > 0)

			for _, v := range res {
				assert.Equal(tt, v.CheckValidity(), true)
			}

		})
	}
}
