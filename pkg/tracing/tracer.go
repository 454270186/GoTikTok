package tracing

import (
	"io"

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

func getParentSpan(operationName, traceID string, isRoot bool) (opentracing.Span, error) {
	// TODO
	return nil, nil
}

func StartSpan(operationName, parentTracID string, isRoot bool) (opentracing.Span, string, error) {
	// TODO
	return nil, "", nil
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