package slog

/*
//notify hook
type NotifyHook struct {

}

func NewNotifyHook() *NotifyHook  {
	return &NotifyHook{}
}

func (*NotifyHook)Levels() []logrus.Level {
	return []logrus.Level{
		logrus.ErrorLevel,
		logrus.FatalLevel,
		logrus.PanicLevel,
	}
}

func (*NotifyHook)Fire(entry *logrus.Entry) error {
	if !viper.GetBool("log.notify.isOpen") {
		return nil
	}

	notifyObj, err  :=  notify.NewNotifyObj(&notify.NotifyConfig{
		To: viper.GetString("log.notify.to"),
		TopicIDs: []int{1181},
	})
	if err != nil {
		return err
	}
	notifyObj.Send(&model.NtfMsg{
		Summary: viper.GetString("appName") + " error happened, please relevant personnel deal it as soon as possible ",
		ContentType: model.CONTENT_TYPE_TEXT,
		Content: model.LogContent{
			Level: strconv.Itoa(int(entry.Level)),
			Time: time.Now().Format(time.RFC3339),
			Body: entry.Message,
		}.String(),
	})

	return nil
}

*/
