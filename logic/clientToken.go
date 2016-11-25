package main

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"
)

func checktoken(xAuthToken string) (err error) {
	//认证token相关：
	//从http header中取X-Auth-Token值。
	//X-Auth-Token结构：userId:expire:hash
	//expire:过期时间戳，如果过期，则不必进行后续步骤。
	//hash计算方式：
	//1、根据userId从token表中取salt值。
	//2、使用salt对userId:expire进行4次迭代的sha-1。
	//如果最终计算出的hash值一致，则验证通过。

	//返回现在时间戳
	var (
		userId string
		expire string
		hash   string
	)

	tokenSplit := strings.Split(xAuthToken, ":")
	if len(tokenSplit) != 3 {
		err = errors.New("token err")
		return err
	}

	userId = tokenSplit[0]
	expire = tokenSplit[1]
	hash = tokenSplit[2]

	//1、判断是否过期
	e, _ := strconv.ParseInt(expire, 10, 64)
	if e < time.Now().Unix() {
		//过期  TODO 是否断开连接？
		err = errors.New("token timeout")
		return err
	}

	//2、hash计算
	//hash计算方式：

	//"token": "1:1477671425846:fcb57f9107bce14e67c57a7f190ab6e4f2876031",
	//salt  9721c3be4ecb55cf
	//TODO 查询salt
	//读取数据库
	//salt := "9721c3be4ecb55cf"
	salt := get_IDtoToken(userId)
	//hashResult := fsha1(userId + ":" + expire + ":" + salt)
	hashResult := nsha1(userId+":"+expire+":"+salt, 1)

	if hashResult == hash {
		return nil
	} else {
		err = errors.New("token failed")
		return err
	}

}

//n time sha1
func nsha1(data string, time int) string {
	for index := 0; index < time; index++ {
		data = fsha1(data)
		fmt.Println(data)
	}
	return data
}

//single time sha1
func fsha1(data string) string {
	t := sha1.New()
	io.WriteString(t, data)
	return fmt.Sprintf("%x", t.Sum(nil))
}

// func main() {
// 	checktoken("1:1482475914:86700932c629920a8dbc41c10aa3c954d828d22474c4afc249a402becade7fce")
// }
