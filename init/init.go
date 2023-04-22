package init

import "sync"
import "github.com/spf13/viper"

func initTWSettings() {
	once := &sync.Once{}
	once.Do(func() {
		viper.SetConfigName("base")
		viper.SetConfigType("yaml")
		viper.AddConfigPath("../configs")
		if err := viper.ReadInConfig(); err != nil {
			panic("Fail to init timing wheel config!")
		}
	})
}

func init() {
	initTWSettings()
}
