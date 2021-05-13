package main

import "github.com/sirupsen/logrus"

// Hook implement the logurs Hook interface
type Hook struct {
	id string
}

// Fire will trigger whenever you log
func (h *Hook) Fire(entry *logrus.Entry) error {
	entry.Data["id"] = h.id
	return nil
}

// Levels is what level the hook will fire on
func (h *Hook) Levels() []logrus.Level {
	return logrus.AllLevels
}

// Logrus demonstrates some basic logrus functionality
func Logrus() {

	logrus.SetFormatter(&logrus.TextFormatter{})
	logrus.SetLevel(logrus.InfoLevel)
	logrus.AddHook(&Hook{"123"})

	fields := logrus.Fields{}
	fields["success"] = true
	fields["complex_struct"] = struct {
		Event string
		When  string
	}{"Something happened", "Just now"}

	x := logrus.WithFields(fields)
	x.Warn("warning!")
	x.Error("error!")

}
