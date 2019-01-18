package common

import (
	"fmt"
	"strconv"
	"strings"
)

type diskInfo struct {
	Device_type   string
	Disk_type     string
	Device_series string
	Device_name   string
	Use           int
	Available     int
	Total         int
	Filesystem    string
	Mount         string
}

func Get_bindEncodeUser(username string) (encUser string) {
	emailSp := strings.Split(username, "@")
	if len(emailSp) == 2 {
		len := len(emailSp[0])
		index := 0

		prefix := ""
		for i := 0; i < len; i++ {
			if index < 2 {
				prefix += emailSp[0][index : index+1]
			} else if index < 4 {
				prefix += "*"
			}

			index++
		}

		for ; index < 4; index++ {
			prefix += "*"
		}

		if len >= 3 {
			prefix += emailSp[0][len-1 : len]
		} else {
			prefix += "*"
		}

		encUser = fmt.Sprintf("%s@%s", prefix, emailSp[1])

	}

	return
}

func Parse_cpuUsage() (use string) {
	cmd := "vmstat"

	results, _ := Shell_cmd_multi(cmd)
	if len(results) < 3 {
		fmt.Printf("cpu parse len err :%d", len(results))
		return "0%"
	}

	header := results[1]
	headers := strings.Split(header, " ")
	index := 0
	for _, v := range headers {
		if v == "id" {
			index++
			break
		}
		if v != "" {
			index++
		}
	}

	content := results[2]
	contents := strings.Split(content, " ")

	for _, v := range contents {
		if v != "" {
			index--
		}
		if index == 0 {
			free, _ := strconv.Atoi(v)
			return fmt.Sprintf("%d", 100-free) + "%"
		}
	}

	fmt.Printf("index :%d len headers :%d len contents :%d\n", index, len(headers), len(contents))
	return "0%"
}

func Parse_memUsage() (total int, free int) {
	cmd := "free"

	results, _ := Shell_cmd_multi(cmd)
	if len(results) < 2 {
		fmt.Printf("mem parse len err :%d", len(results))
		return 0, 0
	}

	header := results[0]
	headers := strings.Split(header, " ")
	index := 1
	totalId := 0
	freeId := 0
	bufId := 0
	cacheId := 0
	for _, v := range headers {
		if v != "" {
			index++
		}
		switch {
		case v == "total":
			totalId = index
		case v == "free":
			freeId = index
		case v == "buffers":
			bufId = index
		case v == "cached":
			cacheId = index
		}
	}

	content := results[1]
	contents := strings.Split(content, " ")
	fmt.Printf("total:%d free:%d buffers:%d cached:%d\n", totalId, freeId, bufId, cacheId)
	cId := 0
	total = 0
	free = 0
	for _, v := range contents {
		if v == "" {
			continue
		}
		cId++
		switch {
		case cId == totalId:
			total, _ = strconv.Atoi(v)
			fmt.Printf("total:%s\n", v)
		case cId == freeId:
			freeSize, _ := strconv.Atoi(v)
			free += freeSize
			fmt.Printf("free:%s\n", v)
		case cId == bufId:
			bufSize, _ := strconv.Atoi(v)
			free += bufSize
			fmt.Printf("buffer:%s\n", v)
		case cId == cacheId:
			cacheSize, _ := strconv.Atoi(v)
			free += cacheSize
			fmt.Printf("cached:%s\n", v)
		}
	}

	fmt.Printf("total :%d free :%d\n", total, free)

	return total, free
}

func Parse_temp() string {
	cmd := "cat /sys/class/thermal/thermal_zone0/temp"

	result, _ := Shell_cmd_single(cmd)

	tmp, _ := strconv.Atoi(result)

	return fmt.Sprintf("%0.1f°C", float64(tmp)/1000)

}

