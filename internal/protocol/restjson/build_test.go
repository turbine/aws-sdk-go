package restjson_test

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/defaults"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/service"
	"github.com/aws/aws-sdk-go/aws/service/serviceinfo"
	"github.com/aws/aws-sdk-go/internal/protocol/restjson"
	"github.com/aws/aws-sdk-go/internal/protocol/xml/xmlutil"
	"github.com/aws/aws-sdk-go/internal/signer/v4"
	"github.com/aws/aws-sdk-go/internal/util"
	"github.com/stretchr/testify/assert"
)

var _ bytes.Buffer // always import bytes
var _ http.Request
var _ json.Marshaler
var _ time.Time
var _ xmlutil.XMLNode
var _ xml.Attr
var _ = ioutil.Discard
var _ = util.Trim("")
var _ = url.Values{}
var _ = io.EOF

type InputService1ProtocolTest struct {
	*service.Service
}

// New returns a new InputService1ProtocolTest client.
func NewInputService1ProtocolTest(config *aws.Config) *InputService1ProtocolTest {
	service := &service.Service{
		ServiceInfo: serviceinfo.ServiceInfo{
			Config:      defaults.DefaultConfig.Merge(config),
			ServiceName: "inputservice1protocoltest",
			APIVersion:  "2014-01-01",
		},
	}
	service.Initialize()

	// Handlers
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Build.PushBack(restjson.Build)
	service.Handlers.Unmarshal.PushBack(restjson.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(restjson.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(restjson.UnmarshalError)

	return &InputService1ProtocolTest{service}
}

// newRequest creates a new request for a InputService1ProtocolTest operation and runs any
// custom request initialization.
func (c *InputService1ProtocolTest) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	return req
}

const opInputService1TestCaseOperation1 = "OperationName"

// InputService1TestCaseOperation1Request generates a request for the InputService1TestCaseOperation1 operation.
func (c *InputService1ProtocolTest) InputService1TestCaseOperation1Request(input *InputService1TestShapeInputService1TestCaseOperation1Input) (req *request.Request, output *InputService1TestShapeInputService1TestCaseOperation1Output) {
	op := &request.Operation{
		Name:       opInputService1TestCaseOperation1,
		HTTPMethod: "GET",
		HTTPPath:   "/2014-01-01/jobsByPipeline/{PipelineId}",
	}

	if input == nil {
		input = &InputService1TestShapeInputService1TestCaseOperation1Input{}
	}

	req = c.newRequest(op, input, output)
	output = &InputService1TestShapeInputService1TestCaseOperation1Output{}
	req.Data = output
	return
}

func (c *InputService1ProtocolTest) InputService1TestCaseOperation1(input *InputService1TestShapeInputService1TestCaseOperation1Input) (*InputService1TestShapeInputService1TestCaseOperation1Output, error) {
	req, out := c.InputService1TestCaseOperation1Request(input)
	err := req.Send()
	return out, err
}

type InputService1TestShapeInputService1TestCaseOperation1Input struct {
	PipelineId *string `location:"uri" type:"string"`

	metadataInputService1TestShapeInputService1TestCaseOperation1Input `json:"-" xml:"-"`
}

type metadataInputService1TestShapeInputService1TestCaseOperation1Input struct {
	SDKShapeTraits bool `type:"structure"`
}

type InputService1TestShapeInputService1TestCaseOperation1Output struct {
	metadataInputService1TestShapeInputService1TestCaseOperation1Output `json:"-" xml:"-"`
}

type metadataInputService1TestShapeInputService1TestCaseOperation1Output struct {
	SDKShapeTraits bool `type:"structure"`
}

type InputService2ProtocolTest struct {
	*service.Service
}

// New returns a new InputService2ProtocolTest client.
func NewInputService2ProtocolTest(config *aws.Config) *InputService2ProtocolTest {
	service := &service.Service{
		ServiceInfo: serviceinfo.ServiceInfo{
			Config:      defaults.DefaultConfig.Merge(config),
			ServiceName: "inputservice2protocoltest",
			APIVersion:  "2014-01-01",
		},
	}
	service.Initialize()

	// Handlers
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Build.PushBack(restjson.Build)
	service.Handlers.Unmarshal.PushBack(restjson.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(restjson.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(restjson.UnmarshalError)

	return &InputService2ProtocolTest{service}
}

// newRequest creates a new request for a InputService2ProtocolTest operation and runs any
// custom request initialization.
func (c *InputService2ProtocolTest) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	return req
}

const opInputService2TestCaseOperation1 = "OperationName"

// InputService2TestCaseOperation1Request generates a request for the InputService2TestCaseOperation1 operation.
func (c *InputService2ProtocolTest) InputService2TestCaseOperation1Request(input *InputService2TestShapeInputService2TestCaseOperation1Input) (req *request.Request, output *InputService2TestShapeInputService2TestCaseOperation1Output) {
	op := &request.Operation{
		Name:       opInputService2TestCaseOperation1,
		HTTPMethod: "GET",
		HTTPPath:   "/2014-01-01/jobsByPipeline/{PipelineId}",
	}

	if input == nil {
		input = &InputService2TestShapeInputService2TestCaseOperation1Input{}
	}

	req = c.newRequest(op, input, output)
	output = &InputService2TestShapeInputService2TestCaseOperation1Output{}
	req.Data = output
	return
}

func (c *InputService2ProtocolTest) InputService2TestCaseOperation1(input *InputService2TestShapeInputService2TestCaseOperation1Input) (*InputService2TestShapeInputService2TestCaseOperation1Output, error) {
	req, out := c.InputService2TestCaseOperation1Request(input)
	err := req.Send()
	return out, err
}

type InputService2TestShapeInputService2TestCaseOperation1Input struct {
	Foo *string `location:"uri" locationName:"PipelineId" type:"string"`

	metadataInputService2TestShapeInputService2TestCaseOperation1Input `json:"-" xml:"-"`
}

type metadataInputService2TestShapeInputService2TestCaseOperation1Input struct {
	SDKShapeTraits bool `type:"structure"`
}

type InputService2TestShapeInputService2TestCaseOperation1Output struct {
	metadataInputService2TestShapeInputService2TestCaseOperation1Output `json:"-" xml:"-"`
}

type metadataInputService2TestShapeInputService2TestCaseOperation1Output struct {
	SDKShapeTraits bool `type:"structure"`
}

type InputService3ProtocolTest struct {
	*service.Service
}

// New returns a new InputService3ProtocolTest client.
func NewInputService3ProtocolTest(config *aws.Config) *InputService3ProtocolTest {
	service := &service.Service{
		ServiceInfo: serviceinfo.ServiceInfo{
			Config:      defaults.DefaultConfig.Merge(config),
			ServiceName: "inputservice3protocoltest",
			APIVersion:  "2014-01-01",
		},
	}
	service.Initialize()

	// Handlers
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Build.PushBack(restjson.Build)
	service.Handlers.Unmarshal.PushBack(restjson.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(restjson.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(restjson.UnmarshalError)

	return &InputService3ProtocolTest{service}
}

// newRequest creates a new request for a InputService3ProtocolTest operation and runs any
// custom request initialization.
func (c *InputService3ProtocolTest) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	return req
}

const opInputService3TestCaseOperation1 = "OperationName"

// InputService3TestCaseOperation1Request generates a request for the InputService3TestCaseOperation1 operation.
func (c *InputService3ProtocolTest) InputService3TestCaseOperation1Request(input *InputService3TestShapeInputService3TestCaseOperation1Input) (req *request.Request, output *InputService3TestShapeInputService3TestCaseOperation1Output) {
	op := &request.Operation{
		Name:       opInputService3TestCaseOperation1,
		HTTPMethod: "GET",
		HTTPPath:   "/2014-01-01/jobsByPipeline/{PipelineId}",
	}

	if input == nil {
		input = &InputService3TestShapeInputService3TestCaseOperation1Input{}
	}

	req = c.newRequest(op, input, output)
	output = &InputService3TestShapeInputService3TestCaseOperation1Output{}
	req.Data = output
	return
}

func (c *InputService3ProtocolTest) InputService3TestCaseOperation1(input *InputService3TestShapeInputService3TestCaseOperation1Input) (*InputService3TestShapeInputService3TestCaseOperation1Output, error) {
	req, out := c.InputService3TestCaseOperation1Request(input)
	err := req.Send()
	return out, err
}

type InputService3TestShapeInputService3TestCaseOperation1Input struct {
	Ascending *string `location:"querystring" locationName:"Ascending" type:"string"`

	PageToken *string `location:"querystring" locationName:"PageToken" type:"string"`

	PipelineId *string `location:"uri" locationName:"PipelineId" type:"string"`

	metadataInputService3TestShapeInputService3TestCaseOperation1Input `json:"-" xml:"-"`
}

type metadataInputService3TestShapeInputService3TestCaseOperation1Input struct {
	SDKShapeTraits bool `type:"structure"`
}

type InputService3TestShapeInputService3TestCaseOperation1Output struct {
	metadataInputService3TestShapeInputService3TestCaseOperation1Output `json:"-" xml:"-"`
}

type metadataInputService3TestShapeInputService3TestCaseOperation1Output struct {
	SDKShapeTraits bool `type:"structure"`
}

type InputService4ProtocolTest struct {
	*service.Service
}

// New returns a new InputService4ProtocolTest client.
func NewInputService4ProtocolTest(config *aws.Config) *InputService4ProtocolTest {
	service := &service.Service{
		ServiceInfo: serviceinfo.ServiceInfo{
			Config:      defaults.DefaultConfig.Merge(config),
			ServiceName: "inputservice4protocoltest",
			APIVersion:  "2014-01-01",
		},
	}
	service.Initialize()

	// Handlers
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Build.PushBack(restjson.Build)
	service.Handlers.Unmarshal.PushBack(restjson.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(restjson.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(restjson.UnmarshalError)

	return &InputService4ProtocolTest{service}
}

// newRequest creates a new request for a InputService4ProtocolTest operation and runs any
// custom request initialization.
func (c *InputService4ProtocolTest) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	return req
}

const opInputService4TestCaseOperation1 = "OperationName"

// InputService4TestCaseOperation1Request generates a request for the InputService4TestCaseOperation1 operation.
func (c *InputService4ProtocolTest) InputService4TestCaseOperation1Request(input *InputService4TestShapeInputService4TestCaseOperation1Input) (req *request.Request, output *InputService4TestShapeInputService4TestCaseOperation1Output) {
	op := &request.Operation{
		Name:       opInputService4TestCaseOperation1,
		HTTPMethod: "POST",
		HTTPPath:   "/2014-01-01/jobsByPipeline/{PipelineId}",
	}

	if input == nil {
		input = &InputService4TestShapeInputService4TestCaseOperation1Input{}
	}

	req = c.newRequest(op, input, output)
	output = &InputService4TestShapeInputService4TestCaseOperation1Output{}
	req.Data = output
	return
}

func (c *InputService4ProtocolTest) InputService4TestCaseOperation1(input *InputService4TestShapeInputService4TestCaseOperation1Input) (*InputService4TestShapeInputService4TestCaseOperation1Output, error) {
	req, out := c.InputService4TestCaseOperation1Request(input)
	err := req.Send()
	return out, err
}

type InputService4TestShapeInputService4TestCaseOperation1Input struct {
	Ascending *string `location:"querystring" locationName:"Ascending" type:"string"`

	Config *InputService4TestShapeStructType `type:"structure"`

	PageToken *string `location:"querystring" locationName:"PageToken" type:"string"`

	PipelineId *string `location:"uri" locationName:"PipelineId" type:"string"`

	metadataInputService4TestShapeInputService4TestCaseOperation1Input `json:"-" xml:"-"`
}

type metadataInputService4TestShapeInputService4TestCaseOperation1Input struct {
	SDKShapeTraits bool `type:"structure"`
}

type InputService4TestShapeInputService4TestCaseOperation1Output struct {
	metadataInputService4TestShapeInputService4TestCaseOperation1Output `json:"-" xml:"-"`
}

type metadataInputService4TestShapeInputService4TestCaseOperation1Output struct {
	SDKShapeTraits bool `type:"structure"`
}

type InputService4TestShapeStructType struct {
	A *string `type:"string"`

	B *string `type:"string"`

	metadataInputService4TestShapeStructType `json:"-" xml:"-"`
}

type metadataInputService4TestShapeStructType struct {
	SDKShapeTraits bool `type:"structure"`
}

type InputService5ProtocolTest struct {
	*service.Service
}

// New returns a new InputService5ProtocolTest client.
func NewInputService5ProtocolTest(config *aws.Config) *InputService5ProtocolTest {
	service := &service.Service{
		ServiceInfo: serviceinfo.ServiceInfo{
			Config:      defaults.DefaultConfig.Merge(config),
			ServiceName: "inputservice5protocoltest",
			APIVersion:  "2014-01-01",
		},
	}
	service.Initialize()

	// Handlers
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Build.PushBack(restjson.Build)
	service.Handlers.Unmarshal.PushBack(restjson.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(restjson.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(restjson.UnmarshalError)

	return &InputService5ProtocolTest{service}
}

// newRequest creates a new request for a InputService5ProtocolTest operation and runs any
// custom request initialization.
func (c *InputService5ProtocolTest) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	return req
}

const opInputService5TestCaseOperation1 = "OperationName"

// InputService5TestCaseOperation1Request generates a request for the InputService5TestCaseOperation1 operation.
func (c *InputService5ProtocolTest) InputService5TestCaseOperation1Request(input *InputService5TestShapeInputService5TestCaseOperation1Input) (req *request.Request, output *InputService5TestShapeInputService5TestCaseOperation1Output) {
	op := &request.Operation{
		Name:       opInputService5TestCaseOperation1,
		HTTPMethod: "POST",
		HTTPPath:   "/2014-01-01/jobsByPipeline/{PipelineId}",
	}

	if input == nil {
		input = &InputService5TestShapeInputService5TestCaseOperation1Input{}
	}

	req = c.newRequest(op, input, output)
	output = &InputService5TestShapeInputService5TestCaseOperation1Output{}
	req.Data = output
	return
}

func (c *InputService5ProtocolTest) InputService5TestCaseOperation1(input *InputService5TestShapeInputService5TestCaseOperation1Input) (*InputService5TestShapeInputService5TestCaseOperation1Output, error) {
	req, out := c.InputService5TestCaseOperation1Request(input)
	err := req.Send()
	return out, err
}

type InputService5TestShapeInputService5TestCaseOperation1Input struct {
	Ascending *string `location:"querystring" locationName:"Ascending" type:"string"`

	Checksum *string `location:"header" locationName:"x-amz-checksum" type:"string"`

	Config *InputService5TestShapeStructType `type:"structure"`

	PageToken *string `location:"querystring" locationName:"PageToken" type:"string"`

	PipelineId *string `location:"uri" locationName:"PipelineId" type:"string"`

	metadataInputService5TestShapeInputService5TestCaseOperation1Input `json:"-" xml:"-"`
}

type metadataInputService5TestShapeInputService5TestCaseOperation1Input struct {
	SDKShapeTraits bool `type:"structure"`
}

type InputService5TestShapeInputService5TestCaseOperation1Output struct {
	metadataInputService5TestShapeInputService5TestCaseOperation1Output `json:"-" xml:"-"`
}

type metadataInputService5TestShapeInputService5TestCaseOperation1Output struct {
	SDKShapeTraits bool `type:"structure"`
}

type InputService5TestShapeStructType struct {
	A *string `type:"string"`

	B *string `type:"string"`

	metadataInputService5TestShapeStructType `json:"-" xml:"-"`
}

type metadataInputService5TestShapeStructType struct {
	SDKShapeTraits bool `type:"structure"`
}

type InputService6ProtocolTest struct {
	*service.Service
}

// New returns a new InputService6ProtocolTest client.
func NewInputService6ProtocolTest(config *aws.Config) *InputService6ProtocolTest {
	service := &service.Service{
		ServiceInfo: serviceinfo.ServiceInfo{
			Config:      defaults.DefaultConfig.Merge(config),
			ServiceName: "inputservice6protocoltest",
			APIVersion:  "2014-01-01",
		},
	}
	service.Initialize()

	// Handlers
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Build.PushBack(restjson.Build)
	service.Handlers.Unmarshal.PushBack(restjson.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(restjson.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(restjson.UnmarshalError)

	return &InputService6ProtocolTest{service}
}

// newRequest creates a new request for a InputService6ProtocolTest operation and runs any
// custom request initialization.
func (c *InputService6ProtocolTest) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	return req
}

const opInputService6TestCaseOperation1 = "OperationName"

// InputService6TestCaseOperation1Request generates a request for the InputService6TestCaseOperation1 operation.
func (c *InputService6ProtocolTest) InputService6TestCaseOperation1Request(input *InputService6TestShapeInputService6TestCaseOperation1Input) (req *request.Request, output *InputService6TestShapeInputService6TestCaseOperation1Output) {
	op := &request.Operation{
		Name:       opInputService6TestCaseOperation1,
		HTTPMethod: "POST",
		HTTPPath:   "/2014-01-01/vaults/{vaultName}/archives",
	}

	if input == nil {
		input = &InputService6TestShapeInputService6TestCaseOperation1Input{}
	}

	req = c.newRequest(op, input, output)
	output = &InputService6TestShapeInputService6TestCaseOperation1Output{}
	req.Data = output
	return
}

func (c *InputService6ProtocolTest) InputService6TestCaseOperation1(input *InputService6TestShapeInputService6TestCaseOperation1Input) (*InputService6TestShapeInputService6TestCaseOperation1Output, error) {
	req, out := c.InputService6TestCaseOperation1Request(input)
	err := req.Send()
	return out, err
}

type InputService6TestShapeInputService6TestCaseOperation1Input struct {
	Body io.ReadSeeker `locationName:"body" type:"blob"`

	Checksum *string `location:"header" locationName:"x-amz-sha256-tree-hash" type:"string"`

	VaultName *string `location:"uri" locationName:"vaultName" type:"string" required:"true"`

	metadataInputService6TestShapeInputService6TestCaseOperation1Input `json:"-" xml:"-"`
}

type metadataInputService6TestShapeInputService6TestCaseOperation1Input struct {
	SDKShapeTraits bool `type:"structure" payload:"Body"`
}

type InputService6TestShapeInputService6TestCaseOperation1Output struct {
	metadataInputService6TestShapeInputService6TestCaseOperation1Output `json:"-" xml:"-"`
}

type metadataInputService6TestShapeInputService6TestCaseOperation1Output struct {
	SDKShapeTraits bool `type:"structure"`
}

type InputService7ProtocolTest struct {
	*service.Service
}

// New returns a new InputService7ProtocolTest client.
func NewInputService7ProtocolTest(config *aws.Config) *InputService7ProtocolTest {
	service := &service.Service{
		ServiceInfo: serviceinfo.ServiceInfo{
			Config:      defaults.DefaultConfig.Merge(config),
			ServiceName: "inputservice7protocoltest",
			APIVersion:  "2014-01-01",
		},
	}
	service.Initialize()

	// Handlers
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Build.PushBack(restjson.Build)
	service.Handlers.Unmarshal.PushBack(restjson.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(restjson.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(restjson.UnmarshalError)

	return &InputService7ProtocolTest{service}
}

// newRequest creates a new request for a InputService7ProtocolTest operation and runs any
// custom request initialization.
func (c *InputService7ProtocolTest) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	return req
}

const opInputService7TestCaseOperation1 = "OperationName"

// InputService7TestCaseOperation1Request generates a request for the InputService7TestCaseOperation1 operation.
func (c *InputService7ProtocolTest) InputService7TestCaseOperation1Request(input *InputService7TestShapeInputService7TestCaseOperation1Input) (req *request.Request, output *InputService7TestShapeInputService7TestCaseOperation1Output) {
	op := &request.Operation{
		Name:       opInputService7TestCaseOperation1,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &InputService7TestShapeInputService7TestCaseOperation1Input{}
	}

	req = c.newRequest(op, input, output)
	output = &InputService7TestShapeInputService7TestCaseOperation1Output{}
	req.Data = output
	return
}

func (c *InputService7ProtocolTest) InputService7TestCaseOperation1(input *InputService7TestShapeInputService7TestCaseOperation1Input) (*InputService7TestShapeInputService7TestCaseOperation1Output, error) {
	req, out := c.InputService7TestCaseOperation1Request(input)
	err := req.Send()
	return out, err
}

type InputService7TestShapeInputService7TestCaseOperation1Input struct {
	Foo *string `locationName:"foo" type:"string"`

	metadataInputService7TestShapeInputService7TestCaseOperation1Input `json:"-" xml:"-"`
}

type metadataInputService7TestShapeInputService7TestCaseOperation1Input struct {
	SDKShapeTraits bool `type:"structure" payload:"Foo"`
}

type InputService7TestShapeInputService7TestCaseOperation1Output struct {
	metadataInputService7TestShapeInputService7TestCaseOperation1Output `json:"-" xml:"-"`
}

type metadataInputService7TestShapeInputService7TestCaseOperation1Output struct {
	SDKShapeTraits bool `type:"structure"`
}

type InputService8ProtocolTest struct {
	*service.Service
}

// New returns a new InputService8ProtocolTest client.
func NewInputService8ProtocolTest(config *aws.Config) *InputService8ProtocolTest {
	service := &service.Service{
		ServiceInfo: serviceinfo.ServiceInfo{
			Config:      defaults.DefaultConfig.Merge(config),
			ServiceName: "inputservice8protocoltest",
			APIVersion:  "2014-01-01",
		},
	}
	service.Initialize()

	// Handlers
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Build.PushBack(restjson.Build)
	service.Handlers.Unmarshal.PushBack(restjson.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(restjson.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(restjson.UnmarshalError)

	return &InputService8ProtocolTest{service}
}

// newRequest creates a new request for a InputService8ProtocolTest operation and runs any
// custom request initialization.
func (c *InputService8ProtocolTest) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	return req
}

const opInputService8TestCaseOperation1 = "OperationName"

// InputService8TestCaseOperation1Request generates a request for the InputService8TestCaseOperation1 operation.
func (c *InputService8ProtocolTest) InputService8TestCaseOperation1Request(input *InputService8TestShapeInputShape) (req *request.Request, output *InputService8TestShapeInputService8TestCaseOperation1Output) {
	op := &request.Operation{
		Name:       opInputService8TestCaseOperation1,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &InputService8TestShapeInputShape{}
	}

	req = c.newRequest(op, input, output)
	output = &InputService8TestShapeInputService8TestCaseOperation1Output{}
	req.Data = output
	return
}

func (c *InputService8ProtocolTest) InputService8TestCaseOperation1(input *InputService8TestShapeInputShape) (*InputService8TestShapeInputService8TestCaseOperation1Output, error) {
	req, out := c.InputService8TestCaseOperation1Request(input)
	err := req.Send()
	return out, err
}

const opInputService8TestCaseOperation2 = "OperationName"

// InputService8TestCaseOperation2Request generates a request for the InputService8TestCaseOperation2 operation.
func (c *InputService8ProtocolTest) InputService8TestCaseOperation2Request(input *InputService8TestShapeInputShape) (req *request.Request, output *InputService8TestShapeInputService8TestCaseOperation2Output) {
	op := &request.Operation{
		Name:       opInputService8TestCaseOperation2,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &InputService8TestShapeInputShape{}
	}

	req = c.newRequest(op, input, output)
	output = &InputService8TestShapeInputService8TestCaseOperation2Output{}
	req.Data = output
	return
}

func (c *InputService8ProtocolTest) InputService8TestCaseOperation2(input *InputService8TestShapeInputShape) (*InputService8TestShapeInputService8TestCaseOperation2Output, error) {
	req, out := c.InputService8TestCaseOperation2Request(input)
	err := req.Send()
	return out, err
}

type InputService8TestShapeInputService8TestCaseOperation1Output struct {
	metadataInputService8TestShapeInputService8TestCaseOperation1Output `json:"-" xml:"-"`
}

type metadataInputService8TestShapeInputService8TestCaseOperation1Output struct {
	SDKShapeTraits bool `type:"structure"`
}

type InputService8TestShapeInputService8TestCaseOperation2Output struct {
	metadataInputService8TestShapeInputService8TestCaseOperation2Output `json:"-" xml:"-"`
}

type metadataInputService8TestShapeInputService8TestCaseOperation2Output struct {
	SDKShapeTraits bool `type:"structure"`
}

type InputService8TestShapeInputShape struct {
	Foo []byte `locationName:"foo" type:"blob"`

	metadataInputService8TestShapeInputShape `json:"-" xml:"-"`
}

type metadataInputService8TestShapeInputShape struct {
	SDKShapeTraits bool `type:"structure" payload:"Foo"`
}

type InputService9ProtocolTest struct {
	*service.Service
}

// New returns a new InputService9ProtocolTest client.
func NewInputService9ProtocolTest(config *aws.Config) *InputService9ProtocolTest {
	service := &service.Service{
		ServiceInfo: serviceinfo.ServiceInfo{
			Config:      defaults.DefaultConfig.Merge(config),
			ServiceName: "inputservice9protocoltest",
			APIVersion:  "2014-01-01",
		},
	}
	service.Initialize()

	// Handlers
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Build.PushBack(restjson.Build)
	service.Handlers.Unmarshal.PushBack(restjson.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(restjson.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(restjson.UnmarshalError)

	return &InputService9ProtocolTest{service}
}

// newRequest creates a new request for a InputService9ProtocolTest operation and runs any
// custom request initialization.
func (c *InputService9ProtocolTest) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	return req
}

const opInputService9TestCaseOperation1 = "OperationName"

// InputService9TestCaseOperation1Request generates a request for the InputService9TestCaseOperation1 operation.
func (c *InputService9ProtocolTest) InputService9TestCaseOperation1Request(input *InputService9TestShapeInputShape) (req *request.Request, output *InputService9TestShapeInputService9TestCaseOperation1Output) {
	op := &request.Operation{
		Name:       opInputService9TestCaseOperation1,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &InputService9TestShapeInputShape{}
	}

	req = c.newRequest(op, input, output)
	output = &InputService9TestShapeInputService9TestCaseOperation1Output{}
	req.Data = output
	return
}

func (c *InputService9ProtocolTest) InputService9TestCaseOperation1(input *InputService9TestShapeInputShape) (*InputService9TestShapeInputService9TestCaseOperation1Output, error) {
	req, out := c.InputService9TestCaseOperation1Request(input)
	err := req.Send()
	return out, err
}

const opInputService9TestCaseOperation2 = "OperationName"

// InputService9TestCaseOperation2Request generates a request for the InputService9TestCaseOperation2 operation.
func (c *InputService9ProtocolTest) InputService9TestCaseOperation2Request(input *InputService9TestShapeInputShape) (req *request.Request, output *InputService9TestShapeInputService9TestCaseOperation2Output) {
	op := &request.Operation{
		Name:       opInputService9TestCaseOperation2,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &InputService9TestShapeInputShape{}
	}

	req = c.newRequest(op, input, output)
	output = &InputService9TestShapeInputService9TestCaseOperation2Output{}
	req.Data = output
	return
}

func (c *InputService9ProtocolTest) InputService9TestCaseOperation2(input *InputService9TestShapeInputShape) (*InputService9TestShapeInputService9TestCaseOperation2Output, error) {
	req, out := c.InputService9TestCaseOperation2Request(input)
	err := req.Send()
	return out, err
}

type InputService9TestShapeFooShape struct {
	Baz *string `locationName:"baz" type:"string"`

	metadataInputService9TestShapeFooShape `json:"-" xml:"-"`
}

type metadataInputService9TestShapeFooShape struct {
	SDKShapeTraits bool `locationName:"foo" type:"structure"`
}

type InputService9TestShapeInputService9TestCaseOperation1Output struct {
	metadataInputService9TestShapeInputService9TestCaseOperation1Output `json:"-" xml:"-"`
}

type metadataInputService9TestShapeInputService9TestCaseOperation1Output struct {
	SDKShapeTraits bool `type:"structure"`
}

type InputService9TestShapeInputService9TestCaseOperation2Output struct {
	metadataInputService9TestShapeInputService9TestCaseOperation2Output `json:"-" xml:"-"`
}

type metadataInputService9TestShapeInputService9TestCaseOperation2Output struct {
	SDKShapeTraits bool `type:"structure"`
}

type InputService9TestShapeInputShape struct {
	Foo *InputService9TestShapeFooShape `locationName:"foo" type:"structure"`

	metadataInputService9TestShapeInputShape `json:"-" xml:"-"`
}

type metadataInputService9TestShapeInputShape struct {
	SDKShapeTraits bool `type:"structure" payload:"Foo"`
}

type InputService10ProtocolTest struct {
	*service.Service
}

// New returns a new InputService10ProtocolTest client.
func NewInputService10ProtocolTest(config *aws.Config) *InputService10ProtocolTest {
	service := &service.Service{
		ServiceInfo: serviceinfo.ServiceInfo{
			Config:      defaults.DefaultConfig.Merge(config),
			ServiceName: "inputservice10protocoltest",
			APIVersion:  "2014-01-01",
		},
	}
	service.Initialize()

	// Handlers
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Build.PushBack(restjson.Build)
	service.Handlers.Unmarshal.PushBack(restjson.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(restjson.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(restjson.UnmarshalError)

	return &InputService10ProtocolTest{service}
}

// newRequest creates a new request for a InputService10ProtocolTest operation and runs any
// custom request initialization.
func (c *InputService10ProtocolTest) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	return req
}

const opInputService10TestCaseOperation1 = "OperationName"

// InputService10TestCaseOperation1Request generates a request for the InputService10TestCaseOperation1 operation.
func (c *InputService10ProtocolTest) InputService10TestCaseOperation1Request(input *InputService10TestShapeInputShape) (req *request.Request, output *InputService10TestShapeInputService10TestCaseOperation1Output) {
	op := &request.Operation{
		Name:       opInputService10TestCaseOperation1,
		HTTPMethod: "POST",
		HTTPPath:   "/path",
	}

	if input == nil {
		input = &InputService10TestShapeInputShape{}
	}

	req = c.newRequest(op, input, output)
	output = &InputService10TestShapeInputService10TestCaseOperation1Output{}
	req.Data = output
	return
}

func (c *InputService10ProtocolTest) InputService10TestCaseOperation1(input *InputService10TestShapeInputShape) (*InputService10TestShapeInputService10TestCaseOperation1Output, error) {
	req, out := c.InputService10TestCaseOperation1Request(input)
	err := req.Send()
	return out, err
}

const opInputService10TestCaseOperation2 = "OperationName"

// InputService10TestCaseOperation2Request generates a request for the InputService10TestCaseOperation2 operation.
func (c *InputService10ProtocolTest) InputService10TestCaseOperation2Request(input *InputService10TestShapeInputShape) (req *request.Request, output *InputService10TestShapeInputService10TestCaseOperation2Output) {
	op := &request.Operation{
		Name:       opInputService10TestCaseOperation2,
		HTTPMethod: "POST",
		HTTPPath:   "/path?abc=mno",
	}

	if input == nil {
		input = &InputService10TestShapeInputShape{}
	}

	req = c.newRequest(op, input, output)
	output = &InputService10TestShapeInputService10TestCaseOperation2Output{}
	req.Data = output
	return
}

func (c *InputService10ProtocolTest) InputService10TestCaseOperation2(input *InputService10TestShapeInputShape) (*InputService10TestShapeInputService10TestCaseOperation2Output, error) {
	req, out := c.InputService10TestCaseOperation2Request(input)
	err := req.Send()
	return out, err
}

type InputService10TestShapeInputService10TestCaseOperation1Output struct {
	metadataInputService10TestShapeInputService10TestCaseOperation1Output `json:"-" xml:"-"`
}

type metadataInputService10TestShapeInputService10TestCaseOperation1Output struct {
	SDKShapeTraits bool `type:"structure"`
}

type InputService10TestShapeInputService10TestCaseOperation2Output struct {
	metadataInputService10TestShapeInputService10TestCaseOperation2Output `json:"-" xml:"-"`
}

type metadataInputService10TestShapeInputService10TestCaseOperation2Output struct {
	SDKShapeTraits bool `type:"structure"`
}

type InputService10TestShapeInputShape struct {
	Foo *string `location:"querystring" locationName:"param-name" type:"string"`

	metadataInputService10TestShapeInputShape `json:"-" xml:"-"`
}

type metadataInputService10TestShapeInputShape struct {
	SDKShapeTraits bool `type:"structure"`
}

type InputService11ProtocolTest struct {
	*service.Service
}

// New returns a new InputService11ProtocolTest client.
func NewInputService11ProtocolTest(config *aws.Config) *InputService11ProtocolTest {
	service := &service.Service{
		ServiceInfo: serviceinfo.ServiceInfo{
			Config:      defaults.DefaultConfig.Merge(config),
			ServiceName: "inputservice11protocoltest",
			APIVersion:  "2014-01-01",
		},
	}
	service.Initialize()

	// Handlers
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Build.PushBack(restjson.Build)
	service.Handlers.Unmarshal.PushBack(restjson.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(restjson.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(restjson.UnmarshalError)

	return &InputService11ProtocolTest{service}
}

// newRequest creates a new request for a InputService11ProtocolTest operation and runs any
// custom request initialization.
func (c *InputService11ProtocolTest) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	return req
}

const opInputService11TestCaseOperation1 = "OperationName"

// InputService11TestCaseOperation1Request generates a request for the InputService11TestCaseOperation1 operation.
func (c *InputService11ProtocolTest) InputService11TestCaseOperation1Request(input *InputService11TestShapeInputShape) (req *request.Request, output *InputService11TestShapeInputService11TestCaseOperation1Output) {
	op := &request.Operation{
		Name:       opInputService11TestCaseOperation1,
		HTTPMethod: "POST",
		HTTPPath:   "/path",
	}

	if input == nil {
		input = &InputService11TestShapeInputShape{}
	}

	req = c.newRequest(op, input, output)
	output = &InputService11TestShapeInputService11TestCaseOperation1Output{}
	req.Data = output
	return
}

func (c *InputService11ProtocolTest) InputService11TestCaseOperation1(input *InputService11TestShapeInputShape) (*InputService11TestShapeInputService11TestCaseOperation1Output, error) {
	req, out := c.InputService11TestCaseOperation1Request(input)
	err := req.Send()
	return out, err
}

const opInputService11TestCaseOperation2 = "OperationName"

// InputService11TestCaseOperation2Request generates a request for the InputService11TestCaseOperation2 operation.
func (c *InputService11ProtocolTest) InputService11TestCaseOperation2Request(input *InputService11TestShapeInputShape) (req *request.Request, output *InputService11TestShapeInputService11TestCaseOperation2Output) {
	op := &request.Operation{
		Name:       opInputService11TestCaseOperation2,
		HTTPMethod: "POST",
		HTTPPath:   "/path",
	}

	if input == nil {
		input = &InputService11TestShapeInputShape{}
	}

	req = c.newRequest(op, input, output)
	output = &InputService11TestShapeInputService11TestCaseOperation2Output{}
	req.Data = output
	return
}

func (c *InputService11ProtocolTest) InputService11TestCaseOperation2(input *InputService11TestShapeInputShape) (*InputService11TestShapeInputService11TestCaseOperation2Output, error) {
	req, out := c.InputService11TestCaseOperation2Request(input)
	err := req.Send()
	return out, err
}

const opInputService11TestCaseOperation3 = "OperationName"

// InputService11TestCaseOperation3Request generates a request for the InputService11TestCaseOperation3 operation.
func (c *InputService11ProtocolTest) InputService11TestCaseOperation3Request(input *InputService11TestShapeInputShape) (req *request.Request, output *InputService11TestShapeInputService11TestCaseOperation3Output) {
	op := &request.Operation{
		Name:       opInputService11TestCaseOperation3,
		HTTPMethod: "POST",
		HTTPPath:   "/path",
	}

	if input == nil {
		input = &InputService11TestShapeInputShape{}
	}

	req = c.newRequest(op, input, output)
	output = &InputService11TestShapeInputService11TestCaseOperation3Output{}
	req.Data = output
	return
}

func (c *InputService11ProtocolTest) InputService11TestCaseOperation3(input *InputService11TestShapeInputShape) (*InputService11TestShapeInputService11TestCaseOperation3Output, error) {
	req, out := c.InputService11TestCaseOperation3Request(input)
	err := req.Send()
	return out, err
}

const opInputService11TestCaseOperation4 = "OperationName"

// InputService11TestCaseOperation4Request generates a request for the InputService11TestCaseOperation4 operation.
func (c *InputService11ProtocolTest) InputService11TestCaseOperation4Request(input *InputService11TestShapeInputShape) (req *request.Request, output *InputService11TestShapeInputService11TestCaseOperation4Output) {
	op := &request.Operation{
		Name:       opInputService11TestCaseOperation4,
		HTTPMethod: "POST",
		HTTPPath:   "/path",
	}

	if input == nil {
		input = &InputService11TestShapeInputShape{}
	}

	req = c.newRequest(op, input, output)
	output = &InputService11TestShapeInputService11TestCaseOperation4Output{}
	req.Data = output
	return
}

func (c *InputService11ProtocolTest) InputService11TestCaseOperation4(input *InputService11TestShapeInputShape) (*InputService11TestShapeInputService11TestCaseOperation4Output, error) {
	req, out := c.InputService11TestCaseOperation4Request(input)
	err := req.Send()
	return out, err
}

const opInputService11TestCaseOperation5 = "OperationName"

// InputService11TestCaseOperation5Request generates a request for the InputService11TestCaseOperation5 operation.
func (c *InputService11ProtocolTest) InputService11TestCaseOperation5Request(input *InputService11TestShapeInputShape) (req *request.Request, output *InputService11TestShapeInputService11TestCaseOperation5Output) {
	op := &request.Operation{
		Name:       opInputService11TestCaseOperation5,
		HTTPMethod: "POST",
		HTTPPath:   "/path",
	}

	if input == nil {
		input = &InputService11TestShapeInputShape{}
	}

	req = c.newRequest(op, input, output)
	output = &InputService11TestShapeInputService11TestCaseOperation5Output{}
	req.Data = output
	return
}

func (c *InputService11ProtocolTest) InputService11TestCaseOperation5(input *InputService11TestShapeInputShape) (*InputService11TestShapeInputService11TestCaseOperation5Output, error) {
	req, out := c.InputService11TestCaseOperation5Request(input)
	err := req.Send()
	return out, err
}

const opInputService11TestCaseOperation6 = "OperationName"

// InputService11TestCaseOperation6Request generates a request for the InputService11TestCaseOperation6 operation.
func (c *InputService11ProtocolTest) InputService11TestCaseOperation6Request(input *InputService11TestShapeInputShape) (req *request.Request, output *InputService11TestShapeInputService11TestCaseOperation6Output) {
	op := &request.Operation{
		Name:       opInputService11TestCaseOperation6,
		HTTPMethod: "POST",
		HTTPPath:   "/path",
	}

	if input == nil {
		input = &InputService11TestShapeInputShape{}
	}

	req = c.newRequest(op, input, output)
	output = &InputService11TestShapeInputService11TestCaseOperation6Output{}
	req.Data = output
	return
}

func (c *InputService11ProtocolTest) InputService11TestCaseOperation6(input *InputService11TestShapeInputShape) (*InputService11TestShapeInputService11TestCaseOperation6Output, error) {
	req, out := c.InputService11TestCaseOperation6Request(input)
	err := req.Send()
	return out, err
}

type InputService11TestShapeInputService11TestCaseOperation1Output struct {
	metadataInputService11TestShapeInputService11TestCaseOperation1Output `json:"-" xml:"-"`
}

type metadataInputService11TestShapeInputService11TestCaseOperation1Output struct {
	SDKShapeTraits bool `type:"structure"`
}

type InputService11TestShapeInputService11TestCaseOperation2Output struct {
	metadataInputService11TestShapeInputService11TestCaseOperation2Output `json:"-" xml:"-"`
}

type metadataInputService11TestShapeInputService11TestCaseOperation2Output struct {
	SDKShapeTraits bool `type:"structure"`
}

type InputService11TestShapeInputService11TestCaseOperation3Output struct {
	metadataInputService11TestShapeInputService11TestCaseOperation3Output `json:"-" xml:"-"`
}

type metadataInputService11TestShapeInputService11TestCaseOperation3Output struct {
	SDKShapeTraits bool `type:"structure"`
}

type InputService11TestShapeInputService11TestCaseOperation4Output struct {
	metadataInputService11TestShapeInputService11TestCaseOperation4Output `json:"-" xml:"-"`
}

type metadataInputService11TestShapeInputService11TestCaseOperation4Output struct {
	SDKShapeTraits bool `type:"structure"`
}

type InputService11TestShapeInputService11TestCaseOperation5Output struct {
	metadataInputService11TestShapeInputService11TestCaseOperation5Output `json:"-" xml:"-"`
}

type metadataInputService11TestShapeInputService11TestCaseOperation5Output struct {
	SDKShapeTraits bool `type:"structure"`
}

type InputService11TestShapeInputService11TestCaseOperation6Output struct {
	metadataInputService11TestShapeInputService11TestCaseOperation6Output `json:"-" xml:"-"`
}

type metadataInputService11TestShapeInputService11TestCaseOperation6Output struct {
	SDKShapeTraits bool `type:"structure"`
}

type InputService11TestShapeInputShape struct {
	RecursiveStruct *InputService11TestShapeRecursiveStructType `type:"structure"`

	metadataInputService11TestShapeInputShape `json:"-" xml:"-"`
}

type metadataInputService11TestShapeInputShape struct {
	SDKShapeTraits bool `type:"structure"`
}

type InputService11TestShapeRecursiveStructType struct {
	NoRecurse *string `type:"string"`

	RecursiveList []*InputService11TestShapeRecursiveStructType `type:"list"`

	RecursiveMap map[string]*InputService11TestShapeRecursiveStructType `type:"map"`

	RecursiveStruct *InputService11TestShapeRecursiveStructType `type:"structure"`

	metadataInputService11TestShapeRecursiveStructType `json:"-" xml:"-"`
}

type metadataInputService11TestShapeRecursiveStructType struct {
	SDKShapeTraits bool `type:"structure"`
}

type InputService12ProtocolTest struct {
	*service.Service
}

// New returns a new InputService12ProtocolTest client.
func NewInputService12ProtocolTest(config *aws.Config) *InputService12ProtocolTest {
	service := &service.Service{
		ServiceInfo: serviceinfo.ServiceInfo{
			Config:      defaults.DefaultConfig.Merge(config),
			ServiceName: "inputservice12protocoltest",
			APIVersion:  "2014-01-01",
		},
	}
	service.Initialize()

	// Handlers
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Build.PushBack(restjson.Build)
	service.Handlers.Unmarshal.PushBack(restjson.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(restjson.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(restjson.UnmarshalError)

	return &InputService12ProtocolTest{service}
}

// newRequest creates a new request for a InputService12ProtocolTest operation and runs any
// custom request initialization.
func (c *InputService12ProtocolTest) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	return req
}

const opInputService12TestCaseOperation1 = "OperationName"

// InputService12TestCaseOperation1Request generates a request for the InputService12TestCaseOperation1 operation.
func (c *InputService12ProtocolTest) InputService12TestCaseOperation1Request(input *InputService12TestShapeInputShape) (req *request.Request, output *InputService12TestShapeInputService12TestCaseOperation1Output) {
	op := &request.Operation{
		Name:       opInputService12TestCaseOperation1,
		HTTPMethod: "POST",
		HTTPPath:   "/path",
	}

	if input == nil {
		input = &InputService12TestShapeInputShape{}
	}

	req = c.newRequest(op, input, output)
	output = &InputService12TestShapeInputService12TestCaseOperation1Output{}
	req.Data = output
	return
}

func (c *InputService12ProtocolTest) InputService12TestCaseOperation1(input *InputService12TestShapeInputShape) (*InputService12TestShapeInputService12TestCaseOperation1Output, error) {
	req, out := c.InputService12TestCaseOperation1Request(input)
	err := req.Send()
	return out, err
}

const opInputService12TestCaseOperation2 = "OperationName"

// InputService12TestCaseOperation2Request generates a request for the InputService12TestCaseOperation2 operation.
func (c *InputService12ProtocolTest) InputService12TestCaseOperation2Request(input *InputService12TestShapeInputShape) (req *request.Request, output *InputService12TestShapeInputService12TestCaseOperation2Output) {
	op := &request.Operation{
		Name:       opInputService12TestCaseOperation2,
		HTTPMethod: "POST",
		HTTPPath:   "/path",
	}

	if input == nil {
		input = &InputService12TestShapeInputShape{}
	}

	req = c.newRequest(op, input, output)
	output = &InputService12TestShapeInputService12TestCaseOperation2Output{}
	req.Data = output
	return
}

func (c *InputService12ProtocolTest) InputService12TestCaseOperation2(input *InputService12TestShapeInputShape) (*InputService12TestShapeInputService12TestCaseOperation2Output, error) {
	req, out := c.InputService12TestCaseOperation2Request(input)
	err := req.Send()
	return out, err
}

type InputService12TestShapeInputService12TestCaseOperation1Output struct {
	metadataInputService12TestShapeInputService12TestCaseOperation1Output `json:"-" xml:"-"`
}

type metadataInputService12TestShapeInputService12TestCaseOperation1Output struct {
	SDKShapeTraits bool `type:"structure"`
}

type InputService12TestShapeInputService12TestCaseOperation2Output struct {
	metadataInputService12TestShapeInputService12TestCaseOperation2Output `json:"-" xml:"-"`
}

type metadataInputService12TestShapeInputService12TestCaseOperation2Output struct {
	SDKShapeTraits bool `type:"structure"`
}

type InputService12TestShapeInputShape struct {
	TimeArg *time.Time `type:"timestamp" timestampFormat:"unix"`

	TimeArgInHeader *time.Time `location:"header" locationName:"x-amz-timearg" type:"timestamp" timestampFormat:"rfc822"`

	metadataInputService12TestShapeInputShape `json:"-" xml:"-"`
}

type metadataInputService12TestShapeInputShape struct {
	SDKShapeTraits bool `type:"structure"`
}

//
// Tests begin here
//

func TestInputService1ProtocolTestURIParameterOnlyWithNoLocationNameCase1(t *testing.T) {
	svc := NewInputService1ProtocolTest(nil)
	svc.Endpoint = "https://test"

	input := &InputService1TestShapeInputService1TestCaseOperation1Input{
		PipelineId: aws.String("foo"),
	}
	req, _ := svc.InputService1TestCaseOperation1Request(input)
	r := req.HTTPRequest

	// build request
	restjson.Build(req)
	assert.NoError(t, req.Error)

	// assert URL
	assert.Equal(t, "https://test/2014-01-01/jobsByPipeline/foo", r.URL.String())

	// assert headers

}

func TestInputService2ProtocolTestURIParameterOnlyWithLocationNameCase1(t *testing.T) {
	svc := NewInputService2ProtocolTest(nil)
	svc.Endpoint = "https://test"

	input := &InputService2TestShapeInputService2TestCaseOperation1Input{
		Foo: aws.String("bar"),
	}
	req, _ := svc.InputService2TestCaseOperation1Request(input)
	r := req.HTTPRequest

	// build request
	restjson.Build(req)
	assert.NoError(t, req.Error)

	// assert URL
	assert.Equal(t, "https://test/2014-01-01/jobsByPipeline/bar", r.URL.String())

	// assert headers

}

func TestInputService3ProtocolTestURIParameterAndQuerystringParamsCase1(t *testing.T) {
	svc := NewInputService3ProtocolTest(nil)
	svc.Endpoint = "https://test"

	input := &InputService3TestShapeInputService3TestCaseOperation1Input{
		Ascending:  aws.String("true"),
		PageToken:  aws.String("bar"),
		PipelineId: aws.String("foo"),
	}
	req, _ := svc.InputService3TestCaseOperation1Request(input)
	r := req.HTTPRequest

	// build request
	restjson.Build(req)
	assert.NoError(t, req.Error)

	// assert URL
	assert.Equal(t, "https://test/2014-01-01/jobsByPipeline/foo?Ascending=true&PageToken=bar", r.URL.String())

	// assert headers

}

func TestInputService4ProtocolTestURIParameterQuerystringParamsAndJSONBodyCase1(t *testing.T) {
	svc := NewInputService4ProtocolTest(nil)
	svc.Endpoint = "https://test"

	input := &InputService4TestShapeInputService4TestCaseOperation1Input{
		Ascending: aws.String("true"),
		Config: &InputService4TestShapeStructType{
			A: aws.String("one"),
			B: aws.String("two"),
		},
		PageToken:  aws.String("bar"),
		PipelineId: aws.String("foo"),
	}
	req, _ := svc.InputService4TestCaseOperation1Request(input)
	r := req.HTTPRequest

	// build request
	restjson.Build(req)
	assert.NoError(t, req.Error)

	// assert body
	assert.NotNil(t, r.Body)
	body, _ := ioutil.ReadAll(r.Body)
	assert.Equal(t, util.Trim(`{"Config":{"A":"one","B":"two"}}`), util.Trim(string(body)))

	// assert URL
	assert.Equal(t, "https://test/2014-01-01/jobsByPipeline/foo?Ascending=true&PageToken=bar", r.URL.String())

	// assert headers

}

func TestInputService5ProtocolTestURIParameterQuerystringParamsHeadersAndJSONBodyCase1(t *testing.T) {
	svc := NewInputService5ProtocolTest(nil)
	svc.Endpoint = "https://test"

	input := &InputService5TestShapeInputService5TestCaseOperation1Input{
		Ascending: aws.String("true"),
		Checksum:  aws.String("12345"),
		Config: &InputService5TestShapeStructType{
			A: aws.String("one"),
			B: aws.String("two"),
		},
		PageToken:  aws.String("bar"),
		PipelineId: aws.String("foo"),
	}
	req, _ := svc.InputService5TestCaseOperation1Request(input)
	r := req.HTTPRequest

	// build request
	restjson.Build(req)
	assert.NoError(t, req.Error)

	// assert body
	assert.NotNil(t, r.Body)
	body, _ := ioutil.ReadAll(r.Body)
	assert.Equal(t, util.Trim(`{"Config":{"A":"one","B":"two"}}`), util.Trim(string(body)))

	// assert URL
	assert.Equal(t, "https://test/2014-01-01/jobsByPipeline/foo?Ascending=true&PageToken=bar", r.URL.String())

	// assert headers
	assert.Equal(t, "12345", r.Header.Get("x-amz-checksum"))

}

func TestInputService6ProtocolTestStreamingPayloadCase1(t *testing.T) {
	svc := NewInputService6ProtocolTest(nil)
	svc.Endpoint = "https://test"

	input := &InputService6TestShapeInputService6TestCaseOperation1Input{
		Body:      aws.ReadSeekCloser(bytes.NewBufferString("contents")),
		Checksum:  aws.String("foo"),
		VaultName: aws.String("name"),
	}
	req, _ := svc.InputService6TestCaseOperation1Request(input)
	r := req.HTTPRequest

	// build request
	restjson.Build(req)
	assert.NoError(t, req.Error)

	// assert body
	assert.NotNil(t, r.Body)
	body, _ := ioutil.ReadAll(r.Body)
	assert.Equal(t, util.Trim(`contents`), util.Trim(string(body)))

	// assert URL
	assert.Equal(t, "https://test/2014-01-01/vaults/name/archives", r.URL.String())

	// assert headers
	assert.Equal(t, "foo", r.Header.Get("x-amz-sha256-tree-hash"))

}

func TestInputService7ProtocolTestStringPayloadCase1(t *testing.T) {
	svc := NewInputService7ProtocolTest(nil)
	svc.Endpoint = "https://test"

	input := &InputService7TestShapeInputService7TestCaseOperation1Input{
		Foo: aws.String("bar"),
	}
	req, _ := svc.InputService7TestCaseOperation1Request(input)
	r := req.HTTPRequest

	// build request
	restjson.Build(req)
	assert.NoError(t, req.Error)

	// assert body
	assert.NotNil(t, r.Body)
	body, _ := ioutil.ReadAll(r.Body)
	assert.Equal(t, util.Trim(`bar`), util.Trim(string(body)))

	// assert URL
	assert.Equal(t, "https://test/", r.URL.String())

	// assert headers

}

func TestInputService8ProtocolTestBlobPayloadCase1(t *testing.T) {
	svc := NewInputService8ProtocolTest(nil)
	svc.Endpoint = "https://test"

	input := &InputService8TestShapeInputShape{
		Foo: []byte("bar"),
	}
	req, _ := svc.InputService8TestCaseOperation1Request(input)
	r := req.HTTPRequest

	// build request
	restjson.Build(req)
	assert.NoError(t, req.Error)

	// assert body
	assert.NotNil(t, r.Body)
	body, _ := ioutil.ReadAll(r.Body)
	assert.Equal(t, util.Trim(`bar`), util.Trim(string(body)))

	// assert URL
	assert.Equal(t, "https://test/", r.URL.String())

	// assert headers

}

func TestInputService8ProtocolTestBlobPayloadCase2(t *testing.T) {
	svc := NewInputService8ProtocolTest(nil)
	svc.Endpoint = "https://test"

	input := &InputService8TestShapeInputShape{}
	req, _ := svc.InputService8TestCaseOperation2Request(input)
	r := req.HTTPRequest

	// build request
	restjson.Build(req)
	assert.NoError(t, req.Error)

	// assert URL
	assert.Equal(t, "https://test/", r.URL.String())

	// assert headers

}

func TestInputService9ProtocolTestStructurePayloadCase1(t *testing.T) {
	svc := NewInputService9ProtocolTest(nil)
	svc.Endpoint = "https://test"

	input := &InputService9TestShapeInputShape{
		Foo: &InputService9TestShapeFooShape{
			Baz: aws.String("bar"),
		},
	}
	req, _ := svc.InputService9TestCaseOperation1Request(input)
	r := req.HTTPRequest

	// build request
	restjson.Build(req)
	assert.NoError(t, req.Error)

	// assert body
	assert.NotNil(t, r.Body)
	body, _ := ioutil.ReadAll(r.Body)
	assert.Equal(t, util.Trim(`{"baz":"bar"}`), util.Trim(string(body)))

	// assert URL
	assert.Equal(t, "https://test/", r.URL.String())

	// assert headers

}

func TestInputService9ProtocolTestStructurePayloadCase2(t *testing.T) {
	svc := NewInputService9ProtocolTest(nil)
	svc.Endpoint = "https://test"

	input := &InputService9TestShapeInputShape{}
	req, _ := svc.InputService9TestCaseOperation2Request(input)
	r := req.HTTPRequest

	// build request
	restjson.Build(req)
	assert.NoError(t, req.Error)

	// assert URL
	assert.Equal(t, "https://test/", r.URL.String())

	// assert headers

}

func TestInputService10ProtocolTestOmitsNullQueryParamsButSerializesEmptyStringsCase1(t *testing.T) {
	svc := NewInputService10ProtocolTest(nil)
	svc.Endpoint = "https://test"

	input := &InputService10TestShapeInputShape{}
	req, _ := svc.InputService10TestCaseOperation1Request(input)
	r := req.HTTPRequest

	// build request
	restjson.Build(req)
	assert.NoError(t, req.Error)

	// assert URL
	assert.Equal(t, "https://test/path", r.URL.String())

	// assert headers

}

func TestInputService10ProtocolTestOmitsNullQueryParamsButSerializesEmptyStringsCase2(t *testing.T) {
	svc := NewInputService10ProtocolTest(nil)
	svc.Endpoint = "https://test"

	input := &InputService10TestShapeInputShape{
		Foo: aws.String(""),
	}
	req, _ := svc.InputService10TestCaseOperation2Request(input)
	r := req.HTTPRequest

	// build request
	restjson.Build(req)
	assert.NoError(t, req.Error)

	// assert URL
	assert.Equal(t, "https://test/path?abc=mno&param-name=", r.URL.String())

	// assert headers

}

func TestInputService11ProtocolTestRecursiveShapesCase1(t *testing.T) {
	svc := NewInputService11ProtocolTest(nil)
	svc.Endpoint = "https://test"

	input := &InputService11TestShapeInputShape{
		RecursiveStruct: &InputService11TestShapeRecursiveStructType{
			NoRecurse: aws.String("foo"),
		},
	}
	req, _ := svc.InputService11TestCaseOperation1Request(input)
	r := req.HTTPRequest

	// build request
	restjson.Build(req)
	assert.NoError(t, req.Error)

	// assert body
	assert.NotNil(t, r.Body)
	body, _ := ioutil.ReadAll(r.Body)
	assert.Equal(t, util.Trim(`{"RecursiveStruct":{"NoRecurse":"foo"}}`), util.Trim(string(body)))

	// assert URL
	assert.Equal(t, "https://test/path", r.URL.String())

	// assert headers

}

func TestInputService11ProtocolTestRecursiveShapesCase2(t *testing.T) {
	svc := NewInputService11ProtocolTest(nil)
	svc.Endpoint = "https://test"

	input := &InputService11TestShapeInputShape{
		RecursiveStruct: &InputService11TestShapeRecursiveStructType{
			RecursiveStruct: &InputService11TestShapeRecursiveStructType{
				NoRecurse: aws.String("foo"),
			},
		},
	}
	req, _ := svc.InputService11TestCaseOperation2Request(input)
	r := req.HTTPRequest

	// build request
	restjson.Build(req)
	assert.NoError(t, req.Error)

	// assert body
	assert.NotNil(t, r.Body)
	body, _ := ioutil.ReadAll(r.Body)
	assert.Equal(t, util.Trim(`{"RecursiveStruct":{"RecursiveStruct":{"NoRecurse":"foo"}}}`), util.Trim(string(body)))

	// assert URL
	assert.Equal(t, "https://test/path", r.URL.String())

	// assert headers

}

func TestInputService11ProtocolTestRecursiveShapesCase3(t *testing.T) {
	svc := NewInputService11ProtocolTest(nil)
	svc.Endpoint = "https://test"

	input := &InputService11TestShapeInputShape{
		RecursiveStruct: &InputService11TestShapeRecursiveStructType{
			RecursiveStruct: &InputService11TestShapeRecursiveStructType{
				RecursiveStruct: &InputService11TestShapeRecursiveStructType{
					RecursiveStruct: &InputService11TestShapeRecursiveStructType{
						NoRecurse: aws.String("foo"),
					},
				},
			},
		},
	}
	req, _ := svc.InputService11TestCaseOperation3Request(input)
	r := req.HTTPRequest

	// build request
	restjson.Build(req)
	assert.NoError(t, req.Error)

	// assert body
	assert.NotNil(t, r.Body)
	body, _ := ioutil.ReadAll(r.Body)
	assert.Equal(t, util.Trim(`{"RecursiveStruct":{"RecursiveStruct":{"RecursiveStruct":{"RecursiveStruct":{"NoRecurse":"foo"}}}}}`), util.Trim(string(body)))

	// assert URL
	assert.Equal(t, "https://test/path", r.URL.String())

	// assert headers

}

func TestInputService11ProtocolTestRecursiveShapesCase4(t *testing.T) {
	svc := NewInputService11ProtocolTest(nil)
	svc.Endpoint = "https://test"

	input := &InputService11TestShapeInputShape{
		RecursiveStruct: &InputService11TestShapeRecursiveStructType{
			RecursiveList: []*InputService11TestShapeRecursiveStructType{
				{
					NoRecurse: aws.String("foo"),
				},
				{
					NoRecurse: aws.String("bar"),
				},
			},
		},
	}
	req, _ := svc.InputService11TestCaseOperation4Request(input)
	r := req.HTTPRequest

	// build request
	restjson.Build(req)
	assert.NoError(t, req.Error)

	// assert body
	assert.NotNil(t, r.Body)
	body, _ := ioutil.ReadAll(r.Body)
	assert.Equal(t, util.Trim(`{"RecursiveStruct":{"RecursiveList":[{"NoRecurse":"foo"},{"NoRecurse":"bar"}]}}`), util.Trim(string(body)))

	// assert URL
	assert.Equal(t, "https://test/path", r.URL.String())

	// assert headers

}

func TestInputService11ProtocolTestRecursiveShapesCase5(t *testing.T) {
	svc := NewInputService11ProtocolTest(nil)
	svc.Endpoint = "https://test"

	input := &InputService11TestShapeInputShape{
		RecursiveStruct: &InputService11TestShapeRecursiveStructType{
			RecursiveList: []*InputService11TestShapeRecursiveStructType{
				{
					NoRecurse: aws.String("foo"),
				},
				{
					RecursiveStruct: &InputService11TestShapeRecursiveStructType{
						NoRecurse: aws.String("bar"),
					},
				},
			},
		},
	}
	req, _ := svc.InputService11TestCaseOperation5Request(input)
	r := req.HTTPRequest

	// build request
	restjson.Build(req)
	assert.NoError(t, req.Error)

	// assert body
	assert.NotNil(t, r.Body)
	body, _ := ioutil.ReadAll(r.Body)
	assert.Equal(t, util.Trim(`{"RecursiveStruct":{"RecursiveList":[{"NoRecurse":"foo"},{"RecursiveStruct":{"NoRecurse":"bar"}}]}}`), util.Trim(string(body)))

	// assert URL
	assert.Equal(t, "https://test/path", r.URL.String())

	// assert headers

}

func TestInputService11ProtocolTestRecursiveShapesCase6(t *testing.T) {
	svc := NewInputService11ProtocolTest(nil)
	svc.Endpoint = "https://test"

	input := &InputService11TestShapeInputShape{
		RecursiveStruct: &InputService11TestShapeRecursiveStructType{
			RecursiveMap: map[string]*InputService11TestShapeRecursiveStructType{
				"bar": {
					NoRecurse: aws.String("bar"),
				},
				"foo": {
					NoRecurse: aws.String("foo"),
				},
			},
		},
	}
	req, _ := svc.InputService11TestCaseOperation6Request(input)
	r := req.HTTPRequest

	// build request
	restjson.Build(req)
	assert.NoError(t, req.Error)

	// assert body
	assert.NotNil(t, r.Body)
	body, _ := ioutil.ReadAll(r.Body)
	assert.Equal(t, util.Trim(`{"RecursiveStruct":{"RecursiveMap":{"bar":{"NoRecurse":"bar"},"foo":{"NoRecurse":"foo"}}}}`), util.Trim(string(body)))

	// assert URL
	assert.Equal(t, "https://test/path", r.URL.String())

	// assert headers

}

func TestInputService12ProtocolTestTimestampValuesCase1(t *testing.T) {
	svc := NewInputService12ProtocolTest(nil)
	svc.Endpoint = "https://test"

	input := &InputService12TestShapeInputShape{
		TimeArg: aws.Time(time.Unix(1422172800, 0)),
	}
	req, _ := svc.InputService12TestCaseOperation1Request(input)
	r := req.HTTPRequest

	// build request
	restjson.Build(req)
	assert.NoError(t, req.Error)

	// assert body
	assert.NotNil(t, r.Body)
	body, _ := ioutil.ReadAll(r.Body)
	assert.Equal(t, util.Trim(`{"TimeArg":1422172800}`), util.Trim(string(body)))

	// assert URL
	assert.Equal(t, "https://test/path", r.URL.String())

	// assert headers

}

func TestInputService12ProtocolTestTimestampValuesCase2(t *testing.T) {
	svc := NewInputService12ProtocolTest(nil)
	svc.Endpoint = "https://test"

	input := &InputService12TestShapeInputShape{
		TimeArgInHeader: aws.Time(time.Unix(1422172800, 0)),
	}
	req, _ := svc.InputService12TestCaseOperation2Request(input)
	r := req.HTTPRequest

	// build request
	restjson.Build(req)
	assert.NoError(t, req.Error)

	// assert URL
	assert.Equal(t, "https://test/path", r.URL.String())

	// assert headers
	assert.Equal(t, "Sun, 25 Jan 2015 08:00:00 GMT", r.Header.Get("x-amz-timearg"))

}
