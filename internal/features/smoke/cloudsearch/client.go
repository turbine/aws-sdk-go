//Package cloudsearch provides gucumber integration tests suppport.
package cloudsearch

import (
	"github.com/turbine/aws-sdk-go/internal/features/shared"
	"github.com/turbine/aws-sdk-go/service/cloudsearch"
	. "github.com/lsegal/gucumber"
)

var _ = shared.Imported

func init() {
	Before("@cloudsearch", func() {
		World["client"] = cloudsearch.New(nil)
	})
}
