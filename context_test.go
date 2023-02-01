package spanner

import (
	"context"
	"testing"
)

func TestContextData(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		trID := NewTraceID()
		sID := NewSpanID()
		cd := &ContextData{
			Trace: &trID,
			Span:  &sID,
		}
		ctx := WithContextData(context.Background(), cd)
		if ctx == nil {
			t.Errorf("unexpected nil value")
			return
		}

		// ensure Value() is covered, too
		v := ctx.Value(ContextKey)
		if v == nil {
			t.Errorf("unexpected nil value")
			return
		}
		if _, ok := v.(ContextData); !ok {
			t.Errorf("type mismatch error: wanted %T ; got %T", &ContextData{}, v)
			return
		}

		storedCD := GetContextData(ctx)
		if storedCD == nil {
			t.Errorf("unexpected nil value")
			return
		}

		if storedCD.Trace == nil || storedCD.Span == nil {
			t.Errorf("unexpected nil value")
			return
		}

		if *storedCD.Trace != trID || *storedCD.Span != sID {
			t.Errorf("content mismatch error")
			return
		}
	})

	t.Run("NilCD", func(t *testing.T) {
		ctx := WithContextData(context.Background(), nil)
		if ctx == nil {
			t.Errorf("unexpected nil value")
			return
		}

		storedCD := GetContextData(ctx)
		if storedCD != nil {
			t.Errorf("unexpected non-nil value")
			return
		}
	})
}
