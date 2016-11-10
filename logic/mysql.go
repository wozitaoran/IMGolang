package main

import (
	"database/sql"
	"goim/libs/proto"

	log "github.com/thinkboy/log4go"
	//"time"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	var err error
	log.Info("mysql conn pool init")

	//获取连接池 sql 包的 Close 方法只有 3个，除了 *sql.Db 是连接池对象，使用中是不会关闭的 其他的两个 Rows.Close 和 Stmt.Close 是需要关的
	db, err = sql.Open("mysql", "root:yxkj@tcp(192.168.19.37:3307)/lexiangccb?charset=utf8")
	db.SetMaxIdleConns(20)
	db.SetMaxOpenConns(20)

	checkErr(err)
}

//insert offline single msg  (fromId, targetId, body, msg_type)
func addSingleOfflinemsg(sendId int64, recvId int64, msgContent string, msg_type int64) {
	//INSERT im_s_offline_msg SET id=?,sendId=?,recvId=?,msgContent=?,msg_type=?,insert_time=NOW(),note=?
	//INSERT im_s_offline_msg SET sendId='2',recvId='1',msgContent='{test:1}',msg_type='1',insert_time=NOW(),note='test'
	stmt, err := db.Prepare("INSERT im_s_offline_msg SET sendId=?,recvId=?,msgContent=?,msg_type=?,insert_time=NOW(),note=?")
	checkErr(err)

	note := "test"
	_, err = stmt.Exec(sendId, recvId, msgContent, msg_type, note)
	checkErr(err)

	stmt.Close()
}

func getSingleOfflineMsg(uid int64) (msgs []proto.RecvMessage, err error) {

	//SELECT sendId,msgContent,msg_type,insert_time FROM im_user WHERE recvId =
	rows, err := db.Query("SELECT sendId,msgContent,msg_type,insert_time FROM im_s_offline_msg WHERE recvId=?", uid)

	for rows.Next() {
		var SendId int64
		var MsgContent string
		var Msg_type int64
		var Insert_time string
		err = rows.Scan(&SendId, &MsgContent, &Msg_type, &Insert_time)
		msg := proto.RecvMessage{"user", Msg_type, MsgContent, SendId, Insert_time}
		msgs = append(msgs, msg)
	}
	log.Debug("getSingleOfflineMsg==========")
	log.Debug(msgs)
	rows.Close()

	//delete

	return
}

//chang status online
//func userConnectMysql() {

//}

//change status offline
//func userDisconnectMysql() {

//}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
