package sisyphus

// by https://betterstack.com/community/guides/logging/logging-in-go/

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"os"

	"github.com/fatih/color"
)

// MARK: 控制台彩色日志输出

// ColorLogHandler defines a custom Handler to implement colorized output.
// 定义一个自定义 Handler 来实现彩色输出
type ColorLogHandler struct {
	handler slog.Handler
}

// NewColorLogHandler creates a new colorized log handler that outputs only to os.Stdout.
// 创建新的彩色日志处理器，输出固定到 os.Stdout
func NewColorLogHandler() *ColorLogHandler {
	return &ColorLogHandler{handler: slog.NewJSONHandler(os.Stdout, nil)}
}

// Handle implements the custom log handling logic.
// 实现自定义的日志处理逻辑
func (h *ColorLogHandler) Handle(ctx context.Context, r slog.Record) error {
	level := r.Level.String()

	// Set colors based on different log levels.
	// 根据不同的日志级别设置颜色
	switch r.Level {
	case slog.LevelDebug:
		level = color.HiMagentaString(level)
	case slog.LevelInfo:
		level = color.HiBlueString(level + " ")
	case slog.LevelWarn:
		level = color.HiYellowString(level + " ")
	case slog.LevelError:
		level = color.HiRedString(level)
	}

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
	timeStr := r.Time.Format("2006/01/02-15:04:05")
	// msg := color.CyanString(r.Message)
	msg := r.Message

	// println(timeStr, level, msg, color.WhiteString(string(b)))
	fmt.Fprintf(os.Stdout, "%s %s %s %s\n",
		timeStr, level, msg, string(b))

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

// // 设置基本的 JSON Handler 或 Text Handler
// baseHandler := slog.NewJSONHandler(os.Stdout, nil)

// // 包装为彩色日志处理器
// colorHandler := sisyphus.NewColorHandler(baseHandler)

// // 创建彩色的 slog 记录器
// logger := slog.New(colorHandler)
// // logger := slog.New(sisyphus.NewColorHandler(slog.NewJSONHandler(os.Stdout, nil)))
// logger.Debug(
// 	"executing database query",
// 	slog.String("query", "SELECT * FROM users"),
// )
// logger.Info("image upload successful", slog.String("image_id", "39ud88"))
// logger.Warn(
// 	"storage is 90% full",
// 	slog.String("available_space", "900.1 MB"),
// )
// logger.Error(
// 	"An error occurred while processing the request",
// 	slog.String("url", "https://example.com"),
// )
