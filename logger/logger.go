package logger

import (
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"path/filepath"
)

const (
	Sep = string(filepath.Separator)
)

func Init(baseDir string) (*logrus.Logger, error) {
	logFile, err := os.OpenFile(baseDir+Sep+"log"+Sep+"log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o666)
	if err != nil {
		return nil, err
	}
	mw := io.MultiWriter(os.Stdout, logFile)

	log := logrus.New()
	log.SetOutput(mw)

	return log, nil
}
