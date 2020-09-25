package logger

import (
	"database/sql/driver"
	"fmt"
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gopkg.in/gemnasium/logrus-graylog-hook.v2"
	"os"
	"reflect"
	"regexp"
	"runtime"
	"strings"
	"time"
	"unicode"
)

var MaxInt64 = ^int64(0)

type Logger struct {
	logrus.Entry
	Level logrus.Level
}

// 日志日志
type Config struct {
	TimeFormat string
	Level      string
	Formatter  string
	Extra      map[string]interface{}
	StorageMode
}

// 日子打印方式
type StorageMode struct {
	File // 文件模式
}

// 文件
type File struct {
	Formatter string
	Path      map[string]string
	Rotation
}

// 文件切割时间和保存时间
type Rotation struct {
	Hours   uint
	Count   uint
	Postfix string
}

func NewLog(viper *viper.Viper, viperKey string) (l *Logger, err error) {
	var (
		config Config
	)
	// 获取日志模块
	if err = viper.UnmarshalKey(viperKey, &config); err != nil {
		return l, err
	}
	if config.TimeFormat != "" {
		config.TimeFormat = strings.TrimSpace(config.TimeFormat)
	}
	level := logrus.InfoLevel
	if config.Level != "" {
		config.Level = strings.TrimSpace(strings.ToLower(config.Level))
		if config.Level == "warn" {
			level = logrus.WarnLevel
		} else if config.Level == "debug" {
			level = logrus.DebugLevel
		} else if config.Level == "error" {
			level = logrus.ErrorLevel
		} else if config.Level == "fatal" {
			level = logrus.FatalLevel
		} else if config.Level == "panic" {
			level = logrus.PanicLevel
		} else if config.Level == "trace" {
			level = logrus.TraceLevel
		} else {
			level = logrus.InfoLevel
		}
	}
	var formatter logrus.Formatter
	if config.Formatter == "json" {
		formatter = &logrus.JSONFormatter{TimestampFormat: config.TimeFormat}
	} else {
		formatter = &logrus.TextFormatter{TimestampFormat: config.TimeFormat}
	}
	log := &logrus.Logger{
		Out:       os.Stderr,
		Formatter: formatter,
		Hooks:     make(logrus.LevelHooks),
		Level:     level,
	}

	if grayAddr, ok := viper.Get("logger.graylog.addr").(string); ok && len(grayAddr) > 0 {
		grayHook := graylog.NewGraylogHook(grayAddr, nil)
		log.AddHook(grayHook)
	}

	if config.Path != nil && len(config.Path) > 0 {
		writerMap := lfshook.WriterMap{}
		if v, ok := config.Path["panic"]; ok && v != "" {
			writerMap[logrus.PanicLevel], _ = rotatelogs.New(
				"./"+v+config.Postfix,
				rotatelogs.WithLinkName(v),                                         // 为最新的日志建立软连接，以方便随着找到当前日志文件
				rotatelogs.WithRotationCount(config.Count),                         // 设置文件清理前最多保存的个数，也可通过WithMaxAge设置最长保存时间，二者取其一
				rotatelogs.WithRotationTime(time.Duration(config.Hours)*time.Hour), // 设置日志分割的时间，例如一天一次
			)
		}
		if v, ok := config.Path["fatal"]; ok && v != "" {
			writerMap[logrus.FatalLevel], _ = rotatelogs.New(
				"./"+v+config.Postfix,
				rotatelogs.WithLinkName(v),
				rotatelogs.WithRotationCount(config.Count),
				rotatelogs.WithRotationTime(time.Duration(config.Hours)*time.Hour),
			)
		}
		if v, ok := config.Path["error"]; ok && v != "" {
			writerMap[logrus.ErrorLevel], _ = rotatelogs.New(
				"./"+v+config.Postfix,
				rotatelogs.WithLinkName(v),
				rotatelogs.WithRotationCount(config.Count),
				rotatelogs.WithRotationTime(time.Duration(config.Hours)*time.Hour),
			)
		}
		if v, ok := config.Path["warn"]; ok && v != "" {
			writerMap[logrus.WarnLevel], _ = rotatelogs.New(
				"./"+v+config.Postfix,
				rotatelogs.WithLinkName(v),
				rotatelogs.WithRotationCount(config.Count),
				rotatelogs.WithRotationTime(time.Duration(config.Hours)*time.Hour),
			)
		}
		if v, ok := config.Path["info"]; ok && v != "" {
			writerMap[logrus.InfoLevel], _ = rotatelogs.New(
				"./"+v+config.Postfix,
				rotatelogs.WithLinkName(v),
				rotatelogs.WithRotationCount(config.Count),
				rotatelogs.WithRotationTime(time.Duration(config.Hours)*time.Hour),
			)
		}
		if v, ok := config.Path["debug"]; ok && v != "" {
			writerMap[logrus.DebugLevel], _ = rotatelogs.New(
				"./"+v+config.Postfix,
				rotatelogs.WithLinkName(v),
				rotatelogs.WithRotationCount(config.Count),
				rotatelogs.WithRotationTime(time.Duration(config.Hours)*time.Hour),
			)
		}
		if v, ok := config.Path["trace"]; ok && v != "" {
			writerMap[logrus.TraceLevel], _ = rotatelogs.New(
				"./"+v+config.Postfix,
				rotatelogs.WithLinkName(v),
				rotatelogs.WithRotationCount(config.Count),
				rotatelogs.WithRotationTime(time.Duration(config.Hours)*time.Hour),
			)
		}
		var lfFormatter logrus.Formatter
		if config.File.Formatter == "json" {
			lfFormatter = &logrus.JSONFormatter{TimestampFormat: config.TimeFormat}
		} else {
			lfFormatter = &logrus.TextFormatter{TimestampFormat: config.TimeFormat}
		}
		lfHook := lfshook.NewHook(writerMap, lfFormatter)
		log.AddHook(lfHook)
	}
	entry := logrus.NewEntry(log)
	if config.Extra != nil && len(config.Extra) > 0 {
		entry = entry.WithFields(config.Extra)
	}
	l = &Logger{Entry: *entry, Level: level}
	return l, err
}

