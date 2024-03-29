package config

type ShowStart struct {
	Email                EmailConfig `yaml:"email"`
	City                 []string    `yaml:"city"`
	TagsSelected         []string    `yaml:"tags_selected"`
	InitialEventID       int64       `yaml:"initial_event_id,omitempty"`
	DBFile               string      `yaml:"db_file"`
	Log                  Log         `yaml:"log"`
	MaxNotFoundCount     int64       `yaml:"maxNotFoundCount"`
	Max404CountToCheck   int64       `yaml:"max404CountToCheck"`
	OtherCityInAfternoon []string    `yaml:"otherCityInAfternoon"`
}

type Simullink struct {
	Email        EmailConfig `yaml:"email"`
	CityCode     string      `yaml:"city_code"`
	URL          string      `yaml:"url"`
	TagsSelected []string    `yaml:"tags_selected,omitempty"`
	DBDir        string      `yaml:"db_dir"`
	Log          Log         `yaml:"log"`
}

type Zhengzai struct {
	Email  EmailConfig `yaml:"email"`
	AdCode string      `yaml:"ad_code"`
	URL    string      `yaml:"url"`
	DBDir  string      `yaml:"db_dir"`
	Log    `yaml:"log"`
}

type EmailConfig struct {
	From     string `yaml:"from"`
	Password string `yaml:"password"`
	Server   string `yaml:"server"`
	Port     int    `yaml:"port"`
	To       string `yaml:"to"`
}

type Log struct {
	LogSuffix string `yaml:"log_suffix"`
	LogDir    string `yaml:"log_dir"`
}
