package main

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	flag "github.com/spf13/pflag"
)

func validateEnvVars() map[string]string {
	var requiredDbEnvVars = [5]string{
		"DB_NAME",
		"DB_HOST",
		"DB_USER",
		"DB_PASSWORD",
		"DB_PORT"}

	dbConfig := make(map[string]string)
	for _, envVar := range requiredDbEnvVars {
		envVarValue, isSet := os.LookupEnv(envVar)
		if !isSet {
			log.Fatal(fmt.Sprintf("%s is a required env variable. Please set it and with a valid value ofcourse.", envVar))
		}
		dbConfig[envVar] = envVarValue
	}

	return dbConfig
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

func main() {
	var requiredLogLevel string
	flag.StringVarP(&requiredLogLevel, "log-level", "l", "info", "Required log level. debug/info/warn/error.")
	var printHelp bool
	flag.BoolVarP(&printHelp, "help", "h", false, "Prints this help content.")
	tablePtr := flag.String("table", "", "table name")
	noDryRunPtr := flag.Bool("no-dryrun", false, "dry-run option. required for performing insert")
	maxOpenConnectionsToDbPtr := flag.Int("max-open-conns-db", 200, "Max Open Connections to the Db.")
	maxIdleConnectionsToDbPtr := flag.Int("max-idle-conns-db", 10, "Max Idle Connections to the Db.")
	flag.Parse()
	if printHelp {
		flag.Usage()
		return
	}
	switch requiredLogLevel {
	case "debug":
		log.SetLevel(log.DebugLevel)
	case "info":
		log.SetLevel(log.InfoLevel)
	case "warn":
		log.SetLevel(log.WarnLevel)
	case "error":
		log.SetLevel(log.ErrorLevel)
	default:
		log.SetLevel(log.InfoLevel)
	}
	log.SetFormatter(&log.JSONFormatter{
		TimestampFormat: time.RFC3339,
	})
	log.SetOutput(os.Stdout)

	log.Info("I am a writer")

	dbConfig := validateEnvVars()

	dSNString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", dbConfig["DB_USER"], dbConfig["DB_STORE_CONFIG_PASSWORD"], dbConfig["DB_HOST"], dbConfig["DB_PORT"], dbConfig["DB_NAME"])
	log.Debug(fmt.Sprintf("dSNString: %s", dSNString))
	db, err := sql.Open("mysql", dSNString)
	checkErr(err)
	db.SetMaxOpenConns(*maxOpenConnectionsToDbPtr)
	db.SetMaxIdleConns(*maxIdleConnectionsToDbPtr)

	err = db.Ping()
	checkErr(err)

	log.Info(dSNString)
	log.Info(*tablePtr)
	log.Info(*noDryRunPtr)
}
