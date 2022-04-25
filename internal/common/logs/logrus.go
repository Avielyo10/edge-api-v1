package logs

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/redhatinsights/edge-api/config"
	lc "github.com/redhatinsights/platform-go-middlewares/logging/cloudwatch"
	"github.com/sirupsen/logrus"
)

// Log is an instance of the global logrus.Logger
var logLevel logrus.Level

// hook is an instance of cloudwatch.hook
var hook *lc.Hook

// InitLogger initializes the API logger
func Init() {
	cfg := config.Get()
	switch cfg.LogLevel {
	case "DEBUG":
		logLevel = logrus.DebugLevel
	case "ERROR":
		logLevel = logrus.ErrorLevel
	default:
		logLevel = logrus.InfoLevel
	}
	logrus.SetReportCaller(true)

	if cfg.Logging != nil && cfg.Logging.Region != "" {
		cred := credentials.NewStaticCredentials(cfg.Logging.AccessKeyID, cfg.Logging.SecretAccessKey, "")
		awsconf := aws.NewConfig().WithRegion(cfg.Logging.Region).WithCredentials(cred)
		hook, err := lc.NewBatchingHook(cfg.Logging.LogGroup, cfg.Hostname, awsconf, 10*time.Second)
		if err != nil {
			logrus.WithFields(logrus.Fields{"error": err.Error()}).Error("Error creating AWS hook")
		}
		logrus.AddHook(hook)
		logrus.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: time.RFC3339Nano,
			FieldMap: logrus.FieldMap{
				logrus.FieldKeyTime:  "time",
				logrus.FieldKeyLevel: "severity",
				logrus.FieldKeyMsg:   "message",
			},
			CallerPrettyfier: func(f *runtime.Frame) (string, string) {
				s := strings.Split(f.Function, ".")
				funcName := s[len(s)-1]
				return funcName, fmt.Sprintf("%s:%d", path.Base(f.File), f.Line)
			},
		})
	}
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logLevel)
}

// FlushLogger Flush batched logging messages
func FlushLogger() {
	if hook != nil {
		err := hook.Flush()
		if err != nil {
			logrus.WithFields(logrus.Fields{"error": err.Error()}).Error("Error flushing batched logging messages")
		}
	}
}

// LogErrorAndPanic Records the error, flushes the buffer, then panics the container
func LogErrorAndPanic(msg string, err error) {
	logrus.WithFields(logrus.Fields{"error": err}).Error(msg)
	FlushLogger()
	panic(err)
}
