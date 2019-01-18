/*
	desc: 断点下载器
	author: licowei
	time: 2018-12-23
	ext:
*/

package downloader

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/ubbey/ubox.downloader/model"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

const (
	default_size   = 1024 * 128 // 128kb
	default_method = "POST"
)

func download(task *model.DownloadTask, args *http_arg, mode int) (n int64, pack_size int64, e error) {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
			switch r := r.(type) {
			case error:
				e = r
			default:
				e = fmt.Errorf("%v", r)
			}
		}
	}()

	if task == nil {
		return 0, 0, fmt.Errorf("download fail , task null")
	}

	if args == nil {
		args = &http_arg{}
		args.Header["Content-type"] = []string{"application/json"}
		args.Off_size = default_size
		args.Method = default_method
	}

	if args.Header["Content-type"] == nil {
		args.Header["Content-type"] = []string{"application/json"}
	}
	args.Header["Range"] = []string{fmt.Sprintf("bytes=%d-%d", task.CurOffset, task.CurOffset+args.Off_size-1)}

	url := task.Url
	method := args.Method

	// Stpe 1. new request parameter
	argsjson, err := json.Marshal(args.Body)
	fmt.Println("call http json, url:[", url, "], argsjson: ", string(argsjson))
	req, err := http.NewRequest(method, url, bytes.NewBuffer(argsjson))
	if err != nil {
		panic("http req err ：" + err.Error())
	}

	for k, v := range args.Header {
		for _, v2 := range v {
			req.Header.Add(k, v2)
		}
	}

	// Stpe 2. execute request
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	resp, err := client.Do(req)
	if err != nil {
		panic("Error in sending request to " + url + " " + err.Error())
	}
	defer resp.Body.Close()

	filepath := strings.Join([]string{task.Dir, task.Name}, "/")

	fmt.Println("file download : ", filepath, " , range : ", resp.Header["Content-Range"])
	dst, err := os.OpenFile(filepath, mode, 0664)
	if err != nil {
		panic("open file err : " + err.Error())
	}
	defer dst.Close()

	if len(resp.Header["Content-Range"]) == 1 {
		size_aray := strings.Split(resp.Header["Content-Range"][0], "/")
		s, _ := strconv.Atoi(size_aray[1])
		pack_size = int64(s)

	} else {
		panic(fmt.Sprintf("download pack err , invaild header %v", resp.Header["Content-Range"]))
	}

	// Copy
	if n, err = io.Copy(dst, resp.Body); err != nil {
		panic("copy data err : " + err.Error())
	}

	return n, pack_size, err
}
