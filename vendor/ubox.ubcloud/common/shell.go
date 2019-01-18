package common

import (
	"os/exec"
	"strings"
)



const (
	CMD_FTP_COUNT = "ps -ef | grep vsftpd | grep -v grep | wc -l"
	CMD_SAMB_COUNT = "ps -ef | grep smbd | grep -v grep | wc -l"

	CMD_FTP_ENABLE = "service vsftpd start"
	CMD_SAMBA_ENABLE = "smbd"

	CMD_FTP_DISABLE = "service vsftpd stop"
	CMD_SAMBA_DISABLE = "smbcontrol smbd shutdown"

	CMD_CPU_USAGE = "vmstat"
)


func Shell_cmd_single(cmd string) (string , error) {
	rs := exec.Command("/bin/bash" , "-c" , cmd)
	out, err := rs.Output()
	if err != nil {
		return "" , err
	}
	ret := string(out)
	trimBlank := strings.Trim(ret, " ")
	rsp := strings.Trim(trimBlank, "\n")

	return rsp , nil
}

func Shell_cmd_notwait(cmd string) ( error) {
	rs := exec.Command("/bin/bash" , "-c" , cmd)
	err  := rs.Start()
	if err != nil {
		return  err
	}

	return nil
}

func Shell_cmd_multi(cmd string) ([]string , error) {
	rs := exec.Command("/bin/bash" , "-c" , cmd)
	out, err := rs.Output()
	if err != nil {
		return nil , err
	}
	ret := string(out)
	rsp := strings.Split(ret, "\n")


	return rsp , nil
}

