package sisyphus

// by https://betterstack.com/community/guides/logging/logging-in-go/
// by https://dusted.codes/creating-a-pretty-console-logger-using-gos-slog-package

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"strings"
)

// 256 Color
// This will be compatible with macOS Terminal.app and iTerm2.app.
// Currently, iTerm2.app supports up to xterm-256color.
const (
	reset = "\033[0m"

	fgBlack   = "232"
	fgRed     = "9"
	fgGreen   = "10"
	fgYellow  = "11"
	fgBlue    = "12"
	fgMagenta = "13"
	fgCyan    = "14"
	fgWhite   = "255"

	bgBlack   = "0"
	bgRed     = "1"
	bgGreen   = "2"
	bgYellow  = "3"
	bgBlue    = "4"
	bgMagenta = "5"
	bgCyan    = "6"
	bgWhite   = "7"
)

// colorizeBg 支持 256 位真彩色背景和前景
func colorizeBg(v string, fg string, bg string, bold bool) string {
	if bold {
		return fmt.Sprintf("\033[1;38;5;%sm\033[48;5;%sm%s%s", fg, bg, v, reset)
	}
	return fmt.Sprintf("\033[38;5;%sm\033[48;5;%sm%s%s", fg, bg, v, reset)
}

// colorizeBg 支持 256 位真彩色前景
func colorize(v string, fg string, bold bool) string {
	if bold {
		return fmt.Sprintf("\033[1;38;5;%sm%s%s", fg, v, reset)
	}
	return fmt.Sprintf("\033[38;5;%sm%s%s", fg, v, reset)
}

// MARK: 控制台彩色日志输出

// ColorLogHandler defines a custom Handler to implement colorized output.
// 定义一个自定义 Handler 来实现彩色输出
type ColorLogHandler struct {
	handler slog.Handler
}

// NewColorLogHandler creates a new colorized log handler that outputs only to os.Stdout.
// 创建新的彩色日志处理器，输出固定到 os.Stdout
func NewColorLogHandler(opts *slog.HandlerOptions) *ColorLogHandler {
	if opts == nil {
		return &ColorLogHandler{handler: slog.NewJSONHandler(os.Stdout, nil)}
	} else {
		return &ColorLogHandler{handler: slog.NewJSONHandler(os.Stdout, opts)}
	}
}

// Handle implements the custom log handling logic.
// 实现自定义的日志处理逻辑
func (h *ColorLogHandler) Handle(ctx context.Context, r slog.Record) error {
	var buildr strings.Builder
	buildr.WriteString(" ")
	buildr.WriteString(r.Level.String())
	buildr.WriteString(" ")
	if r.Level == slog.LevelInfo || r.Level == slog.LevelWarn {
		buildr.WriteString(" ")
	}
	level := buildr.String()

	// Process the attributes in the log and convert them to strings.
	// 处理日志中的属性并转换为字符串
	fields := make(map[string]string, r.NumAttrs())
	r.Attrs(func(a slog.Attr) bool {
		fields[a.Key] = a.Value.String()
		return true
	})

	var b []byte
	var err error
	if len(fields) != 0 {
		// json. | json.MarshalIndent
		// b, err := json.MarshalIndent(fields, "", "  ")
		b, err = json.Marshal(fields)
		if err != nil {
			return err
		}
	}

	// Format the time and message.
	// 格式化时间和消息
	timeStr := r.Time.Format("2006/01/02-15:04:05.000")
	msg := r.Message

	// Set colors based on different log levels.
	// 根据不同的日志级别设置颜色
	switch r.Level {
	case slog.LevelDebug:
		level = colorizeBg(level, fgBlack, bgMagenta, true)
		msg = colorize(msg, fgMagenta, false)
	case slog.LevelInfo:
		level = colorizeBg(level, fgBlack, bgBlue, true)
		// msg = colorize(msg, fgBlue, false)
	case slog.LevelWarn:
		level = colorizeBg(level, fgBlack, bgYellow, true)
		msg = colorize(msg, fgYellow, false)
	case slog.LevelError:
		level = colorizeBg(level, fgBlack, bgRed, true)
		msg = colorize(msg, fgRed, false)
	}

	_, err = fmt.Fprintf(os.Stdout, "%s |%s| %s %s\n",
		timeStr, level, msg, string(b))
	if err != nil {
		return err
	}

	return nil
}

// Enabled checks if the given log level is enabled.
// 方法实现检查给定级别是否启用
func (h *ColorLogHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return h.handler.Enabled(ctx, level)
}

// WithAttrs returns a new Handler with additional attributes.
// 方法实现返回附加属性的新 Handler
func (h *ColorLogHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	// Return a new ColorHandler instance with extra attributes.
	// 返回一个新的 ColorHandler 实例，附加额外的属性
	return &ColorLogHandler{
		handler: h.handler.WithAttrs(attrs),
	}
}

// WithGroup returns a new Handler with a log group name.
// WithGroup 方法实现返回带有日志组名称的新 Handler
func (h *ColorLogHandler) WithGroup(name string) slog.Handler {
	// Return a new ColorHandler instance with a log group name.
	// 返回一个新的 ColorHandler 实例，附加日志组名称
	return &ColorLogHandler{
		handler: h.handler.WithGroup(name),
	}
}
