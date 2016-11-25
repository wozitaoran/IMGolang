package main

import (
	"database/sql"
	"goim/libs/proto"

	log "github.com/thinkboy/log4go"
	//"time"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func initDB() {
	var err error
	log.Info("mysql conn pool init")

	//获取连接池 sql 包的 Close 方法只有 3个，除了 *sql.Db 是连接池对象，使用中是不会关闭的 其他的两个 Rows.Close 和 Stmt.Close 是需要关的
	db, err = sql.Open("mysql", Conf.DBAddrs)
	log.Info("Conf.DBAddrs:\"%s\"", Conf.DBAddrs)
	db.SetMaxIdleConns(20)
	db.SetMaxOpenConns(20)

	checkErr(err)
}

func addMsgRecord(sendId int64, target_type string, targetId int64, msgContent string, msg_type int64) {
	//INSERT im_msg_send SET sendId=?,target_type=?,targetId=?,msg=?,msg_type=?,send_time=NOW()
	stmt, err := db.Prepare("INSERT im_msg_send SET sendId=?,target_type=?,targetId=?,msg=?,msg_type=?,send_time=NOW()")
	checkErr(err)

	_, err = stmt.Exec(sendId, target_type, targetId, msgContent, msg_type)
	checkErr(err)

	stmt.Close()
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

	//删除数据
	stmt, err := db.Prepare("delete from im_s_offline_msg where recvId=?")

	res, err := stmt.Exec(uid)

	affect, err := res.RowsAffected()
	stmt.Close()
	log.Debug("delete from im_s_offline_msg==========")
	log.Debug(affect)

	return
}

//拉取离线群消息
func getSingleOffline_GMsg(uid int64) (gmsgs []proto.Recv_GMessage, err error) {
	db, err := sql.Open("mysql", "root:yxkj@tcp(192.168.19.37:3307)/lexiangccb?charset=utf8")
	if err != nil {
		err.Error()
	}

	rows, err := db.Query("SELECT SendId,groupid,MsgContent,Msg_type,Insert_time FROM im_g_offline_msg WHERE recvId=?", uid)

	for rows.Next() {
		var SendId int64
		var groupid int64

		var MsgContent string
		var Msg_type int64
		var Insert_time string
		err = rows.Scan(&SendId, &groupid, &MsgContent, &Msg_type, &Insert_time)
		msg := proto.Recv_GMessage{"group", Msg_type, groupid, MsgContent, SendId, Insert_time}
		gmsgs = append(gmsgs, msg)
	}

	rows.Close()

	//删除数据
	// stmt, err := db.Prepare("delete from im_g_offline_msg where recvId=?")

	// res, err := stmt.Exec(uid)

	// _, err = res.RowsAffected()
	// stmt.Close()

	return
}

//根据群id查找群成员id
func getGroup_membertoUser_id(gid int64) (uids []int64, err error) {
	rows, err := db.Query("select user_id from group_member where group_id = ?;", gid)

	var uid int64
	for rows.Next() {
		err = rows.Scan(&uid)
		uids = append(uids, uid)
	}
	rows.Close()
	return uids, nil
}

//插入群离线消息
func addSingleOffline_groupmsg(sendId int64, recvId int64, groupid int64, msgContent string, msg_type int64) {
	stmt, err := db.Prepare("insert into im_g_offline_msg (sendId,recvId,groupid,recvTime,msgContent,msg_type,insert_time)values(?,?,?,'',?,?,now());")
	if err != nil {
		err.Error()
	}

	_, err = stmt.Exec(sendId, recvId, groupid, msgContent, msg_type)
	if err != nil {
		err.Error()
	}
	stmt.Close()
}

//根据id读取token
func get_IDtoToken(user_id string) (salt string) {
	rows, err := db.Query("select salt from token where user_id = ?;", user_id)
	if err != nil {
		err.Error()
	}

	for rows.Next() {
		err = rows.Scan(&salt)
	}
	rows.Close()

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
