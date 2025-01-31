package aperture

import (
	"errors"
	"time"

	checkproto "buf.build/gen/go/fluxninja/aperture/protocolbuffers/go/aperture/flowcontrol/check/v1"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/protobuf/encoding/protojson"
)

// Flow is the interface that is returned to the user every time a CheckHTTP call through ApertureClient is made.
// The user can check the status of the check call, response from the server, and end the flow once the workload is executed.
type Flow interface {
	ShouldRun() bool
	SetStatus(status FlowStatus)
	Error() error
	Span() trace.Span
	End() error
	CheckResponse() *checkproto.CheckResponse
}

// TODO: set fail open?

type flow struct {
	span          trace.Span
	err           error
	checkResponse *checkproto.CheckResponse
	statusCode    FlowStatus
	ended         bool
	failOpen      bool
}

// newFlow creates a new flow with default field values.
func newFlow(span trace.Span, failOpen bool) *flow {
	return &flow{
		span:          span,
		checkResponse: nil,
		statusCode:    OK,
		ended:         false,
		failOpen:      failOpen,
	}
}

// ShouldRun returns whether the Flow was allowed to run by Aperture Agent.
// By default, fail-open behavior is enabled. Use DisableFailOpen to disable it.
func (f *flow) ShouldRun() bool {
	if (f.failOpen && f.checkResponse == nil) || (f.checkResponse.DecisionType == checkproto.CheckResponse_DECISION_TYPE_ACCEPTED) {
		return true
	} else {
		return false
	}
}

// DisableFailOpen disables fail-open behavior for the flow.
func (f *flow) DisableFailOpen() {
	f.failOpen = false
}

// CheckResponse returns the response from the server.
func (f *flow) CheckResponse() *checkproto.CheckResponse {
	return f.checkResponse
}

// SetStatus sets the status code of a flow.
// If not set explicitly, defaults to FlowStatus.OK.
func (f *flow) SetStatus(statusCode FlowStatus) {
	f.statusCode = statusCode
}

// Error returns the error that occurred during the flow.
func (f *flow) Error() error {
	return f.err
}

// Span returns the span associated with the flow.
func (f *flow) Span() trace.Span {
	return f.span
}

// End is used to end the flow, using the status code previously set using SetStatus method.
func (f *flow) End() error {
	if f.ended {
		return errors.New("flow already ended")
	}
	f.ended = true

	checkResponseJSONBytes, err := protojson.Marshal(f.checkResponse)
	if err != nil {
		return err
	}
	f.span.SetAttributes(
		attribute.String(flowStatusLabel, f.statusCode.String()),
		attribute.String(checkResponseLabel, string(checkResponseJSONBytes)),
		attribute.Int64(flowEndTimestampLabel, time.Now().UnixNano()),
	)
	f.span.End()
	return nil
}
