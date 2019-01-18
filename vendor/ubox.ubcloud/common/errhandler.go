package common

import "fmt"

func Args_null_err() string {
	return "arguments can't be null"
}

func Args_invaild() string {
	return "arguments invaild"
}
func Args_bind_err(err string) string {
	return fmt.Sprintf("args bind err : %s", err)
}

func Get_cookies_err(err string) string {
	return fmt.Sprintf("Get cookies err : %s", err)
}

func Get_binduser_err(err string) string {
	return fmt.Sprintf("Get bind user err : %s", err)
}

func Get_session_err(err string) string {
	return fmt.Sprintf("Get session err : %s", err)
}
