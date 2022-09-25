package slog

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"time"
)

const (
	TYPE_TERMINAL = "terminal"
	TYPE_FILE     = "file"
)

var logger *logrus.Entry

func Init(level logrus.Level) {
	_, err := logrus.ParseLevel(level.String())
	if err != nil {
		panic(err.Error())
	}

	if viper.GetString("log.output") == TYPE_FILE {
		filePath := "./log/gin-kit.log"
		//fileName := viper.GetString("log.name")
		if viper.GetBool("log.fileRotate.isOpen") {
			maxAge := viper.GetInt64("log.maxAge")             //unit: day
			rotationTime := viper.GetInt64("log.rotationTime") //unit: day
			Day := 24 * time.Hour

			writer, err := rotatelogs.New(filePath+".%Y_%m_%d_%H_%M",
				rotatelogs.WithLinkName(filePath),                            //create soft link which pointer to latest log file
				rotatelogs.WithMaxAge(time.Duration(maxAge)*Day),             //max age
				rotatelogs.WithRotationTime(time.Duration(rotationTime)*Day), //the time between rotation

				//for test
				//rotatelogs.WithMaxAge(time.Minute*3),  //max age
				//rotatelogs.WithRotationTime(time.Minute), //the time between rotation
			)
			if err != nil {
				panic("create log writer failed, err:" + err.Error())
			}
			writerMap := lfshook.WriterMap{
				logrus.InfoLevel:  writer,
				logrus.FatalLevel: writer,
				logrus.DebugLevel: writer,
				logrus.WarnLevel:  writer,
				logrus.ErrorLevel: writer,
				logrus.PanicLevel: writer,
			}
			rotateFileHook := lfshook.NewHook(writerMap, &logrus.JSONFormatter{PrettyPrint: true})
			logrus.AddHook(rotateFileHook)
			logrus.SetOutput(writer)
			logrus.SetReportCaller(true)
		} else {
			file, err := os.OpenFile(filePath, os.O_RDWR|os.O_APPEND|os.O_CREATE, os.ModePerm)
			if err != nil {
				panic("open log file failed, err:" + err.Error())
			}
			logrus.SetOutput(file)
		}
		logrus.SetFormatter(&logrus.JSONFormatter{PrettyPrint: true})
	} else {
		logrus.SetOutput(os.Stdout)
		logrus.SetFormatter(&logrus.TextFormatter{ForceColors: true})
	}

	/*
		//add notify hook
		if viper.GetBool("log.notify.isOpen") {
			notifyHook := NewNotifyHook()
			logrus.AddHook(notifyHook)
		}

	*/

	logrus.SetLevel(level)
	logger = logrus.WithFields(logrus.Fields{
		"appName": viper.GetString("appName"),
		"author":  viper.GetString("author"),
		"ver":     viper.GetString("version"),
	})
}

func SetLevel(l logrus.Level) {
	logrus.SetLevel(l)
}

func GetLogger() *logrus.Entry {
	return logger
}
