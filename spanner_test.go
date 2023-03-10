package spanner_test

import (
	"bytes"
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/zalgonoise/attr"
	"github.com/zalgonoise/logx"
	"github.com/zalgonoise/spanner"
)

func runtime() {
	ctx, startS := spanner.Start(context.Background(), `Runtime:"Main"`)
	defer startS.End()

	_, s := spanner.Start(ctx, "Runtime:Start:A")
	x := runtimeA(ctx, 2)
	s.End()
	_, s = spanner.Start(ctx, "Runtime:Start:E")
	runtimeE(ctx, "Hello", x)
	s.End()
}

func runtimeA(ctx context.Context, i int) int {
	ctx, s := spanner.Start(ctx, "Runtime:A")
	defer s.End()

	s.Event("A: multiply by 2")
	x := i * 2

	return runtimeB(ctx, x)
}

func runtimeB(ctx context.Context, i int) int {
	ctx, s := spanner.Start(ctx, "Runtime:B")
	defer s.End()

	s.Event("B: multiply by 2")
	x := i * 2

	return runtimeC(ctx, x)
}

func runtimeC(ctx context.Context, i int) int {
	ctx, s := spanner.Start(ctx, "Runtime:C")
	defer s.End()

	s.Event("C: multiply by 2")
	x := i * 2

	return x
}

func runtimeD(ctx context.Context, text string) {
	ctx, s := spanner.Start(ctx, "Runtime:D")
	defer s.End()

	fmt.Println(text)
}
func runtimeE(ctx context.Context, text string, i int) {
	ctx, s := spanner.Start(ctx, "Runtime:E")
	defer s.End()

	s.Add(attr.String("text", text), attr.Int("result", i))
	runtimeD(ctx, fmt.Sprintf("%s ; result: %v", text, i))
}
func TestFunctionsWithSpan(t *testing.T) {
	buf := new(bytes.Buffer)
	spanner.To(spanner.Writer(buf))

	runtime()

	spanner.Processor().Flush(context.Background())
	time.Sleep(10 * time.Millisecond)
	t.Log(buf.String())

	t.Error()
}

func TestMainSpan(t *testing.T) {
	buf := new(bytes.Buffer)
	spanner.To(spanner.Writer(buf))
	ctx := logx.InContext(context.Background(), logx.Default())
	ctx, s := spanner.Start(ctx, "main")
	s.End()
	spanner.Processor().Flush(context.Background())
	time.Sleep(10 * time.Millisecond)
	t.Log(buf.String())

	t.Error()
}
