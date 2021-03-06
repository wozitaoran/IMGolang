package main

import (
	"encoding/json"
	"errors"
	inet "goim/libs/net"
	"goim/libs/proto"
	"net"
	"net/rpc"

	log "github.com/thinkboy/log4go"
)

func InitRPC(auther Auther) (err error) {
	var (
		network, addr string
		c             = &RPC{auther: auther}
	)
	rpc.Register(c)
	for i := 0; i < len(Conf.RPCAddrs); i++ {
		log.Info("start listen rpc addr: \"%s\"", Conf.RPCAddrs[i])
		if network, addr, err = inet.ParseNetwork(Conf.RPCAddrs[i]); err != nil {
			log.Error("inet.ParseNetwork() error(%v)", err)
			return
		}
		go rpcListen(network, addr)
	}
	return
}

func rpcListen(network, addr string) {
	l, err := net.Listen(network, addr)
	if err != nil {
		log.Error("net.Listen(\"%s\", \"%s\") error(%v)", network, addr, err)
		panic(err)
	}
	// if process exit, then close the rpc bind
	defer func() {
		log.Info("rpc addr: \"%s\" close", addr)
		if err := l.Close(); err != nil {
			log.Error("listener.Close() error(%v)", err)
		}
	}()
	rpc.Accept(l)
}

// RPC
type RPC struct {
	auther Auther
}

func (r *RPC) Ping(arg *proto.NoArg, reply *proto.NoReply) error {
	return nil
}

// Connect auth and registe login
func (r *RPC) Connect(arg *proto.ConnArg, reply *proto.ConnReply) (err error) {
	if arg == nil {
		err = ErrConnectArgs
		log.Error("Connect() error(%v)", err)
		return
	}
	var (
		uid int64
		seq int32
	)

	// 认证结果判断 ？  //TODO
	uid, reply.RoomId, err = r.auther.Auth(arg.Token)
	if err != nil {
		if seq, err = connect(uid, arg.Server, reply.RoomId); err == nil {
			reply.Key = encode(uid, seq)

			//err 赋值
			err = errors.New("token authentication failed ...")
			return
		}
	}

	if seq, err = connect(uid, arg.Server, reply.RoomId); err == nil {
		reply.Key = encode(uid, seq)

		go checkOfflineMsg(uid, reply.Key, arg.Server) //add by xurui
		go checkOffline_GMsg(uid, reply.Key, arg.Server)
	}
	return
}

func checkOfflineMsg(uid int64, key string, serverId int32) {
	//get from db
	log.Debug("func checkOfflineMsg")
	msgs, err := getSingleOfflineMsg(uid)
	if err != nil {
		log.Error("checkOfflineMsg error:%v", err)
		return
	}
	keyArr := []string{key}

	for _, recMsg := range msgs {
		recvBodyBytes, _ := json.Marshal(recMsg)
		mpushKafka(serverId, keyArr, recvBodyBytes)
	}
}

//拉取离线群消息  并推送
func checkOffline_GMsg(uid int64, key string, serverId int32) {
	//get from db
	log.Debug("func checkOfflineMsg")
	msgs, err := getSingleOffline_GMsg(uid)
	if err != nil {
		log.Error("checkOffline_GMsg error:%v", err)
		return
	}
	keyArr := []string{key}

	for _, recMsg := range msgs {
		recvBodyBytes, _ := json.Marshal(recMsg)
		mpushKafka(serverId, keyArr, recvBodyBytes)
	}
}

// Disconnect notice router offline
func (r *RPC) Disconnect(arg *proto.DisconnArg, reply *proto.DisconnReply) (err error) {
	if arg == nil {
		err = ErrDisconnectArgs
		log.Error("Disconnect() error(%v)", err)
		return
	}
	var (
		uid int64
		seq int32
	)
	if uid, seq, err = decode(arg.Key); err != nil {
		log.Error("decode(\"%s\") error(%s)", arg.Key, err)
		return
	}
	reply.Has, err = disconnect(uid, seq, arg.RoomId)
	return
}
