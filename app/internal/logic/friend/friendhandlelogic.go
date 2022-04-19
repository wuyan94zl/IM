package friend

import (
	"context"
	"encoding/json"
	"github.com/wuyan94zl/go-zero-blog/app/common/im"
	"github.com/wuyan94zl/go-zero-blog/app/models/hasusers"
	"github.com/wuyan94zl/go-zero-blog/app/models/notices"

	"github.com/wuyan94zl/go-zero-blog/app/internal/svc"
	"github.com/wuyan94zl/go-zero-blog/app/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FriendHandleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFriendHandleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendHandleLogic {
	return &FriendHandleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FriendHandleLogic) FriendHandle(req *types.FriendHandleRequest) (resp *types.FriendResponse, err error) {
	id, _ := l.ctx.Value("id").(json.Number).Int64()
	log, err := l.svcCtx.NoticeModel.FindOne(l.ctx, req.ActionLogId)
	if err != nil {
		return nil, err
	}

	if log.SubUserId != id || log.Status == 1 {
		return &types.FriendResponse{
			Status:  false,
			Message: "参数错误或已处理",
		}, nil
	}

	friend, err := l.svcCtx.UserUsersModel.CheckFriend(log.PubUserId, log.SubUserId)
	if len(friend) > 0 || err != nil {
		return &types.FriendResponse{Status: false, Message: "已添加完成"}, nil
	}

	log.Status = 1
	link, _ := l.svcCtx.NoticeModel.FindOne(l.ctx, log.LinkId)
	return l.setFriend(log, link, req.ActionType, id)
	//switch req.ActionType {
	//case 1:
	//	l.svcCtx.UserUsersModel.Insert(l.ctx, &hasusers.UserUsers{UserId: id, HasUserId: log.PubUserId})
	//	l.svcCtx.UserUsersModel.Insert(l.ctx, &hasusers.UserUsers{UserId: log.PubUserId, HasUserId: id})
	//	log.IsAgree = "已同意"
	//	l.svcCtx.NoticeModel.Update(l.ctx, log)
	//	link.IsAgree = "已同意"
	//	l.svcCtx.NoticeModel.Update(l.ctx, link)
	//	strByte, _ := json.Marshal(types.Notice{
	//		Id: log.Id, Tp: log.Tp, IsAgree: log.IsAgree,
	//		NickName: "", Content: log.Content, CreateTime: log.CreateTime.Format("2006-01-06 15:04:01"),
	//	})
	//	im.SendMessageToUid(uint64(id), uint64(log.PubUserId), string(strByte))
	//	//im.SendMessageToUid(uint64(id), uint64(log.PubUserId), "同意了你的好友申请")
	//	return &types.FriendResponse{Status: true, Message: "添加成功"}, nil
	//default:
	//	log.IsAgree = "已拒绝"
	//	l.svcCtx.NoticeModel.Update(l.ctx, log)
	//	link.IsAgree = "被拒绝"
	//	l.svcCtx.NoticeModel.Update(l.ctx, link)
	//	strByte, _ := json.Marshal(types.Notice{
	//		Id: log.Id, Tp: log.Tp, IsAgree: log.IsAgree,
	//		NickName: "", Content: log.Content, CreateTime: log.CreateTime.Format("2006-01-06 15:04:01"),
	//	})
	//	im.SendMessageToUid(uint64(id), uint64(log.PubUserId), string(strByte))
	//	//im.SendMessageToUid(uint64(id), uint64(log.PubUserId), "拒绝了你的好友申请")
	//	return &types.FriendResponse{Status: false, Message: "拒绝添加好友"}, nil
	//}
}

func (l *FriendHandleLogic) setFriend(log, link *notices.Notices, tp, id int64) (resp *types.FriendResponse, err error) {
	rlt := new(types.FriendResponse)
	if tp == 1 {
		l.svcCtx.UserUsersModel.Insert(l.ctx, &hasusers.UserUsers{UserId: id, HasUserId: log.PubUserId})
		l.svcCtx.UserUsersModel.Insert(l.ctx, &hasusers.UserUsers{UserId: log.PubUserId, HasUserId: id})
		log.IsAgree = "已同意"
		link.IsAgree = "已同意"
		rlt.Status = true
		rlt.Message = "添加成功"
	} else {
		log.IsAgree = "已拒绝"
		link.IsAgree = "被拒绝"
		rlt.Status = false
		rlt.Message = "拒绝添加好友"
	}
	l.svcCtx.NoticeModel.Update(l.ctx, log)
	l.svcCtx.NoticeModel.Update(l.ctx, link)
	strByte, _ := json.Marshal(types.Notice{
		Id: log.Id, Tp: log.Tp, IsAgree: log.IsAgree,
		NickName: "", Content: log.Content, CreateTime: log.CreateTime.Format("2006-01-06 15:04:01"),
	})
	im.SendMessageToUid(uint64(id), uint64(log.PubUserId), string(strByte))
	return rlt, nil
}
