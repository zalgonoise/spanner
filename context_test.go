package spanner

import (
	"context"
	"testing"
)

func TestWithNewTrace(t *testing.T) {
	ctx, tr := WithNewTrace(context.Background())

	if tr == nil || ctx == nil {
		t.Errorf("unexpected nil value")
		return
	}

	extracted := ctx.Value(ContextKey)
	if extracted == nil {
		t.Errorf("unexpected nil value")
		return
	}
	extrTrace, ok := extracted.(Trace)
	if !ok {
		t.Errorf("unexpected type")
		return
	}
	if tr.ID() != extrTrace.ID() {
		t.Errorf("trace ID mismatch error: wanted %s ; got %s", tr.ID().String(), extrTrace.ID().String())
		return
	}
}

func TestWithTrace(t *testing.T) {
	tr := newTrace()

	ctx := WithTrace(context.Background(), tr)
	extracted := ctx.Value(ContextKey)
	if extracted == nil {
		t.Errorf("unexpected nil value")
		return
	}
	extrTrace, ok := extracted.(Trace)
	if !ok {
		t.Errorf("unexpected type")
		return
	}
	if tr.ID() != extrTrace.ID() {
		t.Errorf("trace ID mismatch error: wanted %s ; got %s", tr.ID().String(), extrTrace.ID().String())
		return
	}
}

func TestWithSpan(t *testing.T) {
	tr := newTrace()
	s := newSpan(tr, "test")

	ctx := WithSpan(context.Background(), s)
	extracted := ctx.Value(SpanContextKey)
	if extracted == nil {
		t.Errorf("unexpected nil value")
		return
	}
	extrSpan, ok := extracted.(Span)
	if !ok {
		t.Errorf("unexpected type")
		return
	}
	if s.ID() != extrSpan.ID() {
		t.Errorf("span ID mismatch error: wanted %s ; got %s", s.ID().String(), extrSpan.ID().String())
		return
	}

}

func TestGetSpan(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		tr := newTrace()
		s := newSpan(tr, "test")

		ctx := context.WithValue(context.Background(), SpanContextKey, s)

		extrSpan := GetSpan(ctx)

		if s.ID() != extrSpan.ID() {
			t.Errorf("span ID mismatch error: wanted %s ; got %s", s.ID().String(), extrSpan.ID().String())
			return
		}
	})

	t.Run("Fail", func(t *testing.T) {
		t.Run("NoSpanValue", func(t *testing.T) {
			tr := newTrace()

			ctx := context.WithValue(context.Background(), ContextKey, tr)

			extrSpan := GetSpan(ctx)

			if extrSpan != nil {
				t.Error("expected result to be nil")
				return
			}
		})

		t.Run("NotASpan", func(t *testing.T) {
			tr := newTrace()

			ctx := context.WithValue(context.Background(), SpanContextKey, tr)

			extrSpan := GetSpan(ctx)

			if extrSpan != nil {
				t.Error("expected result to be nil")
				return
			}
		})
	})
}
