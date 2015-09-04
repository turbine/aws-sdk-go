// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package codecommit

import (
	"github.com/turbine/aws-sdk-go/aws"
	"github.com/turbine/aws-sdk-go/aws/defaults"
	"github.com/turbine/aws-sdk-go/aws/request"
	"github.com/turbine/aws-sdk-go/aws/service"
	"github.com/turbine/aws-sdk-go/aws/service/serviceinfo"
	"github.com/turbine/aws-sdk-go/internal/protocol/jsonrpc"
	"github.com/turbine/aws-sdk-go/internal/signer/v4"
)

// This is the AWS CodeCommit API Reference. This reference provides descriptions
// of the AWS CodeCommit API.
//
// You can use the AWS CodeCommit API to work with the following objects:
//
//  Repositories Branches Commits  For information about how to use AWS CodeCommit,
// see the AWS CodeCommit User Guide.
type CodeCommit struct {
	*service.Service
}

// Used for custom service initialization logic
var initService func(*service.Service)

// Used for custom request initialization logic
var initRequest func(*request.Request)

// New returns a new CodeCommit client.
func New(config *aws.Config) *CodeCommit {
	service := &service.Service{
		ServiceInfo: serviceinfo.ServiceInfo{
			Config:       defaults.DefaultConfig.Merge(config),
			ServiceName:  "codecommit",
			APIVersion:   "2015-04-13",
			JSONVersion:  "1.1",
			TargetPrefix: "CodeCommit_20150413",
		},
	}
	service.Initialize()

	// Handlers
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Build.PushBack(jsonrpc.Build)
	service.Handlers.Unmarshal.PushBack(jsonrpc.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(jsonrpc.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(jsonrpc.UnmarshalError)

	// Run custom service initialization if present
	if initService != nil {
		initService(service)
	}

	return &CodeCommit{service}
}

// newRequest creates a new request for a CodeCommit operation and runs any
// custom request initialization.
func (c *CodeCommit) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}
