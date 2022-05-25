package im

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/wuyan94zl/IM/app/internal/svc"
	"github.com/wuyan94zl/IM/app/models/hasusers"
	"github.com/wuyan94zl/IM/app/models/messages"
	"github.com/wuyan94zl/IM/app/models/sendqueue"
	"github.com/wuyan94zl/chat"
	"net/http"
	"strconv"
	"time"
)

const (
	sendMessage = 100
)

type cliDetail struct {
	NickName  string `json:"nick_name"`
	Phone     string `json:"phone"`
	HeadProto string `json:"head_proto"`
}

func RunWs(w http.ResponseWriter, r *http.Request, svcCtx *svc.ServiceContext) {
	token, _ := strconv.Atoi(r.FormValue("_token"))
	info, err := svcCtx.UserModel.FindOne(context.Background(), int64(token))
	if err != nil {
		fmt.Println("ws 连接参数错误：", err)
	} else {
		chart.NewServer(w, r, uint64(info.Id), cliDetail{NickName: info.NickName, Phone: info.Mobile, HeadProto: "test uri"}, &data{ctx: svcCtx})
	}
}

type Server struct {
	Ctx *svc.ServiceContext
}

func (i Server) Start() {
	Run(i.Ctx)
}
func (i Server) Stop() {
	fmt.Println("im service was stop...")
}

func Run(ctx *svc.ServiceContext) {
	mux := http.NewServeMux()
	mux.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		token, _ := strconv.Atoi(r.FormValue("_token"))
		info, err := ctx.UserModel.FindOne(context.Background(), int64(token))
		if err != nil {
			fmt.Println("ws 连接错误：", err)
			return
		}
		chart.NewServer(w, r, uint64(info.Id), cliDetail{NickName: info.NickName, Phone: info.Mobile, HeadProto: "test uri"}, &data{ctx: ctx})
	})
	fmt.Printf("Starting Ws Server at 0.0.0.0:9988...\n")
	err := http.ListenAndServe(":9988", mux)
	if err != nil {
		fmt.Println("ws err：", err)
		return
	}
}

type data struct {
	ctx *svc.ServiceContext
}

// 发送消息回调
func (d *data) SendMessage(msg chart.Message) {
	switch msg.Type {
	case sendMessage:
		sendTime, err := time.Parse("2006-01-02 15:01:05", "2022-04-15 22:12:12")
		if err != nil {
			return
		}
		message := messages.Messages{
			ChannelId:  msg.ChannelId,
			SendUserId: int64(msg.UserId),
			Message:    msg.Content,
			CreateTime: sendTime,
		}
		d.ctx.MessageModel.Insert(context.Background(), &message)
	}
}

func (d *data) DelaySendMessage(channelId string, msg chart.Message, sent []uint64) {
	fmt.Println("delay：", channelId, msg, sent)
	var ids []int64
	switch msg.Type {
	case 101:
		users, _ := d.ctx.GroupUserModel.FindUsersByChannelId(channelId)
		for _, u := range users {
			ids = append(ids, u.UserId)
		}
	case 100:
		var users []hasusers.UserUsers
		if channelId == "" {
			ids = append(ids, int64(msg.ToUserId))
		} else {
			if len(sent) == 2 { // 单聊发送人数为2，则无离线消息 return
				return
			}
			users, _ = d.ctx.UserUsersModel.AllChannelIdUsers(channelId)
			for _, u := range users {
				ids = append(ids, u.UserId)
			}
		}
	}

	sentMap := make(map[int64]bool)
	for _, v := range sent {
		sentMap[int64(v)] = true
	}
	msgByte, err := json.Marshal(msg)
	if err != nil {
		return
	}
	for _, uid := range ids {
		if _, ok := sentMap[uid]; !ok {
			d.ctx.SendQueueModel.Insert(context.Background(), &sendqueue.SendQueues{UserId: uid, Message: string(msgByte), SendUserId: int64(msg.UserId)})
		}
	}
}

// LoginServer 登录成功后回调
func (d *data) LoginServer(uid uint64) {
	//list, _ := d.ctx.UserModel.Friends(d.ctx.UserUsersModel, int64(uid))
	list, _ := d.ctx.UserUsersModel.Friends(int64(uid))
	var channelIds []string
	for _, v := range list {
		channelIds = append(channelIds, GenChannelIdByFriend(int64(uid), v.HasUserId))
	}
	groups, _ := d.ctx.GroupUserModel.InGroups(context.Background(), int64(uid))
	for _, v := range groups {
		channelIds = append(channelIds, v.ChannelId)
	}
	chart.JoinChannelIds(uid, channelIds...)
	go func() {
		time.Sleep(time.Second * 1)
		queues, _ := d.ctx.SendQueueModel.FindByUserId(context.Background(), int64(uid))
		for _, queue := range queues {
			SendMessageToUid(uid, uid, queue.Message, 100)
			d.ctx.SendQueueModel.Delete(context.Background(), queue.Id)
		}
	}()
}
func (d *data) LogoutServer(uid uint64) {
	// 退出登陆回调
	fmt.Println("logout ", uid)
}
func (d *data) ErrorLogServer(err error) {
	// 错误消息回调
	fmt.Println("err: ", err)
}
