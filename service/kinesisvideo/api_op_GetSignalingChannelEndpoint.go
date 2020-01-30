// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kinesisvideo

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type GetSignalingChannelEndpointInput struct {
	_ struct{} `type:"structure"`

	// The ARN of the signalling channel for which you want to get an endpoint.
	//
	// ChannelARN is a required field
	ChannelARN *string `min:"1" type:"string" required:"true"`

	// A structure containing the endpoint configuration for the SINGLE_MASTER channel
	// type.
	SingleMasterChannelEndpointConfiguration *SingleMasterChannelEndpointConfiguration `type:"structure"`
}

// String returns the string representation
func (s GetSignalingChannelEndpointInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetSignalingChannelEndpointInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetSignalingChannelEndpointInput"}

	if s.ChannelARN == nil {
		invalidParams.Add(aws.NewErrParamRequired("ChannelARN"))
	}
	if s.ChannelARN != nil && len(*s.ChannelARN) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ChannelARN", 1))
	}
	if s.SingleMasterChannelEndpointConfiguration != nil {
		if err := s.SingleMasterChannelEndpointConfiguration.Validate(); err != nil {
			invalidParams.AddNested("SingleMasterChannelEndpointConfiguration", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetSignalingChannelEndpointInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ChannelARN != nil {
		v := *s.ChannelARN

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "ChannelARN", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SingleMasterChannelEndpointConfiguration != nil {
		v := s.SingleMasterChannelEndpointConfiguration

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "SingleMasterChannelEndpointConfiguration", v, metadata)
	}
	return nil
}

type GetSignalingChannelEndpointOutput struct {
	_ struct{} `type:"structure"`

	// A list of endpoints for the specified signaling channel.
	ResourceEndpointList []ResourceEndpointListItem `type:"list"`
}

// String returns the string representation
func (s GetSignalingChannelEndpointOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetSignalingChannelEndpointOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ResourceEndpointList != nil {
		v := s.ResourceEndpointList

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "ResourceEndpointList", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opGetSignalingChannelEndpoint = "GetSignalingChannelEndpoint"

// GetSignalingChannelEndpointRequest returns a request value for making API operation for
// Amazon Kinesis Video Streams.
//
// Provides an endpoint for the specified signaling channel to send and receive
// messages. This API uses the SingleMasterChannelEndpointConfiguration input
// parameter, which consists of the Protocols and Role properties.
//
// Protocols is used to determine the communication mechanism. For example,
// specifying WSS as the protocol, results in this API producing a secure websocket
// endpoint, and specifying HTTPS as the protocol, results in this API generating
// an HTTPS endpoint.
//
// Role determines the messaging permissions. A MASTER role results in this
// API generating an endpoint that a client can use to communicate with any
// of the viewers on the channel. A VIEWER role results in this API generating
// an endpoint that a client can use to communicate only with a MASTER.
//
//    // Example sending a request using GetSignalingChannelEndpointRequest.
//    req := client.GetSignalingChannelEndpointRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kinesisvideo-2017-09-30/GetSignalingChannelEndpoint
func (c *Client) GetSignalingChannelEndpointRequest(input *GetSignalingChannelEndpointInput) GetSignalingChannelEndpointRequest {
	op := &aws.Operation{
		Name:       opGetSignalingChannelEndpoint,
		HTTPMethod: "POST",
		HTTPPath:   "/getSignalingChannelEndpoint",
	}

	if input == nil {
		input = &GetSignalingChannelEndpointInput{}
	}

	req := c.newRequest(op, input, &GetSignalingChannelEndpointOutput{})
	return GetSignalingChannelEndpointRequest{Request: req, Input: input, Copy: c.GetSignalingChannelEndpointRequest}
}

// GetSignalingChannelEndpointRequest is the request type for the
// GetSignalingChannelEndpoint API operation.
type GetSignalingChannelEndpointRequest struct {
	*aws.Request
	Input *GetSignalingChannelEndpointInput
	Copy  func(*GetSignalingChannelEndpointInput) GetSignalingChannelEndpointRequest
}

// Send marshals and sends the GetSignalingChannelEndpoint API request.
func (r GetSignalingChannelEndpointRequest) Send(ctx context.Context) (*GetSignalingChannelEndpointResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetSignalingChannelEndpointResponse{
		GetSignalingChannelEndpointOutput: r.Request.Data.(*GetSignalingChannelEndpointOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetSignalingChannelEndpointResponse is the response type for the
// GetSignalingChannelEndpoint API operation.
type GetSignalingChannelEndpointResponse struct {
	*GetSignalingChannelEndpointOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetSignalingChannelEndpoint request.
func (r *GetSignalingChannelEndpointResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}