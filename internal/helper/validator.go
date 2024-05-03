package helper

import (
	"context"
)

type Validator interface {
	// Valid function contains validating code for implementor
	Valid(ctx context.Context) (problems map[string]string)
}
