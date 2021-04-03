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

func insertRows(tableName *string, noDryRunPtr *bool, db *sql.DB, sleepTimeBetweenInsertsPtr *int64, rowsToInsert *int64) {
	selectQuery := fmt.Sprintf("select created_on, updated_on, product_id from %s where id >= 1 limit 1;", *tableName)
	log.Debug(selectQuery)
	results, err := db.Query(selectQuery)
	checkErr(err)
	cols, err := results.Columns()
	checkErr(err)
	columnsLength := len(cols)

	// Create a slice of interface{}'s to represent each column,
	// and a second slice to contain pointers to each item in the columns slice.
	columns := make([]interface{}, len(cols))
	columnPointers := make([]interface{}, len(cols))
	for i := range columns {
		columnPointers[i] = &columns[i]
	}
	results.Next()
	// Scan the result into the column pointers...
	err = results.Scan(columnPointers...)
	checkErr(err)

	valueArgs := make([]interface{}, 0, columnsLength)

	for i := range columns {

		// TODO: This columnPointers[i].(*interface{}) interpretation has to be clearly understood
		columnValue := columnPointers[i].(*interface{})
		valueArgs = append(valueArgs, *columnValue)
	}
	destinationInsertQuery := fmt.Sprintf("insert into %s(created_on, updated_on, product_id) values(?, ?, ?)", *tableName)
	log.Debug(destinationInsertQuery)
	if *noDryRunPtr {
		log.WithFields(log.Fields{"noDryRunPtr": false}).Info("Changes will be done to db.")
		for index := int64(0); index < *rowsToInsert; index++ {
			insertTxn, err := db.Begin()
			checkErr(err)
			result, err := insertTxn.Exec(destinationInsertQuery, valueArgs...)
			checkErr(err)
			affected, err := result.RowsAffected()
			checkErr(err)
			log.Debug(fmt.Sprintf("RowsAffected: %d", affected))
			err = insertTxn.Commit()
			checkErr(err)
			if *sleepTimeBetweenInsertsPtr != 0 {
				log.Info(fmt.Sprintf("Relaxing for %d Seconds", *sleepTimeBetweenInsertsPtr))
				time.Sleep(time.Duration(*sleepTimeBetweenInsertsPtr) * time.Second)
			}
		}
	} else {
		log.WithFields(log.Fields{"noDryRunPtr": true}).Info("No changes done to db.")
	}
}

func main() {
	var requiredLogLevel string
	flag.StringVarP(&requiredLogLevel, "log-level", "l", "info", "Required log level. debug/info/warn/error.")
	var printHelp bool
	flag.BoolVarP(&printHelp, "help", "h", false, "Prints this help content.")
	var tableName string
	flag.StringVarP(&tableName, "table", "t", "", "Required table name.")
	noDryRunPtr := flag.Bool("nodryrun", false, "dry-run option. required to be disabled explicitly for performing inserts")
	maxOpenConnectionsToDbPtr := flag.Int("max-open-conns-db", 200, "Max Open Connections to the Db.")
	maxIdleConnectionsToDbPtr := flag.Int("max-idle-conns-db", 10, "Max Idle Connections to the Db.")
	var sleepTimeBetweenInsertsPtr int64
	flag.Int64VarP(&sleepTimeBetweenInsertsPtr, "sleep-between-inserts", "s", 0, "sleep time in seconds between inserts, to do db commits")
	var rowsToInsert int64
	flag.Int64VarP(&rowsToInsert, "rows-to-insert", "r", 1, "Rows to insert.")
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
		TimestampFormat:   time.RFC3339,
		DisableHTMLEscape: true,
	})
	log.SetOutput(os.Stdout)
	if tableName == "" {
		log.Fatal("Please supply the table name using --table")
	}

	log.Info("I am a writer")

	dbConfig := validateEnvVars()

	dsnString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", dbConfig["DB_USER"], dbConfig["DB_PASSWORD"], dbConfig["DB_HOST"], dbConfig["DB_PORT"], dbConfig["DB_NAME"])
	log.Debug("dsnString: ", dsnString)
	db, err := sql.Open("mysql", dsnString)
	checkErr(err)
	db.SetMaxOpenConns(*maxOpenConnectionsToDbPtr)
	db.SetMaxIdleConns(*maxIdleConnectionsToDbPtr)

	err = db.Ping()
	checkErr(err)

	insertRows(&tableName, noDryRunPtr, db, &sleepTimeBetweenInsertsPtr, &rowsToInsert)
	log.Info("load finished")
}
