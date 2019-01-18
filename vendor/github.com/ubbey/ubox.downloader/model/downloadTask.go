package model

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"time"
)

type DownloadTask struct {
	Taskid      string `json:"taskid"`
	Business    string `json:"business"`
	Subbusiness string `json:"subbusiness"`
	User        string `json:"user"`
	Url         string `json:"url"`
	Dir         string `json:"dir"`
	Name        string `json:"name"`
	CheckHash   string `json:"check_hash"`
	CurOffset   int64  `json:"cur_offset"`
	Finish      int64  `json:"finish"`
	Total       int64  `json:"total"`
	Status      int    `json:"status"`
	Method      string `json:"method"`
	Header      string `json:"header"`
	Body        string `json:"body"`
	Ext         string `json:"ext"`
	CreateTime  string `json:"create_time"`
	ModifyTime  string `json:"modify_time"`
}

var task_filed = `F_taskid, F_business, F_subbusiness, F_user, F_url, F_dir, F_packname, F_checkhash, F_cur_offset, F_finish, 
 F_total, F_status, F_method, F_header, F_body, F_ext, F_create_time, F_modify_time`

func get_nowtime() string {
	timestamp := time.Now().Unix()
	return time.Unix(timestamp, 0).Format("2006-01-02 15:04:05") //go 语言 固定的时间格式

}

func (dt *DownloadTask) Generateid() {
	dt.ModifyTime = get_nowtime()
	dt.CreateTime = get_nowtime()
	bt, _ := json.Marshal(*dt)
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(bt))

	id := hex.EncodeToString(md5Ctx.Sum(nil))
	dt.Taskid = id
}

func (dt *DownloadTask) QueryDownloadTaskByid() (t DownloadTask, e error) {
	RWlock.RLock()
	defer RWlock.RUnlock()

	db := GetDbInstance()
	sql := `
		select ` + task_filed + `
		from t_task
		where F_taskid = ?
	`
	stmt, err := db.Prepare(sql)
	defer CloseStmt(stmt)
	if err != nil {
		return DownloadTask{}, err
	}

	_ = stmt.QueryRow(dt.Taskid).Scan(&t.Taskid, &t.Business, &t.Subbusiness, &t.User, &t.Url, &t.Dir, &t.Name, &t.CheckHash, &t.CurOffset, &t.Finish,
		&t.Total, &t.Status, &t.Method, &t.Header, &t.Body, &t.Ext, &t.CreateTime, &t.ModifyTime)

	return
}

func (dt *DownloadTask) QueryAllDownloadTask(business string) (tsList []DownloadTask, e error) {
	RWlock.RLock()
	defer RWlock.RUnlock()

	db := GetDbInstance()
	sql := `
		select ` + task_filed + `
		from t_task where F_business = ?
	`
	stmt, err := db.Prepare(sql)
	defer CloseStmt(stmt)
	if err != nil {
		return nil, err
	}

	rs, _ := stmt.Query(business)

	for rs.Next() {
		t := DownloadTask{}
		rs.Scan(&t.Taskid, &t.Business, &t.Subbusiness, &t.User, &t.Url, &t.Dir, &t.Name, &t.CheckHash, &t.CurOffset, &t.Finish,
			&t.Total, &t.Status, &t.Method, &t.Header, &t.Body, &t.Ext, &t.CreateTime, &t.ModifyTime)
		tsList = append(tsList, t)
	}

	return
}

func (dt *DownloadTask) NewDownloadTask() (int64, error) {
	RWlock.Lock()
	defer RWlock.Unlock()

	db := GetDbInstance()
	sql := `
		insert into t_task(` + task_filed + `)
		values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)
	`
	stmt, err := db.Prepare(sql)
	defer CloseStmt(stmt)
	if err != nil {
		return 0, err
	}

	rs, err := stmt.Exec(dt.Taskid, dt.Business, dt.Subbusiness, dt.User, dt.Url, dt.Dir, dt.Name, dt.CheckHash, dt.CurOffset, dt.Finish, dt.Total, dt.Status,
		dt.Method, dt.Header, dt.Body, dt.Ext, dt.CreateTime, dt.ModifyTime)
	if err != nil {
		return 0, err
	}

	return rs.LastInsertId()
}

func (dt *DownloadTask) RemoveDownloadTaskByid() error {
	RWlock.Lock()
	defer RWlock.Unlock()

	db := GetDbInstance()
	sql := `
		delete from t_task
		where F_id = ?
	`
	stmt, err := db.Prepare(sql)
	defer CloseStmt(stmt)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(dt.Taskid)
	if err != nil {
		return err
	}

	return nil
}

func (dt *DownloadTask) UpdateDownloadTaskByid() (int64, error) {
	RWlock.Lock()
	defer RWlock.Unlock()

	db := GetDbInstance()
	sql := `
		update t_task 
		set F_url = ?, F_business = ?, F_subbusiness = ?, F_user = ?, F_dir = ?, F_packname = ?, F_checkhash = ?, F_cur_offset = ?, F_finish = ?, F_total = ?, F_status = ?, F_method = ?, F_header = ?, F_body = ?, F_ext = ?, F_modify_time = ?
		where F_taskid = ?
	`
	stmt, err := db.Prepare(sql)
	defer CloseStmt(stmt)
	if err != nil {
		return 0, err
	}

	rs, err := stmt.Exec(dt.Url, dt.Business, dt.Subbusiness, dt.User, dt.Dir, dt.Name, dt.CheckHash, dt.CurOffset, dt.Finish, dt.Total, dt.Status,
		dt.Method, dt.Header, dt.Body, dt.Ext, get_nowtime(), dt.Taskid)
	if err != nil {
		return 0, err
	}

	return rs.RowsAffected()
}
