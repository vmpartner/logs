package logs

import (
	"github.com/arthurkiller/rollingwriter"
	"github.com/rs/zerolog"
	"io"
	"os"
	"path"
	"strings"
)

type Log struct {
	logger zerolog.Logger
	w      rollingwriter.RollingWriter
}

var log *Log
var err error

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log = new(Log)
	log.logger = zerolog.New(os.Stdout).With().Timestamp().Caller().Logger()
}

func InitLogsToFile(filePath string) error {

	logPath := path.Dir(filePath)
	fileName := strings.ReplaceAll(path.Base(filePath), path.Ext(filePath), "")

	config := rollingwriter.Config{
		LogPath:                logPath,
		TimeTagFormat:          "060102150405",
		FileName:               fileName,
		MaxRemain:              5,
		RollingPolicy:          rollingwriter.VolumeRolling,
		RollingTimePattern:     "* * * * * *",
		RollingVolumeSize:      "64M",
		WriterMode:             "async",
		BufferWriterThershould: 8 * 1024 * 1024,
		Compress:               true,
	}

	// Create a writer
	log.w, err = rollingwriter.NewWriterFromConfig(&config)
	if err != nil {
		return err
	}

	// Log to file and console
	log.logger = zerolog.New(io.MultiWriter(log.w, os.Stdout)).With().Timestamp().Caller().Logger()

	return nil
}

func Close() {
	if log.w != nil {
		log.w.Close()
	}
	log = nil
}

func Logger() *zerolog.Logger {
	return &log.logger
}

func Debug(text string) {
	log.logger.Debug().Msg(text)
}

func DebugF(format string, v ...interface{}) {
	log.logger.Debug().Msgf(format, v)
}

func Info(text string) {
	log.logger.Info().Msg(text)
}

func InfoF(format string, v ...interface{}) {
	log.logger.Info().Msgf(format, v)
}

func Warn(text string) {
	log.logger.Warn().Msg(text)
}

func WarnF(format string, v ...interface{}) {
	log.logger.Warn().Msgf(format, v)
}

func Error(text string) {
	log.logger.Error().Msg(text)
}

func ErrorF(format string, v ...interface{}) {
	log.logger.Error().Msgf(format, v)
}

func Fatal(text string) {
	log.logger.Fatal().Msg(text)
}

func FatalF(format string, v ...interface{}) {
	log.logger.Fatal().Msgf(format, v)
}
