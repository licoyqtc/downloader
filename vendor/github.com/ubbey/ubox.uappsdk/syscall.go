/***********************************************************************
// Copyright yqtc.com @2017 The source code.
// Copyright (c) 2009-2016 The Bitcoin Core developers
// Distributed under the MIT software license, see the accompanying
// file COPYING or http://www.opensource.org/licenses/mit-license.php.
//******
// Filename:
// Description:
// Author:
// CreateTime:
/***********************************************************************/
package ubox_uappsdk

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
)

const SYSCALL_UNIX_SOCK = "/var/run/sys.sock"

type execargs struct {
	Cmd  string      `json:"cmd"`
	Args interface{} `json:"args"`
}

type execrets struct {
	Err  string      `json:"err"`
	Rets interface{} `json:"rets"`
}

func ExecSysFunc(cmd string, args ...interface{}) (ret interface{}, err error) {
	var (
		errstr = ""
		data   = []byte{}
		n      = 0
		buf    = make([]byte, 8192)
		rets   = execrets{}
	)

	// Step 1. dail unix domain
	conn, err := net.Dial("unix", SYSCALL_UNIX_SOCK)
	if err != nil {
		errstr = fmt.Sprint("ExecSysFunc Dial error:%s", err.Error())
		goto error
	}
	defer conn.Close()

	// Step 2. send syscall data
	data, err = json.Marshal(execargs{cmd, args})
	if err != nil {
		errstr = fmt.Sprint("ExecSysFunc json.Marshal error:%s", err.Error())
		goto error
	}
	n, err = conn.Write(data)
	if err != nil {
		errstr = fmt.Sprint("ExecSysFun conn.Write error:%s", err.Error())
		goto error
	} else if n != len(data) { /* maybe a bug*/
		errstr = fmt.Sprint("ExecSysFun must write data one time")
		goto error
	}

	// Step 3. wait return
	n, err = conn.Read(buf[:])
	if err != nil {
		errstr = fmt.Sprint("ExecSysFun conn.Read error:%s", err.Error())
		goto error
	}
	err = json.Unmarshal(buf[0:n], &rets)
	if err != nil {
		errstr = fmt.Sprint("ExecSysFun json.Unmarshal error:%s", err.Error())
		goto error
	}
	if len(rets.Err) != 0 {
		log.Printf("reject :%+v\n", rets)
		return rets.Rets, err
	} else {
		log.Printf("resolve :%+v\n", rets)
		return rets.Rets, nil
	}

error:
	log.Fatalf("%s\n", errstr)
	return rets.Rets, err
}
