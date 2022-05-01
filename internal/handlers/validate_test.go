package handlers_test

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/MalukiMuthusi/pulseid/internal/handlers"
	"github.com/MalukiMuthusi/pulseid/internal/models"
	"github.com/MalukiMuthusi/pulseid/internal/store"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestValidate(t *testing.T) {

	store := store.NewMockStore()

	type test struct {
		Name     string
		EndPoint string
		Status   int
		Resp     interface{}
	}

	happyCaseResp := map[string]interface{}{"message": "not implemented"}

	tests := []test{
		{
			Name:     "happy case",
			EndPoint: "/validate/some_unique_token",
			Status:   http.StatusNotImplemented,
			Resp:     happyCaseResp,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {

			req, err := http.NewRequest(http.MethodGet, test.EndPoint, nil)
			if err != nil {
				assert.Fail(t, "failed to create a new request in test")
			}

			w := httptest.NewRecorder()

			router := handlers.SetUpRouter(store)

			router.ServeHTTP(w, req)

			assert.Equal(t, test.Status, w.Code)

			var res interface{}

			b, err := ioutil.ReadAll(w.Body)
			if err != nil {
				assert.Fail(t, "failed to read response")
			}

			err = json.Unmarshal(b, &res)
			if err != nil {
				assert.Fail(t, "failed to unMarshal response")
			}

			assert.EqualValues(t, test.Resp, res)
		})
	}
}

func TestCheckValidity(t *testing.T) {

	validToken := models.Token{
		Model: gorm.Model{
			CreatedAt: time.Now(),
		},
	}

	expiredToken := models.Token{
		Model: gorm.Model{
			CreatedAt: time.Now().AddDate(-1, 1, 8),
		},
	}

	recalledToken := models.Token{
		Model: gorm.Model{
			CreatedAt: time.Now().AddDate(1, 1, 8),
		},
		Recalled: true,
	}

	type test struct {
		name     string
		token    models.Token
		validity bool
	}

	tests := []test{
		{
			name:     "happy case",
			token:    validToken,
			validity: true,
		},
		{
			name:     "expired token",
			token:    expiredToken,
			validity: false,
		},
		{
			name:     "recalled token",
			token:    recalledToken,
			validity: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(tt *testing.T) {
			validity := test.token.CheckValidity()

			assert.Equal(tt, test.validity, validity)
		})
	}
}
