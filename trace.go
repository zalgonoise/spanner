package spanner

// // Trace represents a single transaction which creates a set of Spans, as single-event actions.
// //
// // It exposes methods for registering and retrieving the parent SpanID to use in the next
// // Tracer's `Start()` call, and for returning its TraceID.
// type Trace interface {
// 	// ID returns the TraceID
// 	ID() TraceID
// }

// type trace TraceID

// func newTrace() Trace {
// 	return trace(NewTraceID())
// }

// // ID returns the TraceID
// func (t trace) ID() TraceID {
// 	return TraceID(t)
// }
