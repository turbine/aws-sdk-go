//Package sqs provides gucumber integration tests suppport.
package sqs

import (
	"github.com/turbine/aws-sdk-go/internal/features/shared"
	"github.com/turbine/aws-sdk-go/service/sqs"
	. "github.com/lsegal/gucumber"
)

var _ = shared.Imported

func init() {
	Before("@sqs", func() {
		World["client"] = sqs.New(nil)
	})
}
