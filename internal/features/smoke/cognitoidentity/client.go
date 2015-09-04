//Package cognitoidentity provides gucumber integration tests suppport.
package cognitoidentity

import (
	"github.com/turbine/aws-sdk-go/internal/features/shared"
	"github.com/turbine/aws-sdk-go/service/cognitoidentity"
	. "github.com/lsegal/gucumber"
)

var _ = shared.Imported

func init() {
	Before("@cognitoidentity", func() {
		World["client"] = cognitoidentity.New(nil)
	})
}
