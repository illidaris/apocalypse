package log

type Base struct {
	Level  string `toml:"level" json:"level"`   // Log level.
	Format string `toml:"format" json:"format"` // Log format. one of json, text, or console.
}

// GetLevel
/**
 * @Description:
 * @receiver c
 * @return *Level
 */
func (c *Base) GetLevel() *Level {
	l := new(Level)
	l.Unpack(c.Level)
	return l
}

type Std struct {
	Base
}

func (cfg *Std) NewExporter() IExporter {
	return &StdExporter{
		Core: cfg,
	}
}

type File struct {
	Base
	FileDirectory string `toml:"fileDir" json:"fileDir"`       // File directory
	FileName      string `toml:"fileName" json:"fileName"`     // Log filename, leave empty to disable file log.
	MaxSize       int    `toml:"maxSize" json:"maxSize"`       // Max size for a single file, in MB.
	MaxDays       int    `toml:"maxDays" json:"maxDays"`       // Max log keep days, default is never deleting.
	MaxBackups    int    `toml:"maxBackups" json:"maxBackups"` // Maximum number of old log files to retain.
	Compress      bool   `toml:"compress" json:"compress"`     // Compress
}

func (cfg *File) NewExporter() IExporter {
	return &FileExporter{
		Core: cfg,
	}
}

// Config
/**
 * @Description: serializes log related config in toml/json.
 */
type Config struct {
	CallSkip int   `toml:"callSkip" json:"callSkip"` // Log CallSkip
	FileCfg  *File `toml:"file" json:"file"`         // FileCfg log config.
	StdCfg   *Std  `toml:"std" json:"std"`           // StdCfg log config.
}
