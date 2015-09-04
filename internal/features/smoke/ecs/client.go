//Package ecs provides gucumber integration tests suppport.
package ecs

import (
	"github.com/turbine/aws-sdk-go/aws"
	"github.com/turbine/aws-sdk-go/internal/features/shared"
	"github.com/turbine/aws-sdk-go/service/ecs"
	. "github.com/lsegal/gucumber"
)

var _ = shared.Imported

func init() {
	Before("@ecs", func() {
		// FIXME remove custom region
		World["client"] = ecs.New(aws.NewConfig().WithRegion("us-west-2"))
	})
}
