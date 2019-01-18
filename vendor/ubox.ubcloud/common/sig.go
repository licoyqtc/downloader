package common

import (
	"encoding/base64"
	"ubox.golib/crypt"
	"fmt"
)

type Signature struct {
	Username 	string		`json:"username"`
	Boxid		string		`json:"boxid"`
	Timestamp	int64		`json:"timestamp"`
}


func GetSignature(sig string , key string) (st Signature ,e error){
	defer func (){
		err := recover()
		if err != nil {
			e = fmt.Errorf("err :%s",err)
		}
	}()

	rsp := Signature{}
	b , err := base64.StdEncoding.DecodeString(sig)
	if err != nil {
		return rsp , err
	}

	destr , err := crypt.TripleDesDecrypt(b , []byte(key))
	if err != nil {
		return rsp , err
	}

	err = Json_unmarshal(string(destr) , &rsp)

	return rsp , e
}
