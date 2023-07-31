package tracing

import (
	"errors"
	"io"
	"log"

	"github.com/opentracing/opentracing-go"
	jaegerConfig "github.com/uber/jaeger-client-go/config"
)

type JaegerTracer struct {
	Tracer opentracing.Tracer
	Closer io.Closer
}

// Return a new Jaeger Tracer
func NewTracer(serviceName, agentHostPort string) JaegerTracer {
	cfg := jaegerConfig.Configuration{
		ServiceName: serviceName,
		Reporter: &jaegerConfig.ReporterConfig{
			LogSpans: false,
			LocalAgentHostPort: agentHostPort,
		},
		Sampler: &jaegerConfig.SamplerConfig{
			Type: "const",
			Param: 1,
		},
	}

	t, c, err := cfg.NewTracer()
	if err != nil {
		panic("error while initial a Tracer")
	}

	return JaegerTracer{
		Tracer: t,
		Closer: c,
	}
}

// Return parent span(or a new span if no parent) and error
func (j *JaegerTracer) getParentSpan(operationName, traceID string, isRoot bool) (opentracing.Span, error) {
	if j.Tracer == nil {
		log.Println("Tracer is nil")
		return nil, errors.New("tracer is nil")
	}

	parentSpanCtx, err := j.Tracer.Extract(opentracing.TextMap, opentracing.TextMapCarrier{"UBER-TRACE-ID": traceID})
	if err != nil {
		if isRoot {
			return j.Tracer.StartSpan(operationName), nil
		}
		log.Println("error while extract: ", err)
		return nil, err
	}

	return j.Tracer.StartSpan(operationName, opentracing.ChildOf(parentSpanCtx)), nil
}

func (j *JaegerTracer) StartSpan(operationName, parentTracID string, isRoot bool) (opentracing.Span, string, error) {
	parentSpan, err := j.getParentSpan(operationName, parentTracID, isRoot)
	if err != nil {
		log.Println("no span return: ", err)
		return nil, "", errors.New("error while get parent span")
	}

	carrier := opentracing.TextMapCarrier{}
	err = j.Tracer.Inject(parentSpan.Context(), opentracing.TextMap, carrier)
	if err != nil {
		return nil, "", err
	}

	return parentSpan, carrier["uber-trace-id"], nil
}

func (j *JaegerTracer) FinishSpan(span opentracing.Span) {
	if span != nil {
		span.Finish()
	}
}

func (j *JaegerTracer) SpanSetTag(span opentracing.Span, tagKey, tagValue string) {
	if span != nil {
		span.SetTag(tagKey, tagValue)
	}
}