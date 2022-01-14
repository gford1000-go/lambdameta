package lamdbameta

import (
	"context"
	"time"

	"github.com/aws/aws-lambda-go/lambdacontext"
)

type LambdaFunction struct {
	Name    string `json:"function_name"`
	Version string `json:"function_version"`
}

type LogInfo struct {
	GroupName  string `json:"group_name"`
	StreamName string `json:"stream_name"`
}

type Runtime struct {
	MemoryLimit int `json:"memory_limit"`
}

type Meta struct {
	Function  LambdaFunction `json:"lambda"`
	Logging   LogInfo        `json:"logging"`
	Runtime   Runtime        `json:"runtime"`
	RequestID string         `json:"aws_request_id"`
	Timeout   string         `json:"request_timeout"`
}

// GetMeta returns metadata information from the supplied context
func GetMeta(ctx context.Context) *Meta {
	t, _ := ctx.Deadline()
	lc, _ := lambdacontext.FromContext(ctx)
	return &Meta{
		Function: LambdaFunction{
			Name:    lambdacontext.FunctionName,
			Version: lambdacontext.FunctionVersion,
		},
		Logging: LogInfo{
			GroupName:  lambdacontext.LogGroupName,
			StreamName: lambdacontext.LogStreamName,
		},
		Runtime: Runtime{
			MemoryLimit: lambdacontext.MemoryLimitInMB,
		},
		RequestID: lc.AwsRequestID,
		Timeout:   t.Format(time.RFC3339Nano),
	}
}
