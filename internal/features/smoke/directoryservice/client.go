//Package directoryservice provides gucumber integration tests suppport.
package directoryservice

import (
	"github.com/turbine/aws-sdk-go/internal/features/shared"
	"github.com/turbine/aws-sdk-go/service/directoryservice"
	. "github.com/lsegal/gucumber"
)

var _ = shared.Imported

func init() {
	Before("@directoryservice", func() {
		World["client"] = directoryservice.New(nil)
	})
}