func (logger *Logger) Print(args ...interface{}) {
	if args == nil || len(args) == 0 {
		return
	}
	if tp, ok := args[0].(string); ok {
		tp = strings.ToLower(strings.TrimSpace(tp))
		if "sql" == tp && len(args) == 6 {
			logger.printSql(args...)
		} else {
			logger.WithCaller(2).Entry.Print(args...)
		}
	} else {
		logger.WithCaller(2).Entry.Print(args...)
	}
}

func (logger *Logger) WithField(key string, value interface{}) *Logger {
	return &Logger{Entry: *logger.Entry.WithField(key, value)}
}

func (logger *Logger) WithFields(fields map[string]interface{}) *Logger {
	return &Logger{Entry: *logger.Entry.WithFields(fields)}
}

func (logger *Logger) WithError(err error) *Logger {
	return &Logger{Entry: *logger.Entry.WithError(err)}
}

func (logger *Logger) WithCaller(skip int) *Logger {
	if _, ok := logger.Data["codeline"]; ok {
		return logger
	}
	if _, file, line, ok := runtime.Caller(skip); ok {
		return logger.
			WithField("codeline", fmt.Sprintf("%s:%d", file, line))
	}
	return logger
}

func (logger *Logger) Debugf(format string, args ...interface{}) {
	logger.WithCaller(2).Entry.Debugf(format, args...)
}

func (logger *Logger) Infof(format string, args ...interface{}) {
	logger.WithCaller(2).Entry.Infof(format, args...)
}

func (logger *Logger) Printf(format string, args ...interface{}) {
	logger.WithCaller(2).Entry.Printf(format, args...)
}

func (logger *Logger) Warnf(format string, args ...interface{}) {
	logger.WithCaller(2).Entry.Warnf(format, args...)
}

func (logger *Logger) Warningf(format string, args ...interface{}) {
	logger.WithCaller(2).Entry.Warningf(format, args...)
}

func (logger *Logger) Errorf(format string, args ...interface{}) {
	logger.WithCaller(2).Entry.Errorf(format, args...)
}

func (logger *Logger) Fatalf(format string, args ...interface{}) {
	logger.WithCaller(2).Entry.Fatalf(format, args...)
}

func (logger *Logger) Panicf(format string, args ...interface{}) {
	logger.WithCaller(2).Entry.Panicf(format, args...)
}

func (logger *Logger) Debug(args ...interface{}) {
	logger.WithCaller(2).Entry.Debug(args...)
}

func (logger *Logger) Info(args ...interface{}) {
	logger.WithCaller(2).Entry.Info(args...)
}

func (logger *Logger) Warn(args ...interface{}) {
	logger.WithCaller(2).Entry.Warn(args...)
}

func (logger *Logger) Warning(args ...interface{}) {
	logger.WithCaller(2).Entry.Warning(args...)
}

func (logger *Logger) Error(args ...interface{}) {
	logger.WithCaller(2).Entry.Error(args...)
}

func (logger *Logger) Fatal(args ...interface{}) {
	logger.WithCaller(2).Entry.Fatal(args...)
}

func (logger *Logger) Panic(args ...interface{}) {
	logger.WithCaller(2).Entry.Panic(args...)
}

