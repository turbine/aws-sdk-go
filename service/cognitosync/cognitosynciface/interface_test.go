// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package cognitosynciface_test

import (
	"testing"

	"github.com/turbine/aws-sdk-go/service/cognitosync"
	"github.com/turbine/aws-sdk-go/service/cognitosync/cognitosynciface"
	"github.com/stretchr/testify/assert"
)

func TestInterface(t *testing.T) {
	assert.Implements(t, (*cognitosynciface.CognitoSyncAPI)(nil), cognitosync.New(nil))
}
