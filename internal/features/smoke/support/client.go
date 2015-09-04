//Package support provides gucumber integration tests suppport.
package support

import (
	"github.com/turbine/aws-sdk-go/internal/features/shared"
	"github.com/turbine/aws-sdk-go/service/support"
	. "github.com/lsegal/gucumber"
)

var _ = shared.Imported

func init() {
	Before("@support", func() {
		World["client"] = support.New(nil)
	})
}
