package machinelearning_test

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/turbine/aws-sdk-go/aws"
	"github.com/turbine/aws-sdk-go/aws/request"
	"github.com/turbine/aws-sdk-go/internal/test/unit"
	"github.com/turbine/aws-sdk-go/service/machinelearning"
	"github.com/stretchr/testify/assert"
)

var _ = unit.Imported

func TestPredictEndpoint(t *testing.T) {
	ml := machinelearning.New(nil)
	ml.Handlers.Send.Clear()
	ml.Handlers.Send.PushBack(func(r *request.Request) {
		r.HTTPResponse = &http.Response{
			StatusCode: 200,
			Header:     http.Header{},
			Body:       ioutil.NopCloser(bytes.NewReader([]byte("{}"))),
		}
	})

	req, _ := ml.PredictRequest(&machinelearning.PredictInput{
		PredictEndpoint: aws.String("https://localhost/endpoint"),
		MLModelId:       aws.String("id"),
		Record:          map[string]*string{},
	})
	err := req.Send()

	assert.Nil(t, err)
	assert.Equal(t, "https://localhost/endpoint", req.HTTPRequest.URL.String())
}
