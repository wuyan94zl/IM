package actionlogs

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ActionLogsModel = (*customActionLogsModel)(nil)

type (
	// ActionLogsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customActionLogsModel.
	ActionLogsModel interface {
		actionLogsModel
		CheckSendAddFriend(userId, friendId int64) (*ActionLogs, error)
	}

	customActionLogsModel struct {
		*defaultActionLogsModel
	}
)

// NewActionLogsModel returns a model for the database table.
func NewActionLogsModel(conn sqlx.SqlConn) ActionLogsModel {
	return &customActionLogsModel{
		defaultActionLogsModel: newActionLogsModel(conn),
	}
}

func (m *customActionLogsModel) CheckSendAddFriend(userId, friendId int64) (*ActionLogs, error) {
	var resp ActionLogs
	query := fmt.Sprintf("select %s from %s where `pub_user_id` = ? and `sub_user_id` = ? and status = 0 limit 1", actionLogsRows, m.table)
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
