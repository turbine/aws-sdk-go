// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package dynamodbstreamsiface_test

import (
	"testing"

	"github.com/turbine/aws-sdk-go/service/dynamodbstreams"
	"github.com/turbine/aws-sdk-go/service/dynamodbstreams/dynamodbstreamsiface"
	"github.com/stretchr/testify/assert"
)

func TestInterface(t *testing.T) {
	assert.Implements(t, (*dynamodbstreamsiface.DynamoDBStreamsAPI)(nil), dynamodbstreams.New(nil))
}
