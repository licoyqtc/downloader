package common

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"github.com/labstack/echo"
	"math/rand"
	"net/http"
	"regexp"
	"strconv"
	"time"
)

func ReadFile(filename string) string {

	bytes, err := ioutil.ReadFile(filename)
	//fmt.Printf("read byte :%s\n", bytes)
	if err != nil {
		return ""
	}

	return string(bytes)
}

func GetVersion() string {
	verdirCmd := "echo $(df |grep mmcblk0p10 |awk '{print $6}')/upper/usr/local/conf/version"
	dir, _ := Shell_cmd_single(verdirCmd)

	bytes, err := ioutil.ReadFile(dir)

	if err != nil {
		return ""
	}

	return string(bytes)
}

func GetBoxid() string {
	bytes, err := ioutil.ReadFile(F_BOX)

	if err != nil {
		return ""
	}

	return string(bytes)
}

func GetSelfChecksSatus() string {
	bytes, err := ioutil.ReadFile(F_CHECK)

	if err != nil {
		return ""
	}

	return string(bytes)
}

func GetPrivateKey() string {
	bytes, err := ioutil.ReadFile(F_PK)

	if err != nil {
		return ""
	}

	return string(bytes)
}

func GetDeviceDesc() string {
	bytes, err := ioutil.ReadFile(F_DEVICE_DESC)

	if err != nil {
		return ""
	}

	return string(bytes)
}

func WriteFile(filePath string, data []byte) error {
	err := ioutil.WriteFile(filePath, data, 0662)
	Shell_cmd_single("sync")
	return err
}

func Json_unmarshal(data interface{}, out interface{}) error {

	bdata, err := Json_marshal(data)

	if err == nil {
		err = json.Unmarshal(bdata, out)
		return err
	}

	return fmt.Errorf("unknown type")
}

func Json_marshal(data interface{}) ([]byte, error) {

	if sdata, ok := data.(string); ok {
		return []byte(sdata), nil
	}

	return json.Marshal(data)
}

func Path_validate(path *string) bool {
	if *path == "" {
		*path = "/"
		return true
	}

	if (*path)[:1] != "/" {
		*path = "/" + *path
	}
	element := strings.Split(*path, "/")

	back := 0
	dir := 0

	for _, v := range element {

		if v == ".." {
			back++
		} else if v == "." || v == "" {
			continue
		} else {
			dir++
		}
		if back > dir {
			return false
		}
	}

	return true
}

func IsWalletname(nickname *string) bool {

	namestr := *nickname
	b := 0
	e := len(namestr) - 1

	for b < len(namestr) {
		if namestr[b] != ' ' {
			break
		}
		b++
	}

	for e > 0 {
		if namestr[e] != ' ' {
			break
		}
		e--
	}

	if b > e {
		return false
	}
	*nickname = namestr[b : e+1]

	pat := "[\u4E00-\u9FA5]"

	re, _ := regexp.Compile(pat)
	str := re.ReplaceAllString(namestr[b:e+1], "yq")

	if len(str) > 60 {
		return false
	}
	fmt.Printf("len nick : %s %d \n", str, len(str))
	return true
}

func GetRand() int {
	timestamp := time.Now().UnixNano()
	r := rand.New(rand.NewSource(timestamp))
	return r.Int()
}

func GetRandStr() (token string) {
	tempRand := GetRand()
	strRand := strconv.Itoa(tempRand)
	//计算随机数的md5值，做16进制字节码转换
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(strRand))
	randStr := hex.EncodeToString(md5Ctx.Sum(nil))

	return randStr[:16]
}

func Get_nowtime_formatted() string {
	timestamp := time.Now().Unix()
	return time.Unix(timestamp, 0).Format("2006-01-02 15:04:05") //go 语言 固定的时间格式

}

func Get_nowtime_date() string {
	timestamp := time.Now().Unix()
	return time.Unix(timestamp, 0).Format("20060102_150405") //go 语言 固定的时间格式
}

func Get_nowtime_delay(delay int64) string {
	timestamp := time.Now().Unix()
	timestamp += delay
	return time.Unix(timestamp, 0).Format("2006-01-02 15:04:05") //go 语言 固定的时间格式
}

func Get_unix_bystring(str string) (int64, error) {
	loc, _ := time.LoadLocation("Local") //重要：获取时区
	timeNow, err := time.ParseInLocation("2006-01-02 15:04:05", str, loc)

	if err != nil {
		return 0, err
	}
	return timeNow.Unix(), nil
}

func MD5(data string) string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(data))
	encode := hex.EncodeToString(md5Ctx.Sum(nil))
	return encode
}

func SHA256(data string) string {
	sha256Ctx := sha256.New()
	sha256Ctx.Write([]byte(data))
	encode := hex.EncodeToString(sha256Ctx.Sum(nil))
	return encode
}

func IsEmailaddress(email string) bool {

	ret, _ := regexp.MatchString("^[a-zA-Z0-9_.-]+@[a-zA-Z0-9_-]+(\\.[a-zA-Z0-9_-]+)+$", email)
	return ret
}

func IsEthPubaddr(addr string) bool {
	ret, _ := regexp.MatchString("^(0x)?[a-f0-9A-F]{40}$", addr)
	return ret
}

func GenToken(version string) (token string) {
	token = version
	//获取时间
	strtime := strconv.FormatInt(time.Now().Unix(), 16)
	token += strtime
	timestamp := time.Now().UnixNano()
	//获取随机数
	r := rand.New(rand.NewSource(timestamp))
	tempRand := r.Int()
	strRand := strconv.Itoa(tempRand)
	//计算随机数的md5值，做16进制字节码转换
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(strRand))
	token += hex.EncodeToString(md5Ctx.Sum(nil))
	md5Ctx.Reset()
	macData := token
	md5Ctx.Write([]byte(macData))
	hash := hex.EncodeToString(md5Ctx.Sum(nil))
	token += hash[0:3]
	return token
}

func SetCookies(e echo.Context, key string, value string, ttl int) {
	cookie := http.Cookie{
		Name:   key,
		Value:  value,
		Path:   "/",
		MaxAge: ttl,
	}
	e.SetCookie(&cookie)
}

func FormatAddress(addr string) string {
	if addr == "" {
		return ""
	}

	if addr[:2] != "0x" {
		return "0x" + strings.ToLower(addr)
	}
	return strings.ToLower(addr)
}

func ConvertData(src interface{}, p_dest interface{}) (e error) {
	b_src, ok := src.([]byte)

	if !ok {
		b_src, e = json.Marshal(src)
	}

	e = json.Unmarshal(b_src, p_dest)
	return
}
