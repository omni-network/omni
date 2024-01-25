package log

import (
	"log/slog"
	"strings"

	"github.com/omni-network/omni/lib/errors"

	"github.com/muesli/termenv"
	"github.com/spf13/pflag"
)

const (
	loggerConsole = "console"
	loggerJSON    = "json"

	colorForce   = "force"
	colorDisable = "disable"
	colorAuto    = "auto"
)

//nolint:gochecknoglobals // Static mapping.
var loggerFuncs = map[string]func(...func(*options)) *slog.Logger{
	loggerConsole: newConsoleLogger,
	loggerJSON:    newJSONLogger,
}

//nolint:gochecknoglobals // Static mapping.
var colors = map[string]termenv.Profile{
	colorForce:   termenv.TrueColor,
	colorDisable: termenv.Ascii,
	colorAuto:    termenv.ColorProfile(),
}

// DefaultConfig returns a default config.
func DefaultConfig() Config {
	return Config{
		Level:  strings.ToLower(slog.LevelInfo.String()),
		Color:  colorAuto,
		Format: loggerConsole,
	}
}

type Config struct {
	Level  string
	Color  string
	Format string
}

func (c Config) color() (termenv.Profile, error) {
	color := c.Color
	if c.Color == "" {
		color = colorAuto
	}
	resp, ok := colors[color]
	if !ok {
		return 0, errors.New("invalid color", "color", c.Color)
	}

	return resp, nil
}

func (c Config) level() (slog.Level, error) {
	var level slog.Level
	if err := level.UnmarshalText([]byte(c.Level)); err != nil {
		return slog.Level(0), errors.Wrap(err, "parse log level")
	}

	return level, nil
}

func (c Config) loggerFunc() (func(...func(*options)) *slog.Logger, error) {
	f, ok := loggerFuncs[c.Format]
	if !ok {
		return nil, errors.New("invalid format", "format", c.Format)
	}

	return f, nil
}

// make returns a new global as per the config.
func (c Config) make() (*slog.Logger, error) {
	level, err := c.level()
	if err != nil {
		return nil, errors.Wrap(err, "parse log level")
	}

	color, err := c.color()
	if err != nil {
		return nil, errors.Wrap(err, "parse log color")
	}

	loggerFunc, err := c.loggerFunc()
	if err != nil {
		return nil, errors.Wrap(err, "parse log format")
	}

	return loggerFunc(func(o *options) {
		o.Level = level
		o.Color = color
	}), nil
}

// BindFlags binds the standard flags to provide logging config at runtime.
//
//nolint:lll,revive // Long lines are actually more readable here.
func BindFlags(flags *pflag.FlagSet, cfg *Config) {
	flags.StringVar(&cfg.Level, "log-level", cfg.Level, "Log level; debug, info, warn, error")
	flags.StringVar(&cfg.Color, "log-color", cfg.Color, "Log color (only applicable to console format); auto, force, disable")
	flags.StringVar(&cfg.Format, "log-format", cfg.Format, "Log format; console, json")
}
