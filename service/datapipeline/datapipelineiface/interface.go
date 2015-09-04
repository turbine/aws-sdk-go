// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package datapipelineiface provides an interface for the AWS Data Pipeline.
package datapipelineiface

import (
	"github.com/turbine/aws-sdk-go/aws/request"
	"github.com/turbine/aws-sdk-go/service/datapipeline"
)

// DataPipelineAPI is the interface type for datapipeline.DataPipeline.
type DataPipelineAPI interface {
	ActivatePipelineRequest(*datapipeline.ActivatePipelineInput) (*request.Request, *datapipeline.ActivatePipelineOutput)

	ActivatePipeline(*datapipeline.ActivatePipelineInput) (*datapipeline.ActivatePipelineOutput, error)

	AddTagsRequest(*datapipeline.AddTagsInput) (*request.Request, *datapipeline.AddTagsOutput)

	AddTags(*datapipeline.AddTagsInput) (*datapipeline.AddTagsOutput, error)

	CreatePipelineRequest(*datapipeline.CreatePipelineInput) (*request.Request, *datapipeline.CreatePipelineOutput)

	CreatePipeline(*datapipeline.CreatePipelineInput) (*datapipeline.CreatePipelineOutput, error)

	DeactivatePipelineRequest(*datapipeline.DeactivatePipelineInput) (*request.Request, *datapipeline.DeactivatePipelineOutput)

	DeactivatePipeline(*datapipeline.DeactivatePipelineInput) (*datapipeline.DeactivatePipelineOutput, error)

	DeletePipelineRequest(*datapipeline.DeletePipelineInput) (*request.Request, *datapipeline.DeletePipelineOutput)

	DeletePipeline(*datapipeline.DeletePipelineInput) (*datapipeline.DeletePipelineOutput, error)

	DescribeObjectsRequest(*datapipeline.DescribeObjectsInput) (*request.Request, *datapipeline.DescribeObjectsOutput)

	DescribeObjects(*datapipeline.DescribeObjectsInput) (*datapipeline.DescribeObjectsOutput, error)

	DescribeObjectsPages(*datapipeline.DescribeObjectsInput, func(*datapipeline.DescribeObjectsOutput, bool) bool) error

	DescribePipelinesRequest(*datapipeline.DescribePipelinesInput) (*request.Request, *datapipeline.DescribePipelinesOutput)

	DescribePipelines(*datapipeline.DescribePipelinesInput) (*datapipeline.DescribePipelinesOutput, error)

	EvaluateExpressionRequest(*datapipeline.EvaluateExpressionInput) (*request.Request, *datapipeline.EvaluateExpressionOutput)

	EvaluateExpression(*datapipeline.EvaluateExpressionInput) (*datapipeline.EvaluateExpressionOutput, error)

	GetPipelineDefinitionRequest(*datapipeline.GetPipelineDefinitionInput) (*request.Request, *datapipeline.GetPipelineDefinitionOutput)

	GetPipelineDefinition(*datapipeline.GetPipelineDefinitionInput) (*datapipeline.GetPipelineDefinitionOutput, error)

	ListPipelinesRequest(*datapipeline.ListPipelinesInput) (*request.Request, *datapipeline.ListPipelinesOutput)

	ListPipelines(*datapipeline.ListPipelinesInput) (*datapipeline.ListPipelinesOutput, error)

	ListPipelinesPages(*datapipeline.ListPipelinesInput, func(*datapipeline.ListPipelinesOutput, bool) bool) error

	PollForTaskRequest(*datapipeline.PollForTaskInput) (*request.Request, *datapipeline.PollForTaskOutput)

	PollForTask(*datapipeline.PollForTaskInput) (*datapipeline.PollForTaskOutput, error)

	PutPipelineDefinitionRequest(*datapipeline.PutPipelineDefinitionInput) (*request.Request, *datapipeline.PutPipelineDefinitionOutput)

	PutPipelineDefinition(*datapipeline.PutPipelineDefinitionInput) (*datapipeline.PutPipelineDefinitionOutput, error)

	QueryObjectsRequest(*datapipeline.QueryObjectsInput) (*request.Request, *datapipeline.QueryObjectsOutput)

	QueryObjects(*datapipeline.QueryObjectsInput) (*datapipeline.QueryObjectsOutput, error)

	QueryObjectsPages(*datapipeline.QueryObjectsInput, func(*datapipeline.QueryObjectsOutput, bool) bool) error

	RemoveTagsRequest(*datapipeline.RemoveTagsInput) (*request.Request, *datapipeline.RemoveTagsOutput)

	RemoveTags(*datapipeline.RemoveTagsInput) (*datapipeline.RemoveTagsOutput, error)

	ReportTaskProgressRequest(*datapipeline.ReportTaskProgressInput) (*request.Request, *datapipeline.ReportTaskProgressOutput)

	ReportTaskProgress(*datapipeline.ReportTaskProgressInput) (*datapipeline.ReportTaskProgressOutput, error)

	ReportTaskRunnerHeartbeatRequest(*datapipeline.ReportTaskRunnerHeartbeatInput) (*request.Request, *datapipeline.ReportTaskRunnerHeartbeatOutput)

	ReportTaskRunnerHeartbeat(*datapipeline.ReportTaskRunnerHeartbeatInput) (*datapipeline.ReportTaskRunnerHeartbeatOutput, error)

	SetStatusRequest(*datapipeline.SetStatusInput) (*request.Request, *datapipeline.SetStatusOutput)

	SetStatus(*datapipeline.SetStatusInput) (*datapipeline.SetStatusOutput, error)

	SetTaskStatusRequest(*datapipeline.SetTaskStatusInput) (*request.Request, *datapipeline.SetTaskStatusOutput)

	SetTaskStatus(*datapipeline.SetTaskStatusInput) (*datapipeline.SetTaskStatusOutput, error)

	ValidatePipelineDefinitionRequest(*datapipeline.ValidatePipelineDefinitionInput) (*request.Request, *datapipeline.ValidatePipelineDefinitionOutput)

	ValidatePipelineDefinition(*datapipeline.ValidatePipelineDefinitionInput) (*datapipeline.ValidatePipelineDefinitionOutput, error)
}
