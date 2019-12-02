package pkg

type Job struct {
	StartTime string `yaml:"start_time"`
	StopTime  string `yaml:"stop_time"`

	Key        string `yaml:"key"`
	RegionID   int    `yaml:"region_id"`
	PlanID     int    `yaml:"plan_id"`
	OsID       int    `yaml:"os_id"`
	ServerName string `yaml:"server_name"`
}
