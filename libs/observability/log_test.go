package observability

import (
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"testing"
)

func TestLoggable(t *testing.T) {
	t.Run("successful execution", func(t *testing.T) {
		result, _ := Log("successful function", func(context *LogContext) (string, error) {
			context.Add(zap.Int("user_id", 123))

			return "hello world", nil
		})

		assert.Equal(t, result, "hello world")
	})
}
