package tracing

import (
	"fmt"

	"github.com/opentracing/opentracing-go"
)

func (j *JaegerTracer) Gin(ginOperationName, parentTraceID string, isRoot bool) (opentracing.Span, string, error) {
	operationName := fmt.Sprintf("Gin: %s", ginOperationName)
	return j.StartSpan(operationName, parentTraceID, isRoot)
}

func (j *JaegerTracer) RPC(rpcOperationName, parentTraceID string, isRoot bool) (opentracing.Span, string, error) {
	operationName := fmt.Sprintf("Remote: %s", rpcOperationName)
	return j.StartSpan(operationName, parentTraceID, isRoot)
}

func (j *JaegerTracer) DB(dbOperationName, parentTraceID string, isRoot bool) (opentracing.Span, string, error) {
	operationName := fmt.Sprintf("DB: %s", dbOperationName)
	return j.StartSpan(operationName, parentTraceID, isRoot)
}