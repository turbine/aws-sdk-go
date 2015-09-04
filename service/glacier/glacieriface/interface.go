// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package glacieriface provides an interface for the Amazon Glacier.
package glacieriface

import (
	"github.com/turbine/aws-sdk-go/aws/request"
	"github.com/turbine/aws-sdk-go/service/glacier"
)

// GlacierAPI is the interface type for glacier.Glacier.
type GlacierAPI interface {
	AbortMultipartUploadRequest(*glacier.AbortMultipartUploadInput) (*request.Request, *glacier.AbortMultipartUploadOutput)

	AbortMultipartUpload(*glacier.AbortMultipartUploadInput) (*glacier.AbortMultipartUploadOutput, error)

	AbortVaultLockRequest(*glacier.AbortVaultLockInput) (*request.Request, *glacier.AbortVaultLockOutput)

	AbortVaultLock(*glacier.AbortVaultLockInput) (*glacier.AbortVaultLockOutput, error)

	AddTagsToVaultRequest(*glacier.AddTagsToVaultInput) (*request.Request, *glacier.AddTagsToVaultOutput)

	AddTagsToVault(*glacier.AddTagsToVaultInput) (*glacier.AddTagsToVaultOutput, error)

	CompleteMultipartUploadRequest(*glacier.CompleteMultipartUploadInput) (*request.Request, *glacier.ArchiveCreationOutput)

	CompleteMultipartUpload(*glacier.CompleteMultipartUploadInput) (*glacier.ArchiveCreationOutput, error)

	CompleteVaultLockRequest(*glacier.CompleteVaultLockInput) (*request.Request, *glacier.CompleteVaultLockOutput)

	CompleteVaultLock(*glacier.CompleteVaultLockInput) (*glacier.CompleteVaultLockOutput, error)

	CreateVaultRequest(*glacier.CreateVaultInput) (*request.Request, *glacier.CreateVaultOutput)

	CreateVault(*glacier.CreateVaultInput) (*glacier.CreateVaultOutput, error)

	DeleteArchiveRequest(*glacier.DeleteArchiveInput) (*request.Request, *glacier.DeleteArchiveOutput)

	DeleteArchive(*glacier.DeleteArchiveInput) (*glacier.DeleteArchiveOutput, error)

	DeleteVaultRequest(*glacier.DeleteVaultInput) (*request.Request, *glacier.DeleteVaultOutput)

	DeleteVault(*glacier.DeleteVaultInput) (*glacier.DeleteVaultOutput, error)

	DeleteVaultAccessPolicyRequest(*glacier.DeleteVaultAccessPolicyInput) (*request.Request, *glacier.DeleteVaultAccessPolicyOutput)

	DeleteVaultAccessPolicy(*glacier.DeleteVaultAccessPolicyInput) (*glacier.DeleteVaultAccessPolicyOutput, error)

	DeleteVaultNotificationsRequest(*glacier.DeleteVaultNotificationsInput) (*request.Request, *glacier.DeleteVaultNotificationsOutput)

	DeleteVaultNotifications(*glacier.DeleteVaultNotificationsInput) (*glacier.DeleteVaultNotificationsOutput, error)

	DescribeJobRequest(*glacier.DescribeJobInput) (*request.Request, *glacier.JobDescription)

	DescribeJob(*glacier.DescribeJobInput) (*glacier.JobDescription, error)

	DescribeVaultRequest(*glacier.DescribeVaultInput) (*request.Request, *glacier.DescribeVaultOutput)

	DescribeVault(*glacier.DescribeVaultInput) (*glacier.DescribeVaultOutput, error)

	GetDataRetrievalPolicyRequest(*glacier.GetDataRetrievalPolicyInput) (*request.Request, *glacier.GetDataRetrievalPolicyOutput)

	GetDataRetrievalPolicy(*glacier.GetDataRetrievalPolicyInput) (*glacier.GetDataRetrievalPolicyOutput, error)

	GetJobOutputRequest(*glacier.GetJobOutputInput) (*request.Request, *glacier.GetJobOutputOutput)

	GetJobOutput(*glacier.GetJobOutputInput) (*glacier.GetJobOutputOutput, error)

	GetVaultAccessPolicyRequest(*glacier.GetVaultAccessPolicyInput) (*request.Request, *glacier.GetVaultAccessPolicyOutput)

	GetVaultAccessPolicy(*glacier.GetVaultAccessPolicyInput) (*glacier.GetVaultAccessPolicyOutput, error)

	GetVaultLockRequest(*glacier.GetVaultLockInput) (*request.Request, *glacier.GetVaultLockOutput)

	GetVaultLock(*glacier.GetVaultLockInput) (*glacier.GetVaultLockOutput, error)

	GetVaultNotificationsRequest(*glacier.GetVaultNotificationsInput) (*request.Request, *glacier.GetVaultNotificationsOutput)

	GetVaultNotifications(*glacier.GetVaultNotificationsInput) (*glacier.GetVaultNotificationsOutput, error)

	InitiateJobRequest(*glacier.InitiateJobInput) (*request.Request, *glacier.InitiateJobOutput)

	InitiateJob(*glacier.InitiateJobInput) (*glacier.InitiateJobOutput, error)

	InitiateMultipartUploadRequest(*glacier.InitiateMultipartUploadInput) (*request.Request, *glacier.InitiateMultipartUploadOutput)

	InitiateMultipartUpload(*glacier.InitiateMultipartUploadInput) (*glacier.InitiateMultipartUploadOutput, error)

	InitiateVaultLockRequest(*glacier.InitiateVaultLockInput) (*request.Request, *glacier.InitiateVaultLockOutput)

	InitiateVaultLock(*glacier.InitiateVaultLockInput) (*glacier.InitiateVaultLockOutput, error)

	ListJobsRequest(*glacier.ListJobsInput) (*request.Request, *glacier.ListJobsOutput)

	ListJobs(*glacier.ListJobsInput) (*glacier.ListJobsOutput, error)

	ListJobsPages(*glacier.ListJobsInput, func(*glacier.ListJobsOutput, bool) bool) error

	ListMultipartUploadsRequest(*glacier.ListMultipartUploadsInput) (*request.Request, *glacier.ListMultipartUploadsOutput)

	ListMultipartUploads(*glacier.ListMultipartUploadsInput) (*glacier.ListMultipartUploadsOutput, error)

	ListMultipartUploadsPages(*glacier.ListMultipartUploadsInput, func(*glacier.ListMultipartUploadsOutput, bool) bool) error

	ListPartsRequest(*glacier.ListPartsInput) (*request.Request, *glacier.ListPartsOutput)

	ListParts(*glacier.ListPartsInput) (*glacier.ListPartsOutput, error)

	ListPartsPages(*glacier.ListPartsInput, func(*glacier.ListPartsOutput, bool) bool) error

	ListTagsForVaultRequest(*glacier.ListTagsForVaultInput) (*request.Request, *glacier.ListTagsForVaultOutput)

	ListTagsForVault(*glacier.ListTagsForVaultInput) (*glacier.ListTagsForVaultOutput, error)

	ListVaultsRequest(*glacier.ListVaultsInput) (*request.Request, *glacier.ListVaultsOutput)

	ListVaults(*glacier.ListVaultsInput) (*glacier.ListVaultsOutput, error)

	ListVaultsPages(*glacier.ListVaultsInput, func(*glacier.ListVaultsOutput, bool) bool) error

	RemoveTagsFromVaultRequest(*glacier.RemoveTagsFromVaultInput) (*request.Request, *glacier.RemoveTagsFromVaultOutput)

	RemoveTagsFromVault(*glacier.RemoveTagsFromVaultInput) (*glacier.RemoveTagsFromVaultOutput, error)

	SetDataRetrievalPolicyRequest(*glacier.SetDataRetrievalPolicyInput) (*request.Request, *glacier.SetDataRetrievalPolicyOutput)

	SetDataRetrievalPolicy(*glacier.SetDataRetrievalPolicyInput) (*glacier.SetDataRetrievalPolicyOutput, error)

	SetVaultAccessPolicyRequest(*glacier.SetVaultAccessPolicyInput) (*request.Request, *glacier.SetVaultAccessPolicyOutput)

	SetVaultAccessPolicy(*glacier.SetVaultAccessPolicyInput) (*glacier.SetVaultAccessPolicyOutput, error)

	SetVaultNotificationsRequest(*glacier.SetVaultNotificationsInput) (*request.Request, *glacier.SetVaultNotificationsOutput)

	SetVaultNotifications(*glacier.SetVaultNotificationsInput) (*glacier.SetVaultNotificationsOutput, error)

	UploadArchiveRequest(*glacier.UploadArchiveInput) (*request.Request, *glacier.ArchiveCreationOutput)

	UploadArchive(*glacier.UploadArchiveInput) (*glacier.ArchiveCreationOutput, error)

	UploadMultipartPartRequest(*glacier.UploadMultipartPartInput) (*request.Request, *glacier.UploadMultipartPartOutput)

	UploadMultipartPart(*glacier.UploadMultipartPartInput) (*glacier.UploadMultipartPartOutput, error)
}
