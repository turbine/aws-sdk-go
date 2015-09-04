//Package kms provides gucumber integration tests suppport.
package kms

import (
	"github.com/turbine/aws-sdk-go/internal/features/shared"
	"github.com/turbine/aws-sdk-go/service/kms"
	. "github.com/lsegal/gucumber"
)

var _ = shared.Imported

func init() {
	Before("@kms", func() {
		World["client"] = kms.New(nil)
	})
}
