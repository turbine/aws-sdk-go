// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package route53iface_test

import (
	"testing"

	"github.com/turbine/aws-sdk-go/service/route53"
	"github.com/turbine/aws-sdk-go/service/route53/route53iface"
	"github.com/stretchr/testify/assert"
)

func TestInterface(t *testing.T) {
	assert.Implements(t, (*route53iface.Route53API)(nil), route53.New(nil))
}
