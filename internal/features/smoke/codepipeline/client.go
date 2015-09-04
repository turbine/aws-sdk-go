//Package codepipeline provides gucumber integration tests suppport.
package codepipeline

import (
	"github.com/turbine/aws-sdk-go/internal/features/shared"
	"github.com/turbine/aws-sdk-go/service/codepipeline"
	. "github.com/lsegal/gucumber"
)

var _ = shared.Imported

func init() {
	Before("@codepipeline", func() {
		World["client"] = codepipeline.New(nil)
	})
}
