package config

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setEnvs(t *testing.T, m map[string]string) {
	t.Helper()
	for k, v := range m {
		t.Setenv(k, v)
	}
}

func TestLoad(t *testing.T) {
	tests := []struct {
		name      string
		value     map[string]string
		want      *Config
		assertion assert.ErrorAssertionFunc
	}{
		{
			name: "success",
			value: map[string]string{
				"ENV":       "local",
				"PORT":      "8070",
				"API_TOKEN": "test_token",
			},
			want: &Config{
				Env:      "local",
				Port:     "8070",
				APIToken: "test_token",
			},
			assertion: assert.NoError,
		},
		{
			name: "failed. ENV is not local, dev, prod",
			value: map[string]string{
				"ENV":       "invalid_env",
				"PORT":      "8070",
				"API_TOKEN": "test_token",
			},
			want:      nil,
			assertion: assert.Error,
		},
		{
			name: "failed. PORT is not a number",
			value: map[string]string{
				"ENV":       "local",
				"PORT":      "invalid_port",
				"API_TOKEN": "test_token",
			},
			want:      nil,
			assertion: assert.Error,
		},
		{
			name: "failed. API_TOKEN is not set",
			value: map[string]string{
				"ENV":  "local",
				"PORT": "8070",
			},
			want:      nil,
			assertion: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setEnvs(t, tt.value)
			got, err := Load(context.TODO())
			tt.assertion(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestGet(t *testing.T) {
	tests := []struct {
		name      string
		setup     func()
		want      *Config
		assertion assert.ErrorAssertionFunc
	}{
		{
			name: "success",
			setup: func() {
				c = &Config{}
			},
			want:      &Config{},
			assertion: assert.NoError,
		},
		{
			name:      "failed. config is not loaded",
			setup:     func() {},
			want:      nil,
			assertion: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				tt.setup()
			}
			got, err := Get()
			tt.assertion(t, err)
			assert.Equal(t, tt.want, got)

			// cleanup
			c = nil
		})
	}
}
