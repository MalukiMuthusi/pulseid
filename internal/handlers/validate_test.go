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
	"github.com/MalukiMuthusi/pulseid/internal/store/mock"
	"github.com/stretchr/testify/assert"
)

func TestValidate(t *testing.T) {

	store := mock.NewStore()

	type test struct {
		Name     string
		EndPoint string
		Status   int
		Resp     models.TokenValidity
	}

	happyCaseResp := models.TokenValidity{
		Validity: true,
	}

	tests := []test{
		{
			Name:     "happy case",
			EndPoint: "/validate/some_unique_token",
			Status:   http.StatusOK,
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

			router := handlers.SetUpRouter(store, DebugPrintRoute)

			router.ServeHTTP(w, req)

			assert.Equal(t, test.Status, w.Code)

			var res models.TokenValidity

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

	validToken, err := models.NewToken()
	assert.Nil(t, err, "failed to generate a token")

	expiredToken, err := models.NewToken()
	expiredToken.Expiry = time.Now().AddDate(0, 0, -8)
	assert.Nil(t, err, "failed to generate a token")

	recalledToken, err := models.NewToken()
	assert.Nil(t, err, "failed to generate a token")

	recalledToken.Recalled = true

	type test struct {
		name     string
		token    *models.Token
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
