package libs_tracing

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/opentracing/opentracing-go/log"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"io"
	"net/http"
)

var tracing TracingJaegerOperation

type InterfaceTracingJaeger interface {
	InitJaeger() (opentracing.Tracer, io.Closer)
	GeTracer() opentracing.Tracer
}

type TracingJaeger struct {
	serviceName string
	tracer      opentracing.Tracer
}

func NewTracingJaeger(service string) InterfaceTracingJaeger {
	return &TracingJaeger{
		serviceName: service,
	}
}
func (c *TracingJaeger) InitJaeger() (opentracing.Tracer, io.Closer) {
	if c.serviceName == "" {
		c.serviceName = "Init ServiceName"
	}
	cfg := &config.Configuration{
		ServiceName: c.serviceName,
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1000,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: "localhost:6831",
		},
	}
	tracer, closer, err := cfg.NewTracer(
		config.Logger(jaeger.StdLogger),
	)
	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
	}
	opentracing.SetGlobalTracer(tracer)
	c.tracer = tracer
	return tracer, closer
}
func (c *TracingJaeger) GeTracer() opentracing.Tracer {
	return c.tracer
}

type InterfaceTracingJaegerOperation interface {
	SetOperationChild(name string) (span opentracing.Span, ctx context.Context)
	SetOperationParent(ctx2 context.Context, r *http.Request, name string) (span opentracing.Span, ctx context.Context)
	SetLog(key string, val interface{})
	TracingTag(request *http.Request)
	OutgoingContext() context.Context
	SetError(key string, val string)
	FinishChildOperation()
}

type TracingJaegerOperation struct {
	context       context.Context
	span          opentracing.Span
	metadata      *http.Request
	nameOperation string
}

func NewTracingJaegerOperation(ctx context.Context) InterfaceTracingJaegerOperation {
	return &TracingJaegerOperation{
		context: ctx,
	}
}

func (c *TracingJaegerOperation) TracingTag(request *http.Request) {
	if request == nil || c.span == nil {
		fmt.Println("Please call SetOperationParent First")
		return
	}
	c.metadata = request
	for i, j := range c.metadata.Header {
		c.span.SetTag(i, j)
	}
	c.span.SetTag("url", c.metadata.URL.String())
	c.span.SetTag("method", c.metadata.Method)
	c.span.SetTag("host", c.metadata.Host)
	c.span.SetTag("x-trace-id", uuid.New().String())

}
func (c *TracingJaegerOperation) OutgoingContext() context.Context {
	if c.context == nil {
		return context.Background()
	}
	return c.context
}

func (c *TracingJaegerOperation) SetTag(key string, val string) {
	marshal, err := ConvertJSON(val)
	if err != nil {
		return
	}
	c.span.SetTag(key, marshal)
}

func (c *TracingJaegerOperation) SetError(key string, val string) {
	marshal, err := ConvertJSON(val)
	if err != nil {
		return
	}
	c.span.SetTag("error", true)
	c.SetLog(key+"=", marshal)
}

func ConvertJSON(val interface{}) (string, error) {
	marshal, err := json.Marshal(val)
	if err != nil {
		fmt.Println("ConvertJSON=" + err.Error())
		return "", err
	}
	return string(marshal), err
}
func (c *TracingJaegerOperation) SetLog(key string, val interface{}) {
	marshal, err := ConvertJSON(val)
	if err != nil {
		return
	}
	c.span.LogFields(log.Object(key+"=", marshal))
}

func (c *TracingJaegerOperation) SetOperationParent(ctx2 context.Context, r *http.Request, name string) (span opentracing.Span, ctx context.Context) {
	c.nameOperation = name
	c.context = ctx2
	// Memulai span dari konteks yang diberikan
	span, ctx = opentracing.StartSpanFromContext(c.context, r.URL.String()+"."+name)
	// Update span saat ini dan konteks di objek TracingJaegerOperation
	c.span = span
	c.context = ctx
	c.TracingTag(r)
	extract, err := Extract(span.Tracer(), r)
	if err != nil {
		fmt.Println("error SetOperationParent")
		return
	}
	ext.RPCServerOption(extract)
	ext.HTTPMethod.Set(span, "POST")
	ext.HTTPUrl.Set(span, r.URL.String())

	return c.span, c.context
}

func (c *TracingJaegerOperation) SetOperationChild(name string) (span opentracing.Span, ctx context.Context) {
	// Memulai span child dari konteks span parent yang tersimpan dalam objek
	span, ctx = opentracing.StartSpanFromContext(c.context, name, opentracing.FollowsFrom(c.span.Context()))
	// Mengatur span child sebagai span saat ini dalam objek
	c.span = span
	// Mengembalikan span child dan konteksnya
	c.TracingTag(c.metadata)
	return c.span, ctx
}

func (c *TracingJaegerOperation) FinishChildOperation() {
	// Selesaikan operasi pada span child
	if c.span == nil {
		fmt.Println("No child operation is running.")
		return
	}
	defer c.span.Finish()
	// Kembalikan span parent ke konteks dan span yang saat ini dipegang oleh objek TracingJaegerOperation
	c.span = opentracing.SpanFromContext(c.context)
}

func (c *TracingJaegerOperation) SetOperationSubChild(name string, con context.Context) (span opentracing.Span, ctx context.Context) {
	// Memulai span child dari konteks parent yang diberikan

	span, ctx = opentracing.StartSpanFromContext(con, name, opentracing.FollowsFrom(c.span.Context()))
	// Mengatur span child sebagai span saat ini dalam objek
	c.span = span
	// Mengembalikan span child dan konteksnya
	c.TracingTag(c.metadata)
	return c.span, ctx
}

func (c *TracingJaegerOperation) FinishSubChildOperation(con context.Context) {
	// Selesaikan operasi pada span child
	if c.span == nil {
		fmt.Println("No child operation is running.")
		return
	}
	defer c.span.Finish()
	// Kembalikan span parent ke konteks dan span yang saat ini dipegang oleh objek TracingJaegerOperation
	c.span = opentracing.SpanFromContext(con)
}

func StartSpanFromRequest(tracer opentracing.Tracer, r *http.Request, funcDesc string) opentracing.Span {
	spanCtx, _ := Extract(tracer, r)
	return tracer.StartSpan(funcDesc, ext.RPCServerOption(spanCtx))
}
func Inject(span opentracing.Span, request *http.Request) error {
	return span.Tracer().Inject(
		span.Context(),
		opentracing.HTTPHeaders,
		opentracing.HTTPHeadersCarrier(request.Header))
}

func Extract(tracer opentracing.Tracer, r *http.Request) (opentracing.SpanContext, error) {
	return tracer.Extract(
		opentracing.HTTPHeaders,
		opentracing.HTTPHeadersCarrier(r.Header))
}
