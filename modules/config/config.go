package config

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Discord struct {
		Token string `yaml:"token"`
	} `yaml:"discord"`

	Logging struct {
		Enabled bool   `yaml:"enabled"`
		Level   string `yaml:"level"`
		Path    string `yaml:"path"`
	} `yaml:"logging"`

	Music struct {
		MaxQueueSize    int      `yaml:"max_queue_size"`
		DefaultVolume   int      `yaml:"default_volume"`
		AllowedChannels []string `yaml:"allowed_channels"`
	} `yaml:"music"`
}

const defaultConfig = `# Discord Bot Configuration
discord:
  token: "your-token-here" # Replace with your bot token

logging:
  enabled: true
  level: "info"  # debug, info, warn, error
  path: "data/logs"

# TODO: Maybe figure this out LOL
music:
  max_queue_size: 100
  default_volume: 70
  allowed_channels:
    - "music"
    - "dj-booth"
`

// Load reads the config file from disk or creates it with defaults if it doesn't exist
func Load() (*Config, error) {
	// Get the executable's directory
	exe, err := os.Executable()
	if err != nil {
		return nil, fmt.Errorf("failed to get executable path: %w", err)
	}
	exeDir := filepath.Dir(exe)

	// Ensure data directory exists
	dataDir := filepath.Join(exeDir, "data")
	if err := os.MkdirAll(dataDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create data directory: %w", err)
	}

	configPath := filepath.Join(dataDir, "config.yml")

	// Check if config file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		// Create default config file
		if err := os.WriteFile(configPath, []byte(defaultConfig), 0644); err != nil {
			return nil, fmt.Errorf("failed to create default config file: %w", err)
		}
		fmt.Printf("Created default configuration file at: %s\n", configPath)
	}

	// Read config file
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	// Parse YAML
	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed to parse config file: %w", err)
	}

	// Validate required fields
	if err := config.validate(); err != nil {
		return nil, fmt.Errorf("invalid configuration: %w", err)
	}

	return &config, nil
}

// validate checks if all required fields are properly set
func (c *Config) validate() error {
	if c.Discord.Token == "" || c.Discord.Token == "your-token-here" {
		return fmt.Errorf("discord token is required")
	}

	//if c.Music.MaxQueueSize <= 0 {
	//	return fmt.Errorf("max queue size must be positive")
	//}

	//if c.Music.DefaultVolume < 0 || c.Music.DefaultVolume > 100 {
	//	return fmt.Errorf("default volume must be between 0 and 100")
	//}

	return nil
}

// Save writes the current configuration back to the file
func (c *Config) Save() error {
	exe, err := os.Executable()
	if err != nil {
		return fmt.Errorf("failed to get executable path: %w", err)
	}

	configPath := filepath.Join(filepath.Dir(exe), "data", "config.yml")

	data, err := yaml.Marshal(c)
	if err != nil {
		return fmt.Errorf("failed to marshal config: %w", err)
	}

	if err := os.WriteFile(configPath, data, 0644); err != nil {
		return fmt.Errorf("failed to write config file: %w", err)
	}

	return nil
}
