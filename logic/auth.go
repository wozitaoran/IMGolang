package main

import (
	log "github.com/thinkboy/log4go"
	"goim/libs/define"
	"strconv"
	"strings"
)

// developer could implement "Auth" interface for decide how get userId, or roomId
type Auther interface {
	Auth(token string) (userId int64, roomId int32)
}

type DefaultAuther struct {
}

func NewDefaultAuther() *DefaultAuther {
	return &DefaultAuther{}
}

func (a *DefaultAuther) Auth(token string) (userId int64, roomId int32) {
	var err error
	log.Info("token=\"%s\"", token)
	var token0, token1 string
	token0 = strings.Split(token, ",")[0]
	log.Info("token[0]=\"%s\"", token0)
	token1 = strings.Split(token, ",")[1]
	log.Info("token[1]=\"%s\"", token1)
	if userId, err = strconv.ParseInt(token0, 10, 64); err != nil {
		userId = 0
	}
	var roomIdTemp int64
	if roomIdTemp, err = strconv.ParseInt(token1, 10, 64); err != nil {
		roomId = define.NoRoom
	} else {
		roomId = int32(roomIdTemp)
	}
	log.Info("userId=\"%d\",roomid=\"%d\"", userId, roomId)
	return
}
