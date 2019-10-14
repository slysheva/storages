package s3

import (
	"github.com/aws/aws-sdk-go/aws/request"
)

type ThrottlerRetries struct {
	// RetryRules return the retry delay that should be used by the SDK before
	// making another request attempt for the failed request.
	RetryRules(*Request) time.Duration

	// ShouldRetry returns if the failed request is retryable.
	//
	// Implementations may consider request attempt count when determining if a
	// request is retryable, but the SDK will use MaxRetries to limit the
	// number of attempts a request are made.
	ShouldRetry(*Request) bool

	// MaxRetries is the number of times a request may be retried before
	// failing.
	MaxRetries() int
}



func New() ThrottlerRetries{
	return ThrottlerRetries{}
}