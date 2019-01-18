package common

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"strconv"
)

var Dir = []string{"/Images", "/Videos", "/Musics", "/Documents"}


//获取目录dir下的文件大小
func walkDir(dir string, wg *sync.WaitGroup, fileSizes chan<- int64) {
	defer wg.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {//目录
			wg.Add(1)
			subDir := filepath.Join(dir, entry.Name())

			go walkDir(subDir, wg, fileSizes)
		}
		fileSizes <- entry.Size()

	}
}

//sema is a counting semaphore for limiting concurrency in dirents
var sema = make(chan struct{}, 20)

//读取目录dir下的文件信息
func dirents(dir string) []os.FileInfo {
	sema <- struct{}{}
	defer func() { <-sema }()
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		return nil
	}
	return entries
}

//输出文件数量的大小
func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files %0.1fMB\n", nfiles, float64(nbytes) / 1024 /1024)
}

//提供-v 参数会显示程序进度信息
var verbose = flag.Bool("v", false, "show verbose progress messages")

func FileSize(path string) int64 {

	roots := []string{}
	roots = append(roots , path)
	fileSizes := make(chan int64)
	var wg sync.WaitGroup
	for _, root := range roots {
		wg.Add(1)
		go walkDir(root, &wg, fileSizes)
	}
	go func() {
		wg.Wait() //等待goroutine结束
		close(fileSizes)
	}()
	var nfiles, nbytes int64
loop:
	for {
		select {
		case size, ok := <-fileSizes:
			if !ok {
				break loop
			}
			nfiles++
			nbytes += size
		}
	}

	return nbytes
}

func FileExist(filename string)bool{
	if f, err := os.Stat(filename); err != nil {
		return false
	}else{
		if f.IsDir(){
			return false
		}

		return true
	}
}

func DirExist(dir string) bool {
	if _, err := os.Stat(dir); err != nil {
		return false
	}else{
		return true
	}
}

func GetFilename(name string , base string) string{
	index := 0
	for i := len(name) -1 ; i >= 0 ; i --{
		index = i
		if name[i] == '.' {
			break
		}
	}
	if index == 0 {
		index = len(name)
	}

	prefix := name[:index]

	filename := ""
	suffix , _:= strconv.Atoi(base)
	if base != "" {
		suffix += 1
		filename = fmt.Sprintf("%s(%d)%s" ,prefix , suffix , name[index:])
	} else {
		filename = name
		suffix = 1
	}

	if !DirExist(filename) {
		return filename
	}

	return GetFilename(name , strconv.Itoa(suffix))
}