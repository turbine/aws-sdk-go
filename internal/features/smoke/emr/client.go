//Package emr provides gucumber integration tests suppport.
package emr

import (
	"github.com/turbine/aws-sdk-go/internal/features/shared"
	"github.com/turbine/aws-sdk-go/service/emr"
	. "github.com/lsegal/gucumber"
)

var _ = shared.Imported

func init() {
	Before("@emr", func() {
		World["client"] = emr.New(nil)
	})
}
