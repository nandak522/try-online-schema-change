package reader

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	log "github.com/sirupsen/logrus"
)

func CheckErr(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

type BaseItem struct {
	CreatedOn string
	UpdatedOn string
	ProductID string
}

func ReadData(db *sql.DB) *BaseItem {
	selectQuery := "select created_on, updated_on, product_id from try_osc.baseitem where id >= 1 limit 1;"
	log.Debug(selectQuery)
	results, err := db.Query(selectQuery)
	CheckErr(err)
	defer results.Close()

	var bi BaseItem
	results.Next()
	err = results.Scan(&bi.CreatedOn, &bi.UpdatedOn, &bi.ProductID)
	CheckErr(err)

	return &bi
}