func Parse_disk() []diskInfo {
	cmd := "df"

	results, _ := Shell_cmd_multi(cmd)
	header := results[0]
	headers := strings.Split(header, " ")
	index := 0
	useId := 0
	availableId := 0
	mountId := 0
	filesysId := 0
	totalId := 0

	for _, v := range headers {
		if v != "" {
			index++
		}
		switch v {
		case "Used":
			useId = index
		case "Available":
			availableId = index
		case "Mounted":
			mountId = index
		case "Filesystem":
			filesysId = index
		case "1K-blocks":
			totalId = index
		}
	}

	deviceType := []string{"sda", "sdb", "hda", "hdb", "vda", "vdb"}
	ret := []diskInfo{}
	for _, dvType := range deviceType {
		diskname := fmt.Sprintf("/dev/%s", dvType)

		for _, v := range results {
			if strings.Contains(v, diskname) {
				disk := diskInfo{}
				diskSpilt := strings.Split(v, " ")
				di := 0
				for _, f := range diskSpilt {
					if f == "" {
						continue
					}
					di++
					switch di {
					case useId:
						useSize, _ := strconv.Atoi(f)
						disk.Use = useSize
					case availableId:
						avaSize, _ := strconv.Atoi(f)
						disk.Available = avaSize
					case mountId:
						disk.Mount = f
					case filesysId:
						disk.Filesystem = f
					case totalId:
						totalSize, _ := strconv.Atoi(f)
						disk.Total = totalSize
					}

				}
				disk.Disk_type = dvType
				disk.implement_info()
				ret = append(ret, disk)
			}

		}
	}

	return ret
}

func Enable_ftp() error {
	shell_enable := CMD_FTP_ENABLE
	_, err := Shell_cmd_single(shell_enable)
	if err != nil {
		return err
	}

	shell_ps := CMD_FTP_COUNT

	rs, err := Shell_cmd_single(shell_ps)
	if err != nil {
		return err
	}

	count, _ := strconv.Atoi(rs)
	if count == 0 {
		return fmt.Errorf("system failed : start ftp service failed")
	}

	return nil
}

func Disable_ftp() error {
	shell_enable := CMD_FTP_DISABLE
	_, err := Shell_cmd_single(shell_enable)
	if err != nil {
		return err
	}

	shell_ps := CMD_FTP_COUNT

	rs, err := Shell_cmd_single(shell_ps)
	if err != nil {
		return err
	}

	count, _ := strconv.Atoi(rs)
	if count != 0 {
		return fmt.Errorf("system failed : stop ftp service failed")
	}
	return nil
}

func Enable_samba() error {
	shell_enable := CMD_SAMBA_ENABLE
	_, err := Shell_cmd_single(shell_enable)
	if err != nil {
		return err
	}

	shell_ps := CMD_SAMB_COUNT

	rs, err := Shell_cmd_single(shell_ps)
	if err != nil {
		return err
	}

	count, _ := strconv.Atoi(rs)
	if count == 0 {
		return fmt.Errorf("system failed : start samba service failed")

	}
	return nil
}

func Disable_samba() error {
	shell_disable := CMD_SAMBA_DISABLE
	_, err := Shell_cmd_single(shell_disable)
	if err != nil {
		return err
	}

	shell_ps := CMD_SAMB_COUNT

	rs, err := Shell_cmd_single(shell_ps)
	if err != nil {
		return err
	}

	count, _ := strconv.Atoi(rs)
	if count != 0 {
		return fmt.Errorf("system failed : stop samba service failed")
	}

	return nil
}

func (ds *diskInfo) implement_info() {
	queryDevice := "hdparm -I " + ds.Filesystem

	results, _ := Shell_cmd_multi(queryDevice)
	for _, v := range results {
		if strings.Contains(v, "Model Number:") {
			diskInfos := strings.Split(v, "Model Number:")
			start := 0
			end := 0
			opstr := diskInfos[1]

			for i, c := range opstr {
				if c != ' ' {
					start = i
					break
				}
			}

			index := len(opstr) - 1
			for index > 0 {
				if opstr[index] != ' ' {
					end = index
					break
				}
				index--
			}
			fmt.Printf("Model start : %d , end : %d\n", start, end)
			ds.Device_name = opstr[start : end+1]
		} else if strings.Contains(v, "Serial Number:") {
			diskInfos := strings.Split(v, "Serial Number:")
			start := 0
			end := 0
			opstr := diskInfos[1]

			for i, c := range opstr {
				if c != ' ' {
					start = i
					break
				}
			}
			index := len(opstr) - 1
			for index > 0 {
				if opstr[index] != ' ' {
					end = index
					break
				}
				index--
			}
			fmt.Printf("Serial start : %d , end : %d\n", start, end)
			ds.Device_series = opstr[start : end+1]
		}
		queryDeviceRotation := fmt.Sprintf("cat /sys/block/%s/queue/rotational", ds.Disk_type)
		rotation, _ := Shell_cmd_single(queryDeviceRotation)

		if rotation == "1" {
			ds.Device_type = "HHD"
		} else {
			ds.Device_type = "SSD"
		}
	}
}
