package viperconf

import (
	"fmt"
	"github.com/spf13/viper"
)

func ViperExample(path string, name string) *viper.Viper {
	v := viper.New()
	v.AddConfigPath(path)
	v.SetConfigName(name)
	if err := v.ReadInConfig(); err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return v
}