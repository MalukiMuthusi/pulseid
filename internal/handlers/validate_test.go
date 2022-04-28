package handlers_test

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/MalukiMuthusi/pulseid/internal/handlers"
	"github.com/MalukiMuthusi/pulseid/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestValidate(t *testing.T) {

	type test struct {
		Name     string
		EndPoint string
		Status   int
		Resp     interface{}
	}

	happyCaseResp := map[string]interface{}{"message": "not implemented"}

	tokenNotProvidedResp := models.BasicError{
		Code:    "INVALID_TOKEN_PARAMETER",
		Message: "provide a valid token parameter",
	}

	tests := []test{
		{
			Name:     "happy case",
			EndPoint: "/validate/some_unique_token",
			Status:   http.StatusNotImplemented,
			Resp:     happyCaseResp,
		},
		{
			Name:     "token not provided",
			EndPoint: "/validate/",
			Status:   http.StatusUnprocessableEntity,
			Resp:     tokenNotProvidedResp,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {

			req, err := http.NewRequest(http.MethodGet, test.EndPoint, nil)
			if err != nil {
				assert.Fail(t, "failed to create a new request in test")
			}

			w := httptest.NewRecorder()

			router := handlers.SetUpRouter()

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
