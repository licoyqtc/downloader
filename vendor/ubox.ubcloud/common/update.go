package common

import (
	"fmt"
	"strconv"
	"sync"
)

var update_lock sync.Mutex

func Read_packsize(ver string) int {
	update_lock.Lock()
	defer update_lock.Unlock()

	filename := fmt.Sprintf("./packets/%s_progress",ver)
	ubyte := ReadFile(filename)

	size , _ := strconv.Atoi(string(ubyte))
	return size
}

func Write_packsize(ver string , size string) {
	update_lock.Lock()
	defer update_lock.Unlock()

	filename := fmt.Sprintf("./packets/%s_progress",ver)
	_ = WriteFile(filename , []byte(size))

	return
}

