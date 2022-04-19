package notices

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"time"
)

var _ NoticesModel = (*customNoticesModel)(nil)

type (
	// NoticesModel is an interface to be customized, add more methods here,
	// and implement the added methods in customNoticesModel.
	NoticesModel interface {
		noticesModel
		CheckSendAddFriend(userId, friendId int64) (*Notices, error)
		GetListByUserId(userId int64) ([]ListItem, error)
	}

	customNoticesModel struct {
		*defaultNoticesModel
	}
)

type (
	ListItem struct {
		Id         int64     `db:"id" json:"id"`
		Tp         int64     `db:"type" json:"type"`
		IsAgree    string    `db:"is_agree" json:"is_agree"`
		NickName   string    `db:"nick_name" json:"nick_name"`
		Content    string    `db:"content" json:"content"`
		CreateTime time.Time `db:"create_time" json:"create_time"`
		Status     int64     `db:"status" json:"status"`
	}
)

// NewNoticesModel returns a model for the database table.
func NewNoticesModel(conn sqlx.SqlConn) NoticesModel {
	return &customNoticesModel{
		defaultNoticesModel: newNoticesModel(conn),
	}
}

func (m *customNoticesModel) CheckSendAddFriend(userId, friendId int64) (*Notices, error) {
	var resp Notices
	query := fmt.Sprintf("select %s from %s where `pub_user_id` = ? and `sub_user_id` = ? and status = 0 limit 1", noticesRows, m.table)
	err := m.conn.QueryRow(&resp, query, userId, friendId)
	switch err {
	case ErrNotFound:
		return nil, nil
	case nil:
		return &resp, nil
	default:
		return nil, err
	}
}

func (m *customNoticesModel) GetListByUserId(userId int64) ([]ListItem, error) {
	var resp []ListItem
	selectRows := "notices.`id`,notices.`type`,notices.`content`,notices.`is_agree`,notices.`create_time`,notices.`status`,users.`nick_name`"
	query := fmt.Sprintf("select %s from %s join users on %s.pub_user_id = users.id where `sub_user_id` = ? order by id desc limit 20", selectRows, m.table, m.table)
	err := m.conn.QueryRows(&resp, query, userId)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}
