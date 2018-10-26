package echo_stats_handler_test

import (
	"fmt"
	"testing"

	"github.com/labstack/echo"
	"github.com/tkitsunai/echo-stats-handler"
)

type MockContext struct {
	echo.Context
}

func (MockContext) String(code int, s string) error {
	fmt.Printf("code: %v\n result:%v", code, s)
	return nil
}

func TestEchoStatsHandler(t *testing.T) {
	err := echo_stats_handler.Stats(&MockContext{})
	if err != nil {
		t.Fatal("Fail test", err)
	}
}
