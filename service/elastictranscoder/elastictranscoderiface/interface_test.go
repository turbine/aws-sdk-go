// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package elastictranscoderiface_test

import (
	"testing"

	"github.com/turbine/aws-sdk-go/service/elastictranscoder"
	"github.com/turbine/aws-sdk-go/service/elastictranscoder/elastictranscoderiface"
	"github.com/stretchr/testify/assert"
)

func TestInterface(t *testing.T) {
	assert.Implements(t, (*elastictranscoderiface.ElasticTranscoderAPI)(nil), elastictranscoder.New(nil))
}
