//Package efs provides gucumber integration tests suppport.
package efs

import (
	"github.com/turbine/aws-sdk-go/aws"
	"github.com/turbine/aws-sdk-go/internal/features/shared"
	"github.com/turbine/aws-sdk-go/service/efs"
	. "github.com/lsegal/gucumber"
)

var _ = shared.Imported

func init() {
	Before("@efs", func() {
		// FIXME remove custom region
		World["client"] = efs.New(aws.NewConfig().WithRegion("us-west-2"))
	})
}
