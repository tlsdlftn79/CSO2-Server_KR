package mysql

import (
	"database/sql"

	. "github.com/KouKouChan/CSO2-Server/configure"
	. "github.com/KouKouChan/CSO2-Server/kerlong"
)

const (
	sqlPath_BuildTable = "/CSO2-Server/database/sql/buildTable.sql"
	sqlPath_BuildDB    = "/CSO2-Server/database/sql/buildDB.sql"
)

var (
	DBtable = "CSO2Server"
)

func InitDatabase(path string) (*sql.DB, error) {
	DB, err := sql.Open("mysql", Conf.DBUserName+":"+Conf.DBpassword+"@tcp("+Conf.DBaddress+":"+Conf.DBport+")/?charset=utf8&multiStatements=true")

	sqlDB := ReadStringFromFile(path + sqlPath_BuildDB)

	_, err = DB.Exec(sqlDB)
	if err != nil {
		return DB, err
	}

	DB.Close()

	return sql.Open("mysql", Conf.DBUserName+":"+Conf.DBpassword+"@tcp("+Conf.DBaddress+":"+Conf.DBport+")/"+DBtable+"?charset=utf8&multiStatements=true")
}

func CheckDataBaseTable(DB *sql.DB, path string) error {
	sqlTable := ReadStringFromFile(path + sqlPath_BuildTable)

	_, err := DB.Exec(sqlTable)
	return err
}
