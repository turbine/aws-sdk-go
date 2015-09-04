//Package configservice provides gucumber integration tests suppport.
package configservice

import (
	"github.com/turbine/aws-sdk-go/internal/features/shared"
	"github.com/turbine/aws-sdk-go/service/configservice"
	. "github.com/lsegal/gucumber"
)

var _ = shared.Imported

func init() {
	Before("@configservice", func() {
		World["client"] = configservice.New(nil)
	})
}
