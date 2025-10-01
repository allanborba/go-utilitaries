package httprequestmoker

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/allanborba/go-utilitaries/asserts"
	jsonutils "github.com/allanborba/go-utilitaries/json_utils"
)

func TestRequestMocker(t *testing.T) {
	payload := &testDto{Msg: "test"}

	headers := map[string]string{"Authorization": "Bearer 123"}
	result := NewRequestMocker[*testDto](jsonutils.ToJsonBytes(payload), &testController{}).AddHeader(headers).RequestWithResponse()

	asserts.DeepEqual(t, payload, result)
}

type testDto struct {
	Msg string `json:"msg"`
}

type testController struct{}

func (this *testController) Execute(w http.ResponseWriter, r *http.Request) {
	input := &testDto{}
	token := r.Header.Get("Authorization")
	if token != "Bearer 123" {
		panic("token invalid")
	}

	json.NewDecoder(r.Body).Decode(input)

	json.NewEncoder(w).Encode(input)
}
