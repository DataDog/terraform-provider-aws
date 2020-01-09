// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudwatchlogs

const (

	// ErrCodeDataAlreadyAcceptedException for service response error code
	// "DataAlreadyAcceptedException".
	//
	// The event was already logged.
	ErrCodeDataAlreadyAcceptedException = "DataAlreadyAcceptedException"

	// ErrCodeInvalidOperationException for service response error code
	// "InvalidOperationException".
	//
	// The operation is not valid on the specified resource.
	ErrCodeInvalidOperationException = "InvalidOperationException"

	// ErrCodeInvalidParameterException for service response error code
	// "InvalidParameterException".
	//
	// A parameter is specified incorrectly.
	ErrCodeInvalidParameterException = "InvalidParameterException"

	// ErrCodeInvalidSequenceTokenException for service response error code
	// "InvalidSequenceTokenException".
	//
	// The sequence token is not valid. You can get the correct sequence token in
	// the expectedSequenceToken field in the InvalidSequenceTokenException message.
	ErrCodeInvalidSequenceTokenException = "InvalidSequenceTokenException"

	// ErrCodeLimitExceededException for service response error code
	// "LimitExceededException".
	//
	// You have reached the maximum number of resources that can be created.
	ErrCodeLimitExceededException = "LimitExceededException"

	// ErrCodeMalformedQueryException for service response error code
	// "MalformedQueryException".
	//
	// The query string is not valid. Details about this error are displayed in
	// a QueryCompileError object. For more information, see .
	//
	// For more information about valid query syntax, see CloudWatch Logs Insights
	// Query Syntax (https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CWL_QuerySyntax.html).
	ErrCodeMalformedQueryException = "MalformedQueryException"

	// ErrCodeOperationAbortedException for service response error code
	// "OperationAbortedException".
	//
	// Multiple requests to update the same resource were in conflict.
	ErrCodeOperationAbortedException = "OperationAbortedException"

	// ErrCodeResourceAlreadyExistsException for service response error code
	// "ResourceAlreadyExistsException".
	//
	// The specified resource already exists.
	ErrCodeResourceAlreadyExistsException = "ResourceAlreadyExistsException"

	// ErrCodeResourceNotFoundException for service response error code
	// "ResourceNotFoundException".
	//
	// The specified resource does not exist.
	ErrCodeResourceNotFoundException = "ResourceNotFoundException"

	// ErrCodeServiceUnavailableException for service response error code
	// "ServiceUnavailableException".
	//
	// The service cannot complete the request.
	ErrCodeServiceUnavailableException = "ServiceUnavailableException"

	// ErrCodeUnrecognizedClientException for service response error code
	// "UnrecognizedClientException".
	//
	// The most likely cause is an invalid AWS access key ID or secret key.
	ErrCodeUnrecognizedClientException = "UnrecognizedClientException"
)
