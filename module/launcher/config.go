package launcher

import (
	"gopkg.in/ini.v1"
)

type config struct {
	path string
}

type ExceptionHandler[T any] struct {
	Code    int    `json:"code"`
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

func NewExceptionHandler[T any](code int, status bool, message string, data T) ExceptionHandler[T] {
	return ExceptionHandler[T]{
		Code:    code,
		Status:  status,
		Message: message,
		Data:    data,
	}
}

type ExportToFrontend struct {
	Data string `json:"data"`
}

func NewConfig(path string) *config {
	return &config{path: path}
}

func (config config) Read(section, key string) (string, error) {
	cfg, err := ini.Load(config.path)
	if err != nil {
		return "", err
	}
	sec, err := cfg.GetSection(section)
	if err != nil {
		return "", err
	}
	k, err := sec.GetKey(key)
	if err != nil {
		return "", err
	}
	return k.String(), nil
}
func (config config) Write(section, key, value string) error {
	cfg, err := ini.Load(config.path)
	if err != nil {
		cfg = ini.Empty()
	}
	s, err := cfg.GetSection(section)
	if err != nil {
		s, _ = cfg.NewSection(section)
	}
	s.Key(key).SetValue(value)
	return cfg.SaveTo(config.path)
}
