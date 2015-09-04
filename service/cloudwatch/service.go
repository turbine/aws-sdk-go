// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package cloudwatch

import (
	"github.com/turbine/aws-sdk-go/aws"
	"github.com/turbine/aws-sdk-go/aws/defaults"
	"github.com/turbine/aws-sdk-go/aws/request"
	"github.com/turbine/aws-sdk-go/aws/service"
	"github.com/turbine/aws-sdk-go/aws/service/serviceinfo"
	"github.com/turbine/aws-sdk-go/internal/protocol/query"
	"github.com/turbine/aws-sdk-go/internal/signer/v4"
)

// This is the Amazon CloudWatch API Reference. This guide provides detailed
// information about Amazon CloudWatch actions, data types, parameters, and
// errors. For detailed information about Amazon CloudWatch features and their
// associated API calls, go to the Amazon CloudWatch Developer Guide (http://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide).
//
// Amazon CloudWatch is a web service that enables you to publish, monitor,
// and manage various metrics, as well as configure alarm actions based on data
// from metrics. For more information about this product go to http://aws.amazon.com/cloudwatch
// (http://aws.amazon.com/cloudwatch).
//
//  For information about the namespace, metric names, and dimensions that
// other Amazon Web Services products use to send metrics to Cloudwatch, go
// to Amazon CloudWatch Metrics, Namespaces, and Dimensions Reference (http://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/CW_Support_For_AWS.html)
// in the Amazon CloudWatch Developer Guide.
//
// Use the following links to get started using the Amazon CloudWatch API Reference:
//
//   Actions (http://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_Operations.html):
// An alphabetical list of all Amazon CloudWatch actions.  Data Types (http://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_Types.html):
// An alphabetical list of all Amazon CloudWatch data types.  Common Parameters
// (http://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/CommonParameters.html):
// Parameters that all Query actions can use.  Common Errors (http://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/CommonErrors.html):
// Client and server errors that all actions can return.  Regions and Endpoints
// (http://docs.aws.amazon.com/general/latest/gr/index.html?rande.html): Itemized
// regions and endpoints for all AWS products.  WSDL Location (http://monitoring.amazonaws.com/doc/2010-08-01/CloudWatch.wsdl):
// http://monitoring.amazonaws.com/doc/2010-08-01/CloudWatch.wsdl  In addition
// to using the Amazon CloudWatch API, you can also use the following SDKs and
// third-party libraries to access Amazon CloudWatch programmatically.
//
//  AWS SDK for Java Documentation (http://aws.amazon.com/documentation/sdkforjava/)
// AWS SDK for .NET Documentation (http://aws.amazon.com/documentation/sdkfornet/)
// AWS SDK for PHP Documentation (http://aws.amazon.com/documentation/sdkforphp/)
// AWS SDK for Ruby Documentation (http://aws.amazon.com/documentation/sdkforruby/)
//  Developers in the AWS developer community also provide their own libraries,
// which you can find at the following AWS developer centers:
//
//  AWS Java Developer Center (http://aws.amazon.com/java/) AWS PHP Developer
// Center (http://aws.amazon.com/php/) AWS Python Developer Center (http://aws.amazon.com/python/)
// AWS Ruby Developer Center (http://aws.amazon.com/ruby/) AWS Windows and .NET
// Developer Center (http://aws.amazon.com/net/)
type CloudWatch struct {
	*service.Service
}

// Used for custom service initialization logic
var initService func(*service.Service)

// Used for custom request initialization logic
var initRequest func(*request.Request)

// New returns a new CloudWatch client.
func New(config *aws.Config) *CloudWatch {
	service := &service.Service{
		ServiceInfo: serviceinfo.ServiceInfo{
			Config:      defaults.DefaultConfig.Merge(config),
			ServiceName: "monitoring",
			APIVersion:  "2010-08-01",
		},
	}
	service.Initialize()

	// Handlers
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Build.PushBack(query.Build)
	service.Handlers.Unmarshal.PushBack(query.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(query.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(query.UnmarshalError)

	// Run custom service initialization if present
	if initService != nil {
		initService(service)
	}

	return &CloudWatch{service}
}

// newRequest creates a new request for a CloudWatch operation and runs any
// custom request initialization.
func (c *CloudWatch) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}
