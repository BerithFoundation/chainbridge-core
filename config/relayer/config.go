// Copyright 2021 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package relayer

import (
	"fmt"

	"github.com/rs/zerolog"
)

type RelayerConfig struct {
	OpenTelemetryCollectorURL string
	LogLevel                  zerolog.Level
	LogFile                   string
	Address                   string
}

type RawRelayerConfig struct {
	OpenTelemetryCollectorURL string `mapstructure:"OpenTelemetryCollectorURL" json:"opentelemetryCollectorURL"`
	LogLevel                  string `mapstructure:"LogLevel" json:"logLevel" default:"info"`
	LogFile                   string `mapstructure:"LogFile" json:"logFile" default:"out.log"`
	Address                   string `mapstructure:"Address" json:"address" default:""`
}

func (c *RawRelayerConfig) Validate() error {
	return nil
}

// NewRelayerConfig parses RawRelayerConfig into RelayerConfig
func NewRelayerConfig(rawConfig RawRelayerConfig) (RelayerConfig, error) {
	config := RelayerConfig{}
	err := rawConfig.Validate()
	if err != nil {
		return config, err
	}

	logLevel, err := zerolog.ParseLevel(rawConfig.LogLevel)
	if err != nil {
		return config, fmt.Errorf("unknown log level: %s", rawConfig.LogLevel)
	}
	config.LogLevel = logLevel

	config.LogFile = rawConfig.LogFile
	config.OpenTelemetryCollectorURL = rawConfig.OpenTelemetryCollectorURL

	config.Address = rawConfig.Address

	return config, nil
}
