package db

import (
	"database/sql"

	"example.com/todo-app/constants/db_const"
	"example.com/todo-app/structs"
	_ "modernc.org/sqlite"
)

var DB *sql.DB

func SetUpDb() {
	var err error
	DB, err = sql.Open(db_const.SQL_TYPE, db_const.DB_NAME)

	if err != nil {
		panic(db_const.ERR_DB_CONN)
	}
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createDbTables()
}

func createDbTables() {
	sqlsToRun := []structs.SqlStmtAndErr{
		{
			SqlStmt: db_const.CREATE_USERS_TABLE,
			Err:     db_const.ERR_CREATE_USERS,
		},
		{
			SqlStmt: db_const.CREATE_TODOS_TABLE,
			Err:     db_const.ERR_CREATE_TODOS,
		},
		{
			SqlStmt: db_const.CREATE_TODO_USERS_TABLE,
			Err:     db_const.ERR_CREATE_TODO_USERS,
		},
	}

	for _, element := range sqlsToRun {
		execSqlOrPanic(element)
	}
}

func execSqlOrPanic(sqlStmtAndErr structs.SqlStmtAndErr) {
	_, err := DB.Exec(sqlStmtAndErr.SqlStmt)
	if err != nil {
		panic(sqlStmtAndErr.Err)
	}
}
