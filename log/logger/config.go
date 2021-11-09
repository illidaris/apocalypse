package logger

// Config cfg logger
type Config struct {
	StdLevel      string `toml:"std_level" json:"std_level"`     // Std Log level.
	StdFormat     string `toml:"std_format" json:"std_format"`   // Std Log format. one of json, text, or console.
	Level         string `toml:"level" json:"level"`             // Log level.
	Format        string `toml:"format" json:"format"`           // Log format. one of json, text, or console.
	FileDirectory string `toml:"file_dir" json:"file_dir"`       // File directory
	FileName      string `toml:"file_name" json:"file_name"`     // Log filename, leave empty to disable file log.
	MaxSize       int    `toml:"max_size" json:"max_size"`       // Max size for a single file, in MB.
	MaxDays       int    `toml:"max_days" json:"max_days"`       // Max log keep days, default is never deleting.
	MaxBackups    int    `toml:"max_backups" json:"max_backups"` // Maximum number of old log files to retain.
	Compress      bool   `toml:"compress" json:"compress"`       // Compress
}

// GetStdLevel Get Std Level
func (c *Config) GetStdLevel() *Level {
	l := new(Level)
	l.Unpack(c.StdLevel)
	return l
}

// GetLevel Get File Level
func (c *Config) GetLevel() *Level {
	l := new(Level)
	l.Unpack(c.Level)
	return l
}

func defaultConfig() *Config {
	return &Config{
		StdLevel:      "error",
		StdFormat:     "console",
		Level:         "debug",
		Format:        "json",
		FileDirectory: "./logs",
		FileName:      "log",
		MaxSize:       10,
		MaxDays:       7,
		MaxBackups:    30,
		Compress:      false,
	}
}
