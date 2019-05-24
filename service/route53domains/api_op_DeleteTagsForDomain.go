// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package route53domains

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// The DeleteTagsForDomainRequest includes the following elements.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53domains-2014-05-15/DeleteTagsForDomainRequest
type DeleteTagsForDomainInput struct {
	_ struct{} `type:"structure"`

	// The domain for which you want to delete one or more tags.
	//
	// DomainName is a required field
	DomainName *string `type:"string" required:"true"`

	// A list of tag keys to delete.
	//
	// TagsToDelete is a required field
	TagsToDelete []string `type:"list" required:"true"`
}

// String returns the string representation
func (s DeleteTagsForDomainInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteTagsForDomainInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteTagsForDomainInput"}

	if s.DomainName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DomainName"))
	}

	if s.TagsToDelete == nil {
		invalidParams.Add(aws.NewErrParamRequired("TagsToDelete"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53domains-2014-05-15/DeleteTagsForDomainResponse
type DeleteTagsForDomainOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteTagsForDomainOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteTagsForDomain = "DeleteTagsForDomain"

// DeleteTagsForDomainRequest returns a request value for making API operation for
// Amazon Route 53 Domains.
//
// This operation deletes the specified tags for a domain.
//
// All tag operations are eventually consistent; subsequent operations might
// not immediately represent all issued operations.
//
//    // Example sending a request using DeleteTagsForDomainRequest.
//    req := client.DeleteTagsForDomainRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53domains-2014-05-15/DeleteTagsForDomain
func (c *Client) DeleteTagsForDomainRequest(input *DeleteTagsForDomainInput) DeleteTagsForDomainRequest {
	op := &aws.Operation{
		Name:       opDeleteTagsForDomain,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteTagsForDomainInput{}
	}

	req := c.newRequest(op, input, &DeleteTagsForDomainOutput{})
	return DeleteTagsForDomainRequest{Request: req, Input: input, Copy: c.DeleteTagsForDomainRequest}
}

// DeleteTagsForDomainRequest is the request type for the
// DeleteTagsForDomain API operation.
type DeleteTagsForDomainRequest struct {
	*aws.Request
	Input *DeleteTagsForDomainInput
	Copy  func(*DeleteTagsForDomainInput) DeleteTagsForDomainRequest
}

// Send marshals and sends the DeleteTagsForDomain API request.
func (r DeleteTagsForDomainRequest) Send(ctx context.Context) (*DeleteTagsForDomainResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteTagsForDomainResponse{
		DeleteTagsForDomainOutput: r.Request.Data.(*DeleteTagsForDomainOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteTagsForDomainResponse is the response type for the
// DeleteTagsForDomain API operation.
type DeleteTagsForDomainResponse struct {
	*DeleteTagsForDomainOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteTagsForDomain request.
func (r *DeleteTagsForDomainResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}