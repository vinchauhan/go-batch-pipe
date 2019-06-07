package config

//ConfigInfo is a type to return a config for each type of source like file, kafka, rabbit
type ConfigInfo struct {
	config map[string] interface{}
}

func NewConfig(s string) *ConfigInfo {
	configMap := make(map[string]interface{})
	configMap[s] = nil
	return &ConfigInfo{
		config: configMap,
	}
}

//Initialize the reader config info
func (c * ConfigInfo) BuildSpecifReaderConfigInfo() error {

}