package job

// NewConf returns Config with default values
func NewConf() *Config {
	return &Config{
		UpdateEvery:        1,
		AutoDetectionRetry: 0,
		ChartCleanup:       10,
		RetriesMax:         60,
	}
}

type Config struct {
	moduleName         string // standalone struct ?
	jobName            string // standalone struct ?
	OverrideName       string `yaml:"name"`
	UpdateEvery        int    `yaml:"update_every"`
	AutoDetectionRetry int    `yaml:"autodetection_retry"`
	ChartCleanup       int    `yaml:"chart_cleanup"`
	RetriesMax         int    `yaml:"retries"`
}

// TODO: GetModuleName() prepends "go_"
func (c *Config) GetModuleName() string {
	return "go_" + c.moduleName
}

func (c *Config) GetFullName() string {
	if c.jobName == "" {
		return c.GetModuleName()
	}
	return c.GetModuleName() + "_" + c.GetJobName()
}

func (c *Config) GetJobName() string {
	if c.jobName == "" {
		return c.GetModuleName()
	}
	if c.OverrideName == "" {
		return c.jobName
	}
	return c.OverrideName
}

func (c *Config) GetUpdateEvery() int {
	return c.UpdateEvery
}

func (c *Config) SetModuleName(name string) {
	c.moduleName = name
}

func (c *Config) SetJobName(name string) {
	c.jobName = name
}

func (c *Config) SetUpdateEvery(u int) {
	c.UpdateEvery = u
}
