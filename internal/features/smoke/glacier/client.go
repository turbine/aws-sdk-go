//Package glacier provides gucumber integration tests suppport.
package glacier

import (
	"github.com/turbine/aws-sdk-go/internal/features/shared"
	"github.com/turbine/aws-sdk-go/service/glacier"
	. "github.com/lsegal/gucumber"
)

var _ = shared.Imported

func init() {
	Before("@glacier", func() {
		World["client"] = glacier.New(nil)
	})
}
