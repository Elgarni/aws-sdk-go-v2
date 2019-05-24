// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type GetV2LoggingOptionsInput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s GetV2LoggingOptionsInput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetV2LoggingOptionsInput) MarshalFields(e protocol.FieldEncoder) error {

	return nil
}

type GetV2LoggingOptionsOutput struct {
	_ struct{} `type:"structure"`

	// The default log level.
	DefaultLogLevel LogLevel `locationName:"defaultLogLevel" type:"string" enum:"true"`

	// Disables all logs.
	DisableAllLogs *bool `locationName:"disableAllLogs" type:"boolean"`

	// The IAM role ARN AWS IoT uses to write to your CloudWatch logs.
	RoleArn *string `locationName:"roleArn" type:"string"`
}

// String returns the string representation
func (s GetV2LoggingOptionsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetV2LoggingOptionsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if len(s.DefaultLogLevel) > 0 {
		v := s.DefaultLogLevel

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "defaultLogLevel", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.DisableAllLogs != nil {
		v := *s.DisableAllLogs

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "disableAllLogs", protocol.BoolValue(v), metadata)
	}
	if s.RoleArn != nil {
		v := *s.RoleArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "roleArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opGetV2LoggingOptions = "GetV2LoggingOptions"

// GetV2LoggingOptionsRequest returns a request value for making API operation for
// AWS IoT.
//
// Gets the fine grained logging options.
//
//    // Example sending a request using GetV2LoggingOptionsRequest.
//    req := client.GetV2LoggingOptionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) GetV2LoggingOptionsRequest(input *GetV2LoggingOptionsInput) GetV2LoggingOptionsRequest {
	op := &aws.Operation{
		Name:       opGetV2LoggingOptions,
		HTTPMethod: "GET",
		HTTPPath:   "/v2LoggingOptions",
	}

	if input == nil {
		input = &GetV2LoggingOptionsInput{}
	}

	req := c.newRequest(op, input, &GetV2LoggingOptionsOutput{})
	return GetV2LoggingOptionsRequest{Request: req, Input: input, Copy: c.GetV2LoggingOptionsRequest}
}

// GetV2LoggingOptionsRequest is the request type for the
// GetV2LoggingOptions API operation.
type GetV2LoggingOptionsRequest struct {
	*aws.Request
	Input *GetV2LoggingOptionsInput
	Copy  func(*GetV2LoggingOptionsInput) GetV2LoggingOptionsRequest
}

// Send marshals and sends the GetV2LoggingOptions API request.
func (r GetV2LoggingOptionsRequest) Send(ctx context.Context) (*GetV2LoggingOptionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetV2LoggingOptionsResponse{
		GetV2LoggingOptionsOutput: r.Request.Data.(*GetV2LoggingOptionsOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetV2LoggingOptionsResponse is the response type for the
// GetV2LoggingOptions API operation.
type GetV2LoggingOptionsResponse struct {
	*GetV2LoggingOptionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetV2LoggingOptions request.
func (r *GetV2LoggingOptionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}