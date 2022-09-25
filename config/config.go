package config

import (
	"flag"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"os"
	"runtime"
	"syscall"
)

var stdFileHandler *os.File

func Init() {
	path := flag.String("config", "", "the config path of this application")
	flag.Parse()

	fmt.Println(*path)
	viper.SetConfigFile(*path)
	viper.SetConfigType("yaml")
	viper.OnConfigChange(func(in fsnotify.Event) {
		log.Println("config file has been changed")
	})
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

func InitTest() {
	path := "/Users/chuckchen/Study/BackEnd/Project/sparrow/config/dev.yaml"
	fmt.Println(path)
	viper.SetConfigFile(path)
	viper.SetConfigType("yaml")
	viper.OnConfigChange(func(in fsnotify.Event) {
		log.Println("config file has been changed")
	})
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

func RewriteStdErrFile(path string) error {
	if runtime.GOOS == "windows" {
		return nil
	}
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Println(err)
		return err
	}
	stdFileHandler = f //avoid memory gc

	err = syscall.Dup2(int(f.Fd()), int(os.Stderr.Fd()))
	if err != nil {
		log.Println(err)
		return err
	}

	//close fd before stdFileHandler be gc
	runtime.SetFinalizer(stdFileHandler, func(fd *os.File) {
		fd.Close()
	})

	return nil
}
