package hasusers

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"strings"
)

var _ UserUsersModel = (*customUserUsersModel)(nil)

type (
	// UserUsersModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserUsersModel.
	UserUsersModel interface {
		userUsersModel
		CheckFriend(userId, friendId int64) ([]UserUsers, error)
		IsFriend(userId int64, friendId ...interface{}) ([]UserUsers, error)
		Friends(userId int64) ([]UserUsers, error)
	}

	customUserUsersModel struct {
		*defaultUserUsersModel
	}
)

// NewUserUsersModel returns a model for the database table.
func NewUserUsersModel(conn sqlx.SqlConn, c cache.CacheConf) UserUsersModel {
	return &customUserUsersModel{
		defaultUserUsersModel: newUserUsersModel(conn, c),
	}
}

func (m *customUserUsersModel) CheckFriend(userId, friendId int64) ([]UserUsers, error) {
	var resp []UserUsers
	query := fmt.Sprintf("select %s from %s where (`user_id` = ? and `has_user_id` = ?) or (`user_id` = ? and `has_user_id` = ?) limit 2", userUsersRows, m.table)
	err := m.QueryRowsNoCache(&resp, query, userId, friendId, friendId, userId)
	if err == nil {
		return resp, err
	} else {
		return nil, err
	}
}

func (m *customUserUsersModel) IsFriend(userId int64, friendId ...interface{}) ([]UserUsers, error) {
	var resp []UserUsers
	if len(friendId) < 1 {
		return nil, fmt.Errorf("参数错误")
	}
	fid := strings.Repeat(",?", len(friendId))
	query := fmt.Sprintf("select %s from %s where `user_id` = ? and `has_user_id` in (%s) limit 2", userUsersRows, m.table, fid[1:])
	var param []interface{}
	param = append(param, userId)
	param = append(param, friendId...)
	err := m.QueryRowsNoCache(&resp, query, param...)
	if err == nil {
		return resp, err
	} else {
		return nil, err
	}
}

func (m *customUserUsersModel) Friends(userId int64) ([]UserUsers, error) {
	var resp []UserUsers
	query := fmt.Sprintf("select %s from %s where `user_id` = ?", userUsersRows, m.table)
	err := m.QueryRowsNoCache(&resp, query, userId)
	if err == nil {
		return resp, nil
	} else {
		return nil, err
	}
}
