package syslog

import (
	"log/syslog"
	"testing"

	"github.com/dorofeevsa/logrus"
)

func TestLocalhostAddAndPrint(t *testing.T) {
	log := logrus.New()
	hook, err := NewHook("udp", "localhost:514", syslog.LOG_INFO, "")

	if err != nil {
		t.Errorf("Unable to connect to local syslog.")
	}

	log.Hooks.Add(hook)

	for _, level := range hook.Levels() {
		if len(log.Hooks[level]) != 1 {
			t.Errorf("SyslogHook was not added. The length of log.Hooks[%v]: %v", level, len(log.Hooks[level]))
		}
	}

	log.Info("Congratulations!")
}
