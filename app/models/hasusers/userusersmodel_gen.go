// Code generated by goctl. DO NOT EDIT!

package hasusers

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	userUsersFieldNames          = builder.RawFieldNames(&UserUsers{})
	userUsersRows                = strings.Join(userUsersFieldNames, ",")
	userUsersRowsExpectAutoSet   = strings.Join(stringx.Remove(userUsersFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	userUsersRowsWithPlaceHolder = strings.Join(stringx.Remove(userUsersFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheUserUsersIdPrefix = "cache:userUsers:id:"
)

type (
	userUsersModel interface {
		Insert(ctx context.Context, data *UserUsers) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*UserUsers, error)
		Update(ctx context.Context, data *UserUsers) error
		Delete(ctx context.Context, id int64) error
	}

	defaultUserUsersModel struct {
		sqlc.CachedConn
		table string
	}

	UserUsers struct {
		Id        int64 `db:"id"`
		UserId    int64 `db:"user_id"`
		HasUserId int64 `db:"has_user_id"`
	}
)

func newUserUsersModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultUserUsersModel {
	return &defaultUserUsersModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`user_users`",
	}
}

func (m *defaultUserUsersModel) Insert(ctx context.Context, data *UserUsers) (sql.Result, error) {
	userUsersIdKey := fmt.Sprintf("%s%v", cacheUserUsersIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, userUsersRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.UserId, data.HasUserId)
	}, userUsersIdKey)
	return ret, err
}

func (m *defaultUserUsersModel) FindOne(ctx context.Context, id int64) (*UserUsers, error) {
	userUsersIdKey := fmt.Sprintf("%s%v", cacheUserUsersIdPrefix, id)
	var resp UserUsers
	err := m.QueryRowCtx(ctx, &resp, userUsersIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", userUsersRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserUsersModel) Update(ctx context.Context, data *UserUsers) error {
	userUsersIdKey := fmt.Sprintf("%s%v", cacheUserUsersIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, userUsersRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.UserId, data.HasUserId, data.Id)
	}, userUsersIdKey)
	return err
}

func (m *defaultUserUsersModel) Delete(ctx context.Context, id int64) error {
	userUsersIdKey := fmt.Sprintf("%s%v", cacheUserUsersIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, userUsersIdKey)
	return err
}

func (m *defaultUserUsersModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheUserUsersIdPrefix, primary)
}

func (m *defaultUserUsersModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", userUsersRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultUserUsersModel) tableName() string {
	return m.table
}