func (logger *Logger) Debugln(args ...interface{}) {
	logger.WithCaller(2).Entry.Debugln(args...)
}

func (logger *Logger) Infoln(args ...interface{}) {
	logger.WithCaller(2).Entry.Infoln(args...)
}

func (logger *Logger) Println(args ...interface{}) {
	logger.WithCaller(2).Entry.Println(args...)
}

func (logger *Logger) Warnln(args ...interface{}) {
	logger.WithCaller(2).Entry.Warnln(args...)
}

func (logger *Logger) Warningln(args ...interface{}) {
	logger.WithCaller(2).Entry.Warningln(args...)
}

func (logger *Logger) Errorln(args ...interface{}) {
	logger.WithCaller(2).Entry.Errorln(args...)
}

func (logger *Logger) Fatalln(args ...interface{}) {
	logger.WithCaller(2).Entry.Fatalln(args...)
}

func (logger *Logger) Panicln(args ...interface{}) {
	logger.WithCaller(2).Entry.Panicln(args...)
}

func (logger *Logger) V(v int) bool {
	return false
}

var (
	sqlRegexp                = regexp.MustCompile(`\?`)
	numericPlaceHolderRegexp = regexp.MustCompile(`\$\d+`)
)

func (logger *Logger) printSql(args ...interface{}) {
	length := len(args)
	var (
		codeLine, sql string
		params        []interface{}
		latency       time.Duration
		rows          int64
		ok            bool
	)
	if length > 1 {
		codeLine, _ = args[1].(string)
	}
	if length > 2 {
		latency, _ = args[2].(time.Duration)
	}
	if length > 3 {
		sql, ok = args[3].(string)
		if ok {
			sql = strings.TrimSpace(strings.Replace(strings.Replace(strings.Replace(sql, "\r\n", " ", -1), "\n", " ", -1), "\t", " ", -1))
		}
	}
	if length > 4 {
		params, _ = args[4].([]interface{})
	}
	if length > 5 {
		rows, _ = args[5].(int64)
	}
	lg := logger.
		WithField("tag", "SQL").
		WithField("sql", logger.getSql(sql, params))
	if len(codeLine) > 0 {
		lg = lg.WithField("codeline", strings.TrimSpace(codeLine))
	} else {
		lg = lg.WithCaller(9)
	}
	if latency > 0 {
		lg = lg.WithField("latency", fmt.Sprintf("%v", latency))
	}
	if rows != MaxInt64 {
		lg = lg.WithField("rows", fmt.Sprintf("%d rows affected or returned", rows))
	}
	if len(params) <= 0 {
		lg.Info(fmt.Sprintf("%s;", sql))
	} else {
		lg.Info(fmt.Sprintf("%s %v", sql, params))
	}
}

func (logger *Logger) getSql(originSql string, params []interface{}) string {
	var formattedValues []string
	for _, value := range params {
		indirectValue := reflect.Indirect(reflect.ValueOf(value))
		if indirectValue.IsValid() {
			value = indirectValue.Interface()
			if t, ok := value.(time.Time); ok {
				formattedValues = append(formattedValues, fmt.Sprintf("'%v'", t.Format("2006-01-02 15:04:05")))
			} else if b, ok := value.([]byte); ok {
				if str := string(b); logger.isPrintable(str) {
					formattedValues = append(formattedValues, fmt.Sprintf("'%v'", str))
				} else {
					formattedValues = append(formattedValues, "'<binary>'")
				}
			} else if r, ok := value.(driver.Valuer); ok {
				if value, err := r.Value(); err == nil && value != nil {
					formattedValues = append(formattedValues, fmt.Sprintf("'%v'", value))
				} else {
					formattedValues = append(formattedValues, "NULL")
				}
			} else {
				formattedValues = append(formattedValues, fmt.Sprintf("'%v'", value))
			}
		} else {
			formattedValues = append(formattedValues, "NULL")
		}
	}
	if nil == formattedValues {
		return ""
	}

	var sql string
	// differentiate between $n placeholders or else treat like ?
	if numericPlaceHolderRegexp.MatchString(originSql) {
		for index, value := range formattedValues {
			placeholder := fmt.Sprintf(`\$%d([^\d]|$)`, index+1)
			sql = regexp.MustCompile(placeholder).ReplaceAllString(originSql, value+"$1")
		}
	} else {
		formattedValuesLength := len(formattedValues)
		for index, value := range sqlRegexp.Split(originSql, -1) {
			sql += value
			if index < formattedValuesLength {
				sql += formattedValues[index]
			}
		}
	}
	return sql
}

func (logger *Logger) isPrintable(s string) bool {
	for _, r := range s {
		if !unicode.IsPrint(r) {
			return false
		}
	}
	return true
}
