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

type cliDetail struct {
	NickName  string `json:"nick_name"`
	Phone     string `json:"phone"`
	HeadProto string `json:"head_proto"`
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

func (d *data) SendMessage(msg chart.Message) {
	fmt.Println("send message callback ", msg)
	switch msg.Type {
	case sendMessage:
		local, _ := time.LoadLocation("Asia/Shanghai")
		sendTime, err := time.ParseInLocation("2006-01-02 15:01:05", "2022-04-15 22:12:12", local)
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

func (d *data) DelaySendMessage(channelId string, msg chart.Message, uIds []uint64) {
	fmt.Println(channelId, msg, uIds)
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

func JoinChannelIds(uid uint64, channelIds ...string) {
	chart.JoinChannelIds(uid, channelIds...)
}

func SendMessageToUid(uid, toUId uint64, msg string, tp uint8) {
	chart.SendMessageToUid(uid, toUId, msg, tp)
}

func SendMessageToChannelIds(uid uint64, msg string, tp uint8, channelIds ...string) {
	chart.SendMessageToChannelIds(uid, msg, tp, channelIds...)
}

func SendMessageToUser(uid uint64) {

}
