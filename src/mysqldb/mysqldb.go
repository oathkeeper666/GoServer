package mysqldb

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"logger"
)

var connector DBConnector

type DBConnector struct {
	db *sql.DB
}

// connect db
func ConnectToDb() bool {
	if connector.db != nil {
		logger.WRITE_WARNING("db have been connected.")
		return false
	}
	var err error
	connector.db, err = sql.Open("mysql", "wubang:123456@unix(/tmp/mysql.sock)/mysql")	// user:password@protocol/dbname
	if err != nil {
		logger.WRITE_ERROR("open database failed, error is %v", err)
		return false
	}

	/*var name string
	var host string
	row := connector.db.QueryRow("select Host, user from user where user = ?;", "wubang")
	err2 := row.Scan(&host, &name);
	if err2 == nil {
		logger.WRITE_DEBUG("host, name = %s, %s", host, name)
	} else {
		logger.WRITE_WARNING("err2 is %v", err2)
	}*/

	return true
}

// close db connection
func CloseDB() {
	if connector.db != nil {
		connector.db.Close()
	}
}


// 判断数据库连接是否中断
func IsConnected() bool {
	if connector.db != nil {
		return true
	}
	return false
}