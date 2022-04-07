package user

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UsersModel = (*customUsersModel)(nil)

type (
	// UsersModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUsersModel.
	UsersModel interface {
		usersModel
		FindRawByName(ctx context.Context, userName string) (*Users, error)
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
