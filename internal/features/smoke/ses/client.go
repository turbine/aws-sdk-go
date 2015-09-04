//Package ses provides gucumber integration tests suppport.
package ses

import (
	"github.com/turbine/aws-sdk-go/internal/features/shared"
	"github.com/turbine/aws-sdk-go/service/ses"
	. "github.com/lsegal/gucumber"
)

var _ = shared.Imported

func init() {
	Before("@ses", func() {
		World["client"] = ses.New(nil)
	})
}
