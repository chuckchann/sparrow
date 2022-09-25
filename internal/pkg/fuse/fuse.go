package fuse

import (
	"github.com/afex/hystrix-go/hystrix"
	"github.com/spf13/viper"
)

func Init() {
	hystrix.ConfigureCommand(viper.GetString("appName"), hystrix.CommandConfig{
		Timeout:                3000, // how long to wait for command to complete
		MaxConcurrentRequests:  3,    // how many commands of the same type can run at the same time
		SleepWindow:            3000, // sleep window
		ErrorPercentThreshold:  10,   // error percent threshold
		RequestVolumeThreshold: 5,    // request threshold
	})
}
