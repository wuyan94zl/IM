package groupusers

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ GroupUsersModel = (*customGroupUsersModel)(nil)

type (
	// GroupUsersModel is an interface to be customized, add more methods here,
	// and implement the added methods in customGroupUsersModel.
	GroupUsersModel interface {
		groupUsersModel
		TranCreate(ctx context.Context, session sqlx.Session, groupUserItem *GroupUsers) error
	}

	customGroupUsersModel struct {
		*defaultGroupUsersModel
	}
)

// NewGroupUsersModel returns a model for the database table.
func NewGroupUsersModel(conn sqlx.SqlConn) GroupUsersModel {
	return &customGroupUsersModel{
		defaultGroupUsersModel: newGroupUsersModel(conn),
	}
}

func (m *defaultGroupUsersModel) TranCreate(ctx context.Context, session sqlx.Session, data *GroupUsers) error {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, groupUsersRowsExpectAutoSet)
	_, err := session.ExecCtx(ctx, query, data.GroupId, data.UserId, data.ChannelId, data.IsManager)
	if err != nil {
		return err
	}
	return nil
}
