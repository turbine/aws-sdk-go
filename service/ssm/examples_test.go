// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package ssm_test

import (
	"bytes"
	"fmt"
	"time"

	"github.com/turbine/aws-sdk-go/aws"
	"github.com/turbine/aws-sdk-go/service/ssm"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleSSM_CreateAssociation() {
	svc := ssm.New(nil)

	params := &ssm.CreateAssociationInput{
		InstanceId: aws.String("InstanceId"),   // Required
		Name:       aws.String("DocumentName"), // Required
	}
	resp, err := svc.CreateAssociation(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSSM_CreateAssociationBatch() {
	svc := ssm.New(nil)

	params := &ssm.CreateAssociationBatchInput{
		Entries: []*ssm.CreateAssociationBatchRequestEntry{ // Required
			{ // Required
				InstanceId: aws.String("InstanceId"),
				Name:       aws.String("DocumentName"),
			},
			// More values...
		},
	}
	resp, err := svc.CreateAssociationBatch(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSSM_CreateDocument() {
	svc := ssm.New(nil)

	params := &ssm.CreateDocumentInput{
		Content: aws.String("DocumentContent"), // Required
		Name:    aws.String("DocumentName"),    // Required
	}
	resp, err := svc.CreateDocument(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSSM_DeleteAssociation() {
	svc := ssm.New(nil)

	params := &ssm.DeleteAssociationInput{
		InstanceId: aws.String("InstanceId"),   // Required
		Name:       aws.String("DocumentName"), // Required
	}
	resp, err := svc.DeleteAssociation(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSSM_DeleteDocument() {
	svc := ssm.New(nil)

	params := &ssm.DeleteDocumentInput{
		Name: aws.String("DocumentName"), // Required
	}
	resp, err := svc.DeleteDocument(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSSM_DescribeAssociation() {
	svc := ssm.New(nil)

	params := &ssm.DescribeAssociationInput{
		InstanceId: aws.String("InstanceId"),   // Required
		Name:       aws.String("DocumentName"), // Required
	}
	resp, err := svc.DescribeAssociation(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSSM_DescribeDocument() {
	svc := ssm.New(nil)

	params := &ssm.DescribeDocumentInput{
		Name: aws.String("DocumentName"), // Required
	}
	resp, err := svc.DescribeDocument(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSSM_GetDocument() {
	svc := ssm.New(nil)

	params := &ssm.GetDocumentInput{
		Name: aws.String("DocumentName"), // Required
	}
	resp, err := svc.GetDocument(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSSM_ListAssociations() {
	svc := ssm.New(nil)

	params := &ssm.ListAssociationsInput{
		AssociationFilterList: []*ssm.AssociationFilter{ // Required
			{ // Required
				Key:   aws.String("AssociationFilterKey"),   // Required
				Value: aws.String("AssociationFilterValue"), // Required
			},
			// More values...
		},
		MaxResults: aws.Int64(1),
		NextToken:  aws.String("NextToken"),
	}
	resp, err := svc.ListAssociations(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSSM_ListDocuments() {
	svc := ssm.New(nil)

	params := &ssm.ListDocumentsInput{
		DocumentFilterList: []*ssm.DocumentFilter{
			{ // Required
				Key:   aws.String("DocumentFilterKey"),   // Required
				Value: aws.String("DocumentFilterValue"), // Required
			},
			// More values...
		},
		MaxResults: aws.Int64(1),
		NextToken:  aws.String("NextToken"),
	}
	resp, err := svc.ListDocuments(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSSM_UpdateAssociationStatus() {
	svc := ssm.New(nil)

	params := &ssm.UpdateAssociationStatusInput{
		AssociationStatus: &ssm.AssociationStatus{ // Required
			Date:           aws.Time(time.Now()),                // Required
			Message:        aws.String("StatusMessage"),         // Required
			Name:           aws.String("AssociationStatusName"), // Required
			AdditionalInfo: aws.String("StatusAdditionalInfo"),
		},
		InstanceId: aws.String("InstanceId"),   // Required
		Name:       aws.String("DocumentName"), // Required
	}
	resp, err := svc.UpdateAssociationStatus(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}
