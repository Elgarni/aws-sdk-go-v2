// Code generated by smithy-go-codegen DO NOT EDIT.

package secretsmanager

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Retrieves the contents of the encrypted fields SecretString or SecretBinary from
// the specified version of a secret, whichever contains content. Minimum
// permissions To run this command, you must have the following permissions:
//
// *
// secretsmanager:GetSecretValue
//
// * kms:Decrypt - required only if you use a
// customer-managed Amazon Web Services KMS key to encrypt the secret. You do not
// need this permission to use the account's default Amazon Web Services managed
// CMK for Secrets Manager.
//
// Related operations
//
// * To create a new version of the
// secret with different encrypted information, use PutSecretValue.
//
// * To retrieve
// the non-encrypted details for the secret, use DescribeSecret.
func (c *Client) GetSecretValue(ctx context.Context, params *GetSecretValueInput, optFns ...func(*Options)) (*GetSecretValueOutput, error) {
	if params == nil {
		params = &GetSecretValueInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetSecretValue", params, optFns, c.addOperationGetSecretValueMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetSecretValueOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetSecretValueInput struct {

	// Specifies the secret containing the version that you want to retrieve. You can
	// specify either the Amazon Resource Name (ARN) or the friendly name of the
	// secret. If you specify an ARN, we generally recommend that you specify a
	// complete ARN. You can specify a partial ARN too—for example, if you don’t
	// include the final hyphen and six random characters that Secrets Manager adds at
	// the end of the ARN when you created the secret. A partial ARN match can work as
	// long as it uniquely matches only one secret. However, if your secret has a name
	// that ends in a hyphen followed by six characters (before Secrets Manager adds
	// the hyphen and six characters to the ARN) and you try to use that as a partial
	// ARN, then those characters cause Secrets Manager to assume that you’re
	// specifying a complete ARN. This confusion can cause unexpected results. To avoid
	// this situation, we recommend that you don’t create secret names ending with a
	// hyphen followed by six characters. If you specify an incomplete ARN without the
	// random suffix, and instead provide the 'friendly name', you must not include the
	// random suffix. If you do include the random suffix added by Secrets Manager, you
	// receive either a ResourceNotFoundException or an AccessDeniedException error,
	// depending on your permissions.
	//
	// This member is required.
	SecretId *string

	// Specifies the unique identifier of the version of the secret that you want to
	// retrieve. If you specify both this parameter and VersionStage, the two
	// parameters must refer to the same secret version. If you don't specify either a
	// VersionStage or VersionId then the default is to perform the operation on the
	// version with the VersionStage value of AWSCURRENT. This value is typically a
	// UUID-type (https://wikipedia.org/wiki/Universally_unique_identifier) value with
	// 32 hexadecimal digits.
	VersionId *string

	// Specifies the secret version that you want to retrieve by the staging label
	// attached to the version. Staging labels are used to keep track of different
	// versions during the rotation process. If you specify both this parameter and
	// VersionId, the two parameters must refer to the same secret version . If you
	// don't specify either a VersionStage or VersionId, then the default is to perform
	// the operation on the version with the VersionStage value of AWSCURRENT.
	VersionStage *string

	noSmithyDocumentSerde
}

type GetSecretValueOutput struct {

	// The ARN of the secret.
	ARN *string

	// The date and time that this version of the secret was created.
	CreatedDate *time.Time

	// The friendly name of the secret.
	Name *string

	// The decrypted part of the protected secret information that was originally
	// provided as binary data in the form of a byte array. The response parameter
	// represents the binary data as a base64-encoded
	// (https://tools.ietf.org/html/rfc4648#section-4) string. This parameter is not
	// used if the secret is created by the Secrets Manager console. If you store
	// custom information in this field of the secret, then you must code your Lambda
	// rotation function to parse and interpret whatever you store in the SecretString
	// or SecretBinary fields.
	SecretBinary []byte

	// The decrypted part of the protected secret information that was originally
	// provided as a string. If you create this secret by using the Secrets Manager
	// console then only the SecretString parameter contains data. Secrets Manager
	// stores the information as a JSON structure of key/value pairs that the Lambda
	// rotation function knows how to parse. If you store custom information in the
	// secret by using the CreateSecret, UpdateSecret, or PutSecretValue API operations
	// instead of the Secrets Manager console, or by using the Other secret type in the
	// console, then you must code your Lambda rotation function to parse and interpret
	// those values.
	SecretString *string

	// The unique identifier of this version of the secret.
	VersionId *string

	// A list of all of the staging labels currently attached to this version of the
	// secret.
	VersionStages []string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetSecretValueMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetSecretValue{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetSecretValue{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpGetSecretValueValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetSecretValue(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opGetSecretValue(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "secretsmanager",
		OperationName: "GetSecretValue",
	}
}
