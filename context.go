package spanner

import (
	"context"
)

// SpannerContextKey is a unique type to use as key when storing traces in context
type SpannerContextKey string

// ContextKey is the package's default context key for storing ContextData in context
var ContextKey SpannerContextKey = "spanner"

type SpanContext context.Context

type ContextData struct {
	Trace *TraceID
	Span  *SpanID
}

type spanDataContext struct {
	context.Context
	value ContextData
}

func (c *spanDataContext) Value(key any) any {
	if key == ContextKey {
		return c.value
	}
	return c.Context.Value(key)
}

func WithContextData(ctx context.Context, cd *ContextData) context.Context {
	if cd != nil {
		return &spanDataContext{
			Context: ctx,
			value:   *cd,
		}
	}
	return ctx
}

func GetContextData(ctx context.Context) *ContextData {
	if c, ok := ctx.(*spanDataContext); ok {
		return &c.value
	}
	return nil
}

// // TraceContextKey is a unique type to use as key when storing traces in context
// type TraceContextKey string

// // ContextKey is the package's default context key for storing traces in context
// var ContextKey TraceContextKey = "spanner"

// // ContextData stores a reference to the Trace and (parent) Span to pass along a context
// type ContextData struct {
// 	Trace *TraceID
// 	Span  *SpanID
// }

// func WithContextData(ctx context.Context, cd *ContextData) context.Context {
// 	return context.WithValue(ctx, ContextKey, *cd)
// }

// // GetTrace returns the Trace from the input context `ctx`, or nil if it doesn't have one
// func GetTrace(ctx context.Context) *TraceID {
// 	v := ctx.Value(ContextKey)
// 	if v == nil {
// 		return nil
// 	}
// 	if sd, ok := v.(ContextData); ok {
// 		return sd.Trace

// 	}
// 	return nil
// }

// // GetSpan returns the Span from the input context `ctx`, or nil if it doesn't have one
// func GetSpan(ctx context.Context) *SpanID {
// 	v := ctx.Value(ContextKey)
// 	if v == nil {
// 		return nil
// 	}
// 	if sd, ok := v.(ContextData); ok {
// 		return sd.Span
// 	}
// 	return nil
// }

// func GetContextData(ctx context.Context) *ContextData {
// 	v := ctx.Value(ContextKey)
// 	if v == nil {
// 		return nil
// 	}
// 	if sd, ok := v.(ContextData); ok {
// 		return &sd
// 	}
// 	return nil
// }
