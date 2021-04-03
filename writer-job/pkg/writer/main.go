package writer

import (
	"database/sql"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
)

func CheckErr(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

func InsertRows(tableName *string, noDryRunPtr *bool, db *sql.DB, sleepTimeBetweenInsertsPtr *int64, rowsToInsert *int64) {
	selectQuery := fmt.Sprintf("select created_on, updated_on, product_id from %s where id >= 1 limit 1;", *tableName)
	log.Debug(selectQuery)
	results, err := db.Query(selectQuery)
	CheckErr(err)
	cols, err := results.Columns()
	CheckErr(err)
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
	CheckErr(err)

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
			CheckErr(err)
			result, err := insertTxn.Exec(destinationInsertQuery, valueArgs...)
			CheckErr(err)
			affected, err := result.RowsAffected()
			CheckErr(err)
			log.Debug(fmt.Sprintf("RowsAffected: %d", affected))
			err = insertTxn.Commit()
			CheckErr(err)
			if *sleepTimeBetweenInsertsPtr != 0 {
				log.Info(fmt.Sprintf("Relaxing for %d Seconds", *sleepTimeBetweenInsertsPtr))
				time.Sleep(time.Duration(*sleepTimeBetweenInsertsPtr) * time.Second)
			}
		}
	} else {
		log.WithFields(log.Fields{"noDryRunPtr": true}).Info("No changes done to db.")
	}
}
