// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package comprehendmedical

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type StopRxNormInferenceJobInput struct {
	_ struct{} `type:"structure"`

	// The identifier of the job.
	//
	// JobId is a required field
	JobId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s StopRxNormInferenceJobInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StopRxNormInferenceJobInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StopRxNormInferenceJobInput"}

	if s.JobId == nil {
		invalidParams.Add(aws.NewErrParamRequired("JobId"))
	}
	if s.JobId != nil && len(*s.JobId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("JobId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type StopRxNormInferenceJobOutput struct {
	_ struct{} `type:"structure"`

	// The identifier generated for the job. To get the status of job, use this
	// identifier with the DescribeRxNormInferenceJob operation.
	JobId *string `min:"1" type:"string"`
}

// String returns the string representation
func (s StopRxNormInferenceJobOutput) String() string {
	return awsutil.Prettify(s)
}

const opStopRxNormInferenceJob = "StopRxNormInferenceJob"

// StopRxNormInferenceJobRequest returns a request value for making API operation for
// AWS Comprehend Medical.
//
// Stops an InferRxNorm inference job in progress.
//
//    // Example sending a request using StopRxNormInferenceJobRequest.
//    req := client.StopRxNormInferenceJobRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/comprehendmedical-2018-10-30/StopRxNormInferenceJob
func (c *Client) StopRxNormInferenceJobRequest(input *StopRxNormInferenceJobInput) StopRxNormInferenceJobRequest {
	op := &aws.Operation{
		Name:       opStopRxNormInferenceJob,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &StopRxNormInferenceJobInput{}
	}

	req := c.newRequest(op, input, &StopRxNormInferenceJobOutput{})

	return StopRxNormInferenceJobRequest{Request: req, Input: input, Copy: c.StopRxNormInferenceJobRequest}
}

// StopRxNormInferenceJobRequest is the request type for the
// StopRxNormInferenceJob API operation.
type StopRxNormInferenceJobRequest struct {
	*aws.Request
	Input *StopRxNormInferenceJobInput
	Copy  func(*StopRxNormInferenceJobInput) StopRxNormInferenceJobRequest
}

// Send marshals and sends the StopRxNormInferenceJob API request.
func (r StopRxNormInferenceJobRequest) Send(ctx context.Context) (*StopRxNormInferenceJobResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StopRxNormInferenceJobResponse{
		StopRxNormInferenceJobOutput: r.Request.Data.(*StopRxNormInferenceJobOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StopRxNormInferenceJobResponse is the response type for the
// StopRxNormInferenceJob API operation.
type StopRxNormInferenceJobResponse struct {
	*StopRxNormInferenceJobOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StopRxNormInferenceJob request.
func (r *StopRxNormInferenceJobResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}