/***********************************************************************
// Copyright yqtc.com @2017 The source code.
// Copyright (c) 2009-2016 The Bitcoin Core developers
// Distributed under the MIT software license, see the accompanying
// file COPYING or http://www.opensource.org/licenses/mit-license.php.
//******
// Filename: sdk.go
// Description: 盒子提供给uapp调用的SDK接口
// Author: licowei
// CreateTime: 2019-01-16
/***********************************************************************/
package sdk

import (
	"encoding/json"
	//	"github.com/yqtc.com/ubox.golib/log"
	"github.com/yqtc.com/ubox.uapp/uvm/sdk/syscall"
)

/*
	interface
*/

const (
	// user interface
	SDK_USER_GET_USER = "/sdk/user/get_user"

	// device interface
	SDK_DEVICE_REGISTER_UPNP = "/sdk/device/register_upnp"
	SDK_DEVICE_OPEN_UPNP     = "/sdk/device/open_upnp"

	// samba interface
	SDK_SAMBA_OPERATE = "/sdk/samba/operate"
	SDK_SAMBA_CONFIG  = "/sdk/samba/config"

	// downloader interface
	SDK_DOWNLOADER_DOWNLOAD      = "/sdk/downloader/download"
	SDK_DOWNLOADER_TASK_INFO     = "/sdk/downloader/task_info"
	SDK_DOWNLOADER_DELETE        = "/sdk/downloader/delete"
	SDK_DOWNLOADER_GET_LIST      = "/sdk/downloader/get_list"
	SDK_DOWNLOADER_CHANGE_STATUS = "/sdk/downloader/change_status"
	SDK_DOWNLOADER_REMOVEALL     = "/sdk/downloader/removeall"

	// disk interface
	SDK_DISK_CLEAR_UERDATA = "/sdk/disk/clear_userdata"

	UAPP_PVR = "PVR"

	UPNP_EXTENDKEY_BOXID    = "boxId"
	UPNP_EXTENDKEY_VERSION  = "version"
	UPNP_EXTENDKEY_MAC      = "mac"
	UPNP_EXTENDKEY_BINDUSER = "bindUser"
)

//// Discovery /////////

//// Download  /////////

//// HttpServer/////////

//// File      /////////

//// Web3	   /////////

//// device  /////

/*
 注册upnp接口
*/
//刚启动进程调用，每次头盔绑定，解绑时调用（绑定头盔改变得瑟时候调用）
type Sdk_device_registerUpnp_req struct {
	Upnpst       string   `json:"upnpst"`       //固定为"pvr"
	Description  string   `json:"description"`  //xml,config文件，（xml，里面有绑定的头盔字段）
	Search_delay int      `json:"search_delay"` //传过去的时间（3分钟）
	Extend_key   []string `json:"extend_key"`   //固定这两个扩展字段boxId，version等
	Ext          string   `json:"ext"`          //预留
}
type Sdk_device_registerUpnp_rsp struct {
	Err_no int `json:"err_no"`
	/*
		-1001 - 请求解析失败
		1002  - 参数为空
	*/
	Err_msg string `json:"err_msg"`
}

func Sdk_device_registerUpnp(req Sdk_device_registerUpnp_req) (Sdk_device_registerUpnp_rsp, error) {

	rsp := Sdk_device_registerUpnp_rsp{}

	dret, err := syscall.ExecSysFunc(SDK_DEVICE_REGISTER_UPNP, req)
	if err != nil {
		return rsp, err
	}
	bret, _ := json.Marshal(dret)

	err = json.Unmarshal(bret, &rsp)

	return rsp, err
}

/*
 打开upnp搜索开关
*/
type Sdk_device_openUpnp_req struct {
	Upnpst string `json:"upnpst"` //固定为"pvr"
	Ext    string `json:"ext"`    //预留
}

//重复两次调用会返回错误码，先调用此接口，成功才更新pvr_token
type Sdk_device_openUpnp_rsp struct {
	Err_no int `json:"err_no"`
	/*
		    -1001 - 请求解析失败
			1002  - 参数为空
			11001 - 未注册UPNP
			11002 - 开关已打开
	*/
	Err_msg string `json:"err_msg"`
	Delay   int64  `json:"delay"`
}

