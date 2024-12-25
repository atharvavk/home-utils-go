package app

import (
	"database/sql"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type EnvConfig struct {
	DBHost     string `json:"dbhost"`
	DBUser     string `json:"dbusername"`
	DBPassword string `json:"dbpassword"`
	DBSchema   string `json:"dbschema"`
	AppPort    int    `json:"appport"`
}

type AppContext struct {
	ServerPort int
	Sql        *sql.DB
	Logger     *zap.Logger
}

func IntializeApp() AppContext {
	envConfig := readEnvConfig()
	dbConn := createDBConnection(envConfig)
	logger := initializeLogger()
	return AppContext{
		ServerPort: envConfig.AppPort,
		Sql:        dbConn,
		Logger:     logger,
	}
}

func readEnvConfig() EnvConfig {
	conf := viper.New()
	conf.SetConfigFile(".env.json")
	err := conf.ReadInConfig()
	if err != nil {
		panic(err)
	}
	var env EnvConfig
	err = conf.Unmarshal(&env)
	if err != nil {
		panic(err)
	}
	return env
}

func createDBConnection(env EnvConfig) *sql.DB {
	mysqlConfig := mysql.Config{
		DBName:               env.DBSchema,
		User:                 env.DBUser,
		Addr:                 env.DBHost,
		Passwd:               env.DBPassword,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	}
	conn, err := sql.Open("mysql", mysqlConfig.FormatDSN())
	if err != nil {
		panic(err)
	}
	return conn
}

func initializeLogger() *zap.Logger {
	logLevel := zap.InfoLevel
	consoleEncoder := zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
		MessageKey:    "msg",
		LevelKey:      "level",
		TimeKey:       "time",
		NameKey:       "logger",
		CallerKey:     "caller",
		StacktraceKey: "stacktrace",
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeLevel:   zapcore.LowercaseLevelEncoder,
		EncodeTime:    zapcore.ISO8601TimeEncoder,
		EncodeCaller:  zapcore.ShortCallerEncoder,
	})
	stdout := zapcore.Lock(os.Stdout)
	core := zapcore.NewCore(consoleEncoder, stdout, logLevel)
	logger := zap.New(core)
	return logger
}
