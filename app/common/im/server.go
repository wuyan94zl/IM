package im

import (
	"context"
	"fmt"
	"github.com/wuyan94zl/chart"
	"github.com/wuyan94zl/go-zero-blog/app/internal/svc"
	"github.com/wuyan94zl/go-zero-blog/app/models/messages"
	"net/http"
	"strconv"
	"time"
)

const (
	AesKey         = "wuyan94zl1asdfghjklqwertyuiopzas"
	publicChanelId = "wuyan94zl:im:public"
	sendMessage    = 100
)

func Run(ctx *svc.ServiceContext) {
	mux := http.NewServeMux()
	mux.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		token, _ := strconv.Atoi(r.FormValue("_token"))
		info, err := ctx.UserModel.FindOne(context.Background(), int64(token))
		if err != nil {
			return
		}
		chart.NewServer(w, r, uint64(info.Id), &data{ctx: ctx})
	})
	err := http.ListenAndServe(":8899", mux)
	if err != nil {
		return
	}
}

type data struct {
	ctx *svc.ServiceContext
}

func (d *data) SendMessage(msg chart.Message) {
	fmt.Println("send message callback ", msg)
	switch msg.Type {
	case sendMessage:
		local, _ := time.LoadLocation("Asia/Shanghai")
		sendTime, err := time.ParseInLocation("2006-01-02 15:01:05", "2022-04-15 22:12:12",local)
		fmt.Println(sendTime, err, msg.SendTime)
		if err != nil {
			return
		}
		message := messages.Messages{
			ChannelId:  msg.ChannelId,
			SendUserId: int64(msg.UserId),
			Message:    msg.Content,
			//CreateTime: sendTime,
		}
		d.ctx.MessageModel.Insert(context.Background(), &message)
	}
	fmt.Println("send message callback ", msg.ChannelId, msg.Content, msg.Type, msg.SendTime, msg.UserId)
}

// LoginServer 登录成功后回调
func (d *data) LoginServer(uid uint64) {
	list, _ := d.ctx.UserModel.Friends(d.ctx.UserUsersModel, int64(uid))
	var channelIds []string
	for _, v := range list {
		channelIds = append(channelIds, GenChannelIdByFriend(int64(uid), v.Id))
	}
	channelIds = append(channelIds, publicChanelId)
	chart.JoinChannelIds(uid, channelIds...)
}
func (d *data) LogoutServer(uid uint64) {
	// 退出登陆回调
	fmt.Println("logout ", uid)
}
func (d *data) ErrorLogServer(err error) {
	// 错误消息回调
	fmt.Println("err: ", err)
}

func SendMessageToUid(uid, toUId uint64, msg string) {
	chart.SendMessageToUid(uid, toUId, msg)
}

func SendMessageToUser(uid uint64) {
	chart.SendMessageToChannelIds(uid, "添加", "ddd")
}
