package auth

import (
	"context"
	"log/slog"
)

// CasbinSlogLogger is the implementation for a Logger using golang log.
type CasbinSlogLogger struct {
	enabled bool
	slogger *slog.Logger
	ctx     context.Context
	level   slog.Level
}

func NewCasbinSlogLogger(slogger *slog.Logger) *CasbinSlogLogger {
	return &CasbinSlogLogger{
		enabled: true,
		slogger: slogger,
		ctx:     context.Background(),
	}
}

func (l *CasbinSlogLogger) EnableLog(enable bool) {
	l.enabled = enable
}

func (l *CasbinSlogLogger) IsEnabled() bool {
	return l.enabled
}

func (l *CasbinSlogLogger) LogModel(model [][]string) {
	if !l.enabled {
		return
	}

	l.slogger.LogAttrs(l.ctx, level, "Casbin Model", slog.Any("Model", model))
}

func (l *CasbinSlogLogger) LogEnforce(matcher string, request []interface{}, result bool, explains [][]string) {
	if !l.enabled {
		return
	}

	l.slogger.LogAttrs(l.ctx, level, "Casbin Enforce", slog.Group("Enforce", slog.Any("Request", request), slog.Bool("Result", result), slog.Any("Explain", explains)))
}

func (l *CasbinSlogLogger) LogPolicy(policy map[string][][]string) {
	if !l.enabled {
		return
	}

	l.slogger.LogAttrs(l.ctx, level, "Casbin Policy", slog.Any("Policy", policy))
}

func (l *CasbinSlogLogger) LogRole(roles []string) {
	if !l.enabled {
		return
	}

	l.slogger.LogAttrs(l.ctx, level, "Casbin Role", slog.Any("Roles", roles))
}

// LogError implements log.Logger.
func (l *CasbinSlogLogger) LogError(err error, msg ...string) {
	if !l.enabled {
		return
	}

	l.slogger.LogAttrs(l.ctx, l.level, "Casbin Error", slog.Any("Error", err), slog.Any("Message", msg))
}
