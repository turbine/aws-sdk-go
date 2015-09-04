//Package route53 provides gucumber integration tests suppport.
package route53

import (
	"github.com/turbine/aws-sdk-go/internal/features/shared"
	"github.com/turbine/aws-sdk-go/service/route53"
	. "github.com/lsegal/gucumber"
)

var _ = shared.Imported

func init() {
	Before("@route53", func() {
		World["client"] = route53.New(nil)
	})
}
