package main

import (
	"errors"
	"goim/libs/define"
	"strconv"
	"strings"

	log "github.com/thinkboy/log4go"
)

// developer could implement "Auth" interface for decide how get userId, or roomId
type Auther interface {
	Auth(token string) (userId int64, roomId int32, err error)
}

type DefaultAuther struct {
}

func NewDefaultAuther() *DefaultAuther {
	return &DefaultAuther{}
}

func (a *DefaultAuther) Auth(token string) (userId int64, roomId int32, err error) {
	//TODO mysql token check
	log.Info("token=\"%s\"", token)

	var userId_temp, roomId_temp string
	var token_temp string
	userId_temp = strings.Split(token, ",")[0]
	log.Info("token[0]=\"%s\"", userId_temp)

	roomId_temp = strings.Split(token, ",")[1]
	log.Info("token[1]=\"%s\"", roomId_temp)

	token_temp = strings.Split(token, ",")[2]
	log.Info("token[1]=\"%s\"", token_temp)

	if userId, err = strconv.ParseInt(userId_temp, 10, 64); err != nil {
		userId = 0
		roomId = define.NoRoom
		return
	}

	// token认证
	token_Rerr := checktoken(token_temp)
	if token_Rerr != nil {
		log.Info("token bad")
		log.Error(token_Rerr)
		err = errors.New("token authentication failed ...")
		return
	}

	log.Info("userId=\"%d\",roomid=\"%d\"", userId, roomId)
	return
}
