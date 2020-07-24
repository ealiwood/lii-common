package setting

import "github.com/spf13/viper"

type Setting struct {
	vp *viper.Viper
}

func New(configName, configPath, configType string) (*Setting, error) {
	vp := viper.New()
	vp.SetConfigName(configName)
	vp.AddConfigPath(configPath)
	vp.SetConfigType(configType)
	if err := vp.ReadInConfig(); err != nil {
		return nil, err
	}
	return &Setting{vp}, nil
}

func (s *Setting) UnmarshalAll(v interface{}) error {
	return s.vp.Unmarshal(v)
}
