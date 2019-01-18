package model

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"sync"
)

var (
	sql_create_downloadtask string = `
	CREATE TABLE IF NOT EXISTS t_task (
	F_taskid TEXT PRIMARY KEY,
 	F_business TEXT NULL,
	F_subbusiness TEXT NULL,
	F_url TEXT NULL,
	F_user TEXT NULL,
	F_dir TEXT NULL,
	F_packname TEXT NULL,
	F_checkhash TEXT NULL,
	F_finish INTEGER NULL,
	F_total INTEGER NULL,
	F_cur_offset INTEGER NULL,
	F_status INTEGER NULL,
	F_method TEXT NULL,
	F_header TEXT NULL,
	F_body 	 TEXT NULL,
	F_ext	 TEXT NULL,
	F_create_time TEXT NULL,
	F_modify_time TEXT NULL)
`
	sql_index_list = []string{
		"create index IF NOT EXISTS I_url on t_task(F_url)",
	}

	instance *sql.DB

	RWlock sync.RWMutex
)

const DOWNLOADER_DB = "/usr/local/dbfile/downloader"

func GetDbInstance() *sql.DB {
	if instance == nil {
		var err error
		instance, err = sql.Open("sqlite3", DOWNLOADER_DB)
		if err != nil {
			log.Printf("open db file err : %s", err.Error())
			return nil
		}
		log.Println("init sqlite db done")
	}
	return instance
}

func init() {
	execSql(sql_create_downloadtask)

	for _, v := range sql_index_list {
		execSql(v)
	}
}

func execSql(sql string) (rowAffect int64, err error) {
	db := GetDbInstance()
	stmt, err := db.Prepare(sql)
	defer CloseStmt(stmt)
	if err != nil {
		log.Printf("prepare sql : %s err : %s\n", sql, err.Error())
		return 0, err
	}
	res, err := stmt.Exec()
	if err != nil {
		log.Printf("exec sql : %s err : %s\n", sql, err.Error())
		return 0, err
	}
	return res.RowsAffected()
}

func CloseStmt(s *sql.Stmt) {
	if s != nil {
		s.Close()
	}
}