func (r *Sdk_device_openUpnp_rsp) Result(errno int, errmsg string) {
	r.Err_no = errno
	r.Err_msg = errmsg
	//log.TRACE("SDK OpenUpnpRsp err_no : %d , err_msg : %s", errno, errmsg)
}

func Sdk_device_openUpnp(req Sdk_device_openUpnp_req) (Sdk_device_openUpnp_rsp, error) {

	rsp := Sdk_device_openUpnp_rsp{}

	dret, err := syscall.ExecSysFunc(SDK_DEVICE_OPEN_UPNP, req)
	if err != nil {
		return rsp, err
	}
	bret, _ := json.Marshal(dret)

	err = json.Unmarshal(bret, &rsp)

	return rsp, err
}

////  user  /////

/*
 获取用户信息接口
*/

type Sdk_user_getUser_req struct {
	Token string `json:"token"`
	Ext   string `json:"ext"`
}
type Sdk_user_getUser_rsp struct {
	Err_no int `json:"err_no"`
	/*
		-1001 - 请求解析失败
		1002  - 参数为空
		1103  - token过期
	*/
	Err_msg  string `json:"err_msg"`
	Username string `json:"username"`
	BoxId    string `json:"boxId"`
} //token失效错误码，只判断是否返回错误

func Sdk_user_getUser(req Sdk_user_getUser_req) (Sdk_user_getUser_rsp, error) {

	rsp := Sdk_user_getUser_rsp{}

	dret, err := syscall.ExecSysFunc(SDK_USER_GET_USER, req)
	if err != nil {
		return rsp, err
	}
	bret, _ := json.Marshal(dret)

	err = json.Unmarshal(bret, &rsp)

	return rsp, err
}

////  samba  /////

/*
 配置samba服务接口
*/

type Sdk_samba_config_req struct {
	Business     string `json:"business"`     //固定为"pvr"
	Sub_business string `json:"sub_business"` //头盔sn
	Ext          string `json:"ext"`          //预留
}
type Sdk_samba_config_rsp struct {
	Err_no int `json:"err_no"`
	/*
		-1001 - 请求解析失败
		1002  - 参数为空
	*/
	Err_msg     string `json:"err_msg"`
	Samba_uname string `json:"samba_uname"`
	Samba_pwd   string `json:"samba_pwd"`
	Root_path   string `json:"root_path"`
}

//绑定时调用
func Sdk_samba_config(req Sdk_samba_config_req) (Sdk_samba_config_rsp, error) {

	rsp := Sdk_samba_config_rsp{}

	dret, err := syscall.ExecSysFunc(SDK_SAMBA_CONFIG, req)
	if err != nil {
		return rsp, err
	}
	bret, _ := json.Marshal(dret)

	err = json.Unmarshal(bret, &rsp)

	return rsp, err
}

/*
 samba操作接口
*/

type Sdk_samba_operate_req struct {
	Action int    `json:"action"` //1，2，3（重启）
	Ext    string `json:"ext"`
}
type Sdk_samba_operate_rsp struct {
	Err_no int `json:"err_no"`
	/*
		-1001 - 请求解析失败
		1002  - 参数为空
	*/
	Err_msg     string `json:"err_msg"`
	Samba_uname string `json:"samba_uname"`
	Samba_pwd   string `json:"samba_pwd"`
	Root_path   string `json:"root_path"`
}

//头显连接时调用
func Sdk_samba_operate(req Sdk_samba_operate_req) (Sdk_samba_operate_rsp, error) {

	rsp := Sdk_samba_operate_rsp{}

	dret, err := syscall.ExecSysFunc(SDK_SAMBA_OPERATE, req)
	if err != nil {
		return rsp, err
	}
	bret, _ := json.Marshal(dret)

	err = json.Unmarshal(bret, &rsp)

	return rsp, err
}

