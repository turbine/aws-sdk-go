// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package directoryservice

import (
	"github.com/turbine/aws-sdk-go/aws"
	"github.com/turbine/aws-sdk-go/aws/defaults"
	"github.com/turbine/aws-sdk-go/aws/request"
	"github.com/turbine/aws-sdk-go/aws/service"
	"github.com/turbine/aws-sdk-go/aws/service/serviceinfo"
	"github.com/turbine/aws-sdk-go/internal/protocol/jsonrpc"
	"github.com/turbine/aws-sdk-go/internal/signer/v4"
)

// This is the AWS Directory Service API Reference. This guide provides detailed
// information about AWS Directory Service operations, data types, parameters,
// and errors.
type DirectoryService struct {
	*service.Service
}

// Used for custom service initialization logic
var initService func(*service.Service)

// Used for custom request initialization logic
var initRequest func(*request.Request)

// New returns a new DirectoryService client.
func New(config *aws.Config) *DirectoryService {
	service := &service.Service{
		ServiceInfo: serviceinfo.ServiceInfo{
			Config:       defaults.DefaultConfig.Merge(config),
			ServiceName:  "ds",
			APIVersion:   "2015-04-16",
			JSONVersion:  "1.1",
			TargetPrefix: "DirectoryService_20150416",
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

	return &DirectoryService{service}
}

// newRequest creates a new request for a DirectoryService operation and runs any
// custom request initialization.
func (c *DirectoryService) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}
