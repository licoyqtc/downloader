package common

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type pocFile struct {
	Name     string
	Coinbase string
	Start    int
	Quantity int
}

type PocFiles []pocFile

func (pfs PocFiles) Len() int {
	return len(pfs)
}
func (pfs PocFiles) Swap(i, j int) { // 重写 Swap() 方法
	pfs[i], pfs[j] = pfs[j], pfs[i]
}
func (pfs PocFiles) Less(i, j int) bool { // 重写 Less() 方法， 从小到大排序
	return pfs[j].Start > pfs[i].Start
}

func TrimPocFiles(files []string) (pfs PocFiles, e error) {

	tmpPfs := PocFiles{}
	for _, v := range files {
		//过滤还未生成完成的任务
		if strings.HasSuffix(v, ".orig") || strings.HasSuffix(v, ".dest") {
			fmt.Println("found not complete poc file name : " + v)
			continue
		}
		//过滤不符合规则的文件
		elements := strings.Split(v, "_")
		if len(elements) != 3 {
			fmt.Println("found invaild poc file name : " + v)
			continue
		}

		pf := pocFile{}
		pf.Name = v
		pf.Coinbase = elements[0]
		pf.Start, _ = strconv.Atoi(elements[1])
		pf.Quantity, _ = strconv.Atoi(elements[2])
		if pf.Quantity != UBBEY_NONCE_QUANTITY {
			fmt.Println("found invaild quantity number file name : " + v)
			continue
		}

		tmpPfs = append(tmpPfs, pf)

	}
	sort.Sort(tmpPfs)
	if len(tmpPfs) == 0 {
		return
	}

	start := tmpPfs[0].Start
	for _, v := range tmpPfs {
		if v.Coinbase == "0000000000000000000000000000000000000000" {
			continue
		}
		if v.Start != start {
			errstr := fmt.Sprintf("trim poc file err , start pos not match : %d %d", v.Start, start)
			fmt.Println(errstr)
			e = fmt.Errorf(errstr)
			return
		}

		start += v.Quantity
		pfs = append(pfs, v)
	}

	return
}
