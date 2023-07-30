package tracing

import (
	"errors"
	"io"
	"log"

	"github.com/opentracing/opentracing-go"
	jaegerConfig "github.com/uber/jaeger-client-go/config"
)

var (
	Tracer opentracing.Tracer
	Closer io.Closer
)

// Init a Jaeger tracer
func InitTracer(serviceName, agentHostPort string) error {
	if Tracer != nil && Closer != nil {
		return nil
	}

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

	Tracer, Closer = t, c
	return nil
}

// Return parent span(or a new span if no parent) and error
func getParentSpan(operationName, traceID string, isRoot bool) (opentracing.Span, error) {
	if Tracer == nil {
		log.Println("Tracer is nil")
		return nil, errors.New("tracer is nil")
	}

	parentSpanCtx, err := Tracer.Extract(opentracing.TextMap, opentracing.TextMapCarrier{"UBER-TRACE-ID": traceID})
	if err != nil {
		if isRoot {
			return Tracer.StartSpan(operationName), nil
		}
		log.Println("error while extract: ", err)
		return nil, err
	}

	return Tracer.StartSpan(operationName, opentracing.ChildOf(parentSpanCtx)), nil
}

func StartSpan(operationName, parentTracID string, isRoot bool) (opentracing.Span, string, error) {
	parentSpan, err := getParentSpan(operationName, parentTracID, isRoot)
	if err != nil {
		log.Println("no span return: ", err)
		return nil, "", errors.New("error while get parent span")
	}

	carrier := opentracing.TextMapCarrier{}
	err = Tracer.Inject(parentSpan.Context(), opentracing.TextMap, carrier)
	if err != nil {
		return nil, "", err
	}


	return parentSpan, carrier["uber-trace-id"], nil
}

func FinishSpan(span opentracing.Span) {
	if span != nil {
		span.Finish()
	}
}

func SpanSetTag(span opentracing.Span, tagKey, tagValue string) {
	if span != nil {
		span.SetTag(tagKey, tagValue)
	}
}