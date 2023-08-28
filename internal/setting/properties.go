package setting

type ServerProperties struct {
	RunMode  string
	HttpPort string
}

type AppProperties struct {
	LoggerFileName   string
	LoggerLevel      string
	LoggerMaxSize    int
	LoggerMaxBackups int
	LoggerMaxAge     int
	UploadSavePath   string
	UploadServerUrl  string
	UploadMaxSize    int64
	UploadAllowExt   []interface{}
}