////  downloader  /////

/*
 下载任务接口
*/

type Sdk_downloader_task_download_req struct {
	Url         string `json:"url"`
	Business    string `json:"business"`    //"pvr"
	Subbusiness string `json:"subbusiness"` //pvr_sn
	Dir         string `json:"dir"`         //路径
	Packname    string `json:"packname"`    //文件名
	Method      string `json:"method"`      //http请求方式
	Header      string `json:"header"`      //
	Body        string `json:"body"`        //
	Ext         string `json:"ext"`
}
type Sdk_downloader_task_download_rsp struct {
	/*
		-1001 - 请求解析失败
		1002  - 参数为空
		10001 - 下载失败
	*/
	Err_no  int    `json:"err_no"`
	Err_msg string `json:"err_msg"`
	Taskid  string `json:"taskid"`
}

func (r *Sdk_downloader_task_download_rsp) Result(errno int, errmsg string) {
	r.Err_no = errno
	r.Err_msg = errmsg
	//log.TRACE("SDK DownloadRsp err_no : %d , err_msg : %s", errno, errmsg)
}

func Sdk_downloader_task_download(req Sdk_downloader_task_download_req) (Sdk_downloader_task_download_rsp, error) {

	rsp := Sdk_downloader_task_download_rsp{}

	dret, err := syscall.ExecSysFunc(SDK_DOWNLOADER_DOWNLOAD, req)
	if err != nil {
		return rsp, err
	}
	bret, _ := json.Marshal(dret)

	err = json.Unmarshal(bret, &rsp)

	return rsp, err
}

/*
 更改任务状态接口
*/

type Sdk_downloader_change_status_req struct {
	Taskid string `json:"taskid"`
	Action int    `json:"action"` //
	Ext    string `json:"ext"`
}
type Sdk_downloader_change_status_rsp struct {
	/*
		-1001 - 请求解析失败
		1002  - 参数为空
		10101 - 任务未找到
	*/
	Err_no  int    `json:"err_no"`
	Err_msg string `json:"err_msg"`
	Taskid  string `json:"taskid"`
}

func Sdk_downloader_change_status(req Sdk_downloader_change_status_req) (Sdk_downloader_change_status_rsp, error) {

	rsp := Sdk_downloader_change_status_rsp{}

	dret, err := syscall.ExecSysFunc(SDK_DOWNLOADER_CHANGE_STATUS, req)
	if err != nil {
		return rsp, err
	}
	bret, _ := json.Marshal(dret)

	err = json.Unmarshal(bret, &rsp)

	return rsp, err
}

/*
 查询任务信息接口
*/

type Sdk_downloader_taskinfo_req struct {
	Taskid string `json:"taskid"`
	Ext    string `json:"ext"`
}
type Sdk_downloader_taskinfo_rsp struct {
	/*
		-1001 - 请求解析失败
		1002  - 参数为空
		10101 - 任务未找到
	*/
	Err_no     int    `json:"err_no"`
	Err_msg    string `json:"err_msg"`
	Taskid     string `json:"taskid"`
	Total      int64  `json:"total"`
	Finish     int64  `json:"finish"`
	Status     int    `json:"status"`
	Name       string `json:"name"`
	CreateTime int64  `json:"create_time"`
	ModifyTime int64  `json:"modify_time"`
}

func Sdk_downloader_taskinfo(req Sdk_downloader_taskinfo_req) (Sdk_downloader_taskinfo_rsp, error) {

	rsp := Sdk_downloader_taskinfo_rsp{}

	dret, err := syscall.ExecSysFunc(SDK_DOWNLOADER_TASK_INFO, req)
	if err != nil {
		return rsp, err
	}
	bret, _ := json.Marshal(dret)

	err = json.Unmarshal(bret, &rsp)

	return rsp, err
}

/*
 删除任务接口
*/

