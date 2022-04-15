package user

import (
	"context"
	"fmt"
	"github.com/wuyan94zl/go-zero-blog/app/models/hasusers"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"strings"
)

var _ UsersModel = (*customUsersModel)(nil)

type (
	// UsersModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUsersModel.
	UsersModel interface {
		usersModel
		FindRawByName(ctx context.Context, userName string) (*Users, error)
		Friends(hasUserModel hasusers.UserUsersModel, userId int64) ([]Users, error)
	}

	customUsersModel struct {
		*defaultUsersModel
		sqlWhere string
	}
)

// NewUsersModel returns a model for the database table.
func NewUsersModel(conn sqlx.SqlConn, c cache.CacheConf) UsersModel {
	return &customUsersModel{
		defaultUsersModel: newUsersModel(conn, c),
	}
}

func (m *customUsersModel) FindRawByName(ctx context.Context, userName string) (*Users, error) {
	var resp Users
	err := m.QueryRowNoCache(&resp, fmt.Sprintf("select %s from %s where `user_name` = ? limit 1", usersRows, m.table), userName)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *customUsersModel) Friends(hasUserModel hasusers.UserUsersModel, userId int64) ([]Users, error) {
	// 关联数据
	hasUsers, err := hasUserModel.Friends(userId)
	if err != nil {
		return nil, err
	}
	var friendIds []interface{}
	for _, friend := range hasUsers {
		friendIds = append(friendIds, friend.HasUserId)
	}

	var resp []Users
	ids := strings.Repeat(" ,?", len(friendIds))
	query := fmt.Sprintf("select %s from %s where `id` in (%s)", usersRows, m.table, ids[2:])
	err = m.QueryRowsNoCacheCtx(context.Background(), &resp, query, friendIds...)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
