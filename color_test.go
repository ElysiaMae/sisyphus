package sisyphus_test

import (
	"log/slog"
	"testing"

	"github.com/elysiamae/sisyphus"
)

func TestColor(t *testing.T) {
	cl := slog.New(sisyphus.NewColorLogHandler())

	cl.Debug("Debug Message")
	cl.Info("Info Message")
	cl.Warn("Warn Message")
	cl.Error("Error Message")
}
