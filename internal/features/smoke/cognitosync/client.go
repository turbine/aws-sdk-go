//Package cognitosync provides gucumber integration tests suppport.
package cognitosync

import (
	"github.com/turbine/aws-sdk-go/internal/features/shared"
	"github.com/turbine/aws-sdk-go/service/cognitosync"
	. "github.com/lsegal/gucumber"
)

var _ = shared.Imported

func init() {
	Before("@cognitosync", func() {
		World["client"] = cognitosync.New(nil)
	})
}
