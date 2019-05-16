package grace

import (
	"context"
	"syscall"
	"testing"
)

func TestAll(t *testing.T) {
	Init(syscall.SIGTERM)
	Go(func(ctx context.Context) {
		t.Log("backend goroutine running")
	})
	Go(func(ctx context.Context) {
		t.Log("backend goroutine running")
		Go(func(ctx context.Context) {
			t.Log("dynamic goroutine running")
		})
	})
	Run()
}
