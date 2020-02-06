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

func InitLogsToFile(filePath string, configRewrite ...rollingwriter.Config) error {

	// Paths
	logPath := path.Dir(filePath)
	fileName := strings.ReplaceAll(path.Base(filePath), path.Ext(filePath), "")

	// Main config
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

	// Rewrite config
	if len(configRewrite) > 0 {
		configRewrite[0].LogPath = config.LogPath
		configRewrite[0].FileName = config.FileName
		config = configRewrite[0]
	}

	// Create a writer
	log.w, err = rollingwriter.NewWriterFromConfig(&config)
	if err != nil {
		return err
	}

	// Log to file and console
	log.logger = zerolog.New(io.MultiWriter(log.w, os.Stdout)).With().Timestamp().Logger()

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
	log.logger.Debug().Caller(1).Msg(text)
}

func DebugF(format string, v ...interface{}) {
	log.logger.Debug().Caller(1).Msgf(format, v...)
}

func Info(text string) {
	log.logger.Info().Caller(1).Msg(text)
}

func InfoF(format string, v ...interface{}) {
	log.logger.Info().Caller(1).Msgf(format, v...)
}

func Warn(text string) {
	log.logger.Warn().Caller(1).Msg(text)
}

func WarnF(format string, v ...interface{}) {
	log.logger.Warn().Caller(1).Msgf(format, v...)
}

func Error(text string) {
	log.logger.Error().Caller(1).Msg(text)
}

func ErrorF(format string, v ...interface{}) {
	log.logger.Error().Caller(1).Msgf(format, v...)
}

func Fatal(text string) {
	log.logger.Fatal().Caller(1).Msg(text)
}

func FatalF(format string, v ...interface{}) {
	log.logger.Fatal().Caller(1).Msgf(format, v...)
}

func SendErr(err error) {
	log.logger.Err(err).Caller(1).Send()
}
