package handlers_test

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/MalukiMuthusi/pulseid/internal/handlers"
	"github.com/stretchr/testify/assert"
)

func TestInactive(t *testing.T) {
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
			EndPoint: "/inactive",
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