type Sdk_downloader_delete_req struct {
	Taskid string `json:"taskid"`
	IfDel  int    `json:"if_del"`
	Ext    string `json:"ext"`
}
type Sdk_downloader_delete_rsp struct {
	/*
		-1001 - 请求解析失败
		1002  - 参数为空
		10101 - 任务未找到
	*/
	Err_no  int    `json:"err_no"`
	Err_msg string `json:"err_msg"`
	Taskid  string `json:"taskid"`
}

func Sdk_downloader_delete(req Sdk_downloader_delete_req) (Sdk_downloader_delete_rsp, error) {

	rsp := Sdk_downloader_delete_rsp{}

	dret, err := syscall.ExecSysFunc(SDK_DOWNLOADER_DELETE, req)
	if err != nil {
		return rsp, err
	}
	bret, _ := json.Marshal(dret)

	err = json.Unmarshal(bret, &rsp)

	return rsp, err
}

/*
 查询任务列表接口
*/

type Sdk_downloader_getlist_req struct {
	Business     string `json:"business"`     //固定为"pvr"
	Sub_business string `json:"sub_business"` //头盔sn
	Ext          string `json:"ext"`          //预留
}
type Sdk_downloader_getlist_rsp struct {
	/*
		-1001 - 请求解析失败
		1002  - 参数为空
	*/
	Err_no  int        `json:"err_no"`
	Err_msg string     `json:"err_msg"`
	List    []taskInfo `json:"list"`
}

type taskInfo struct {
	Taskid     string `json:"taskid"`
	Total      int64  `json:"total"`
	Finish     int64  `json:"finish"`
	Status     int    `json:"status"`
	Name       string `json:"name"`
	CreateTime int64  `json:"create_time"`
	ModifyTime int64  `json:"modify_time"`
}

func Sdk_downloader_getlist(req Sdk_downloader_getlist_req) (Sdk_downloader_getlist_rsp, error) {

	rsp := Sdk_downloader_getlist_rsp{}

	dret, err := syscall.ExecSysFunc(SDK_DOWNLOADER_GET_LIST, req)
	if err != nil {
		return rsp, err
	}
	bret, _ := json.Marshal(dret)

	err = json.Unmarshal(bret, &rsp)

	return rsp, err
}

/*
 清空任务列表接口
*/

type Sdk_downloader_removeall_req struct {
	Business     string `json:"business"`     //固定为"pvr"
	Sub_business string `json:"sub_business"` //头盔sn
	Ext          string `json:"ext"`          //预留
}
type Sdk_downloader_removeall_rsp struct {
	/*
		-1001 - 请求解析失败
		1002  - 参数为空
	*/
	Err_no  int    `json:"err_no"`
	Err_msg string `json:"err_msg"`
}

func Sdk_downloader_removeall(req Sdk_downloader_removeall_req) (Sdk_downloader_removeall_rsp, error) {

	rsp := Sdk_downloader_removeall_rsp{}

	dret, err := syscall.ExecSysFunc(SDK_DOWNLOADER_REMOVEALL, req)
	if err != nil {
		return rsp, err
	}
	bret, _ := json.Marshal(dret)

	err = json.Unmarshal(bret, &rsp)

	return rsp, err
}

////  disk  /////

/*
 清理用户数据
*/

type Sdk_disk_clear_userdata_req struct {
	Business     string `json:"business"`     //固定为"pvr"
	Sub_business string `json:"sub_business"` //头盔sn
	Ext          string `json:"ext"`          //预留
}
type Sdk_disk_clear_userdata_rsp struct {
	/*
		-1001 - 请求解析失败
		1002  - 参数为空
	*/
	Err_no  int    `json:"err_no"`
	Err_msg string `json:"err_msg"`
}

func Sdk_disk_clear_userdata(req Sdk_disk_clear_userdata_req) (Sdk_disk_clear_userdata_rsp, error) {

	rsp := Sdk_disk_clear_userdata_rsp{}

	dret, err := syscall.ExecSysFunc(SDK_DISK_CLEAR_UERDATA, req)
	if err != nil {
		return rsp, err
	}
	bret, _ := json.Marshal(dret)

	err = json.Unmarshal(bret, &rsp)

	return rsp, err
}
