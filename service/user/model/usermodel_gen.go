// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	userFieldNames          = builder.RawFieldNames(&User{})
	userRows                = strings.Join(userFieldNames, ",")
	userRowsExpectAutoSet   = strings.Join(stringx.Remove(userFieldNames, "`id`", "`create_at`", "`created_at`", "`create_time`", "`update_at`", "`updated_at`", "`update_time`"), ",")
	userRowsWithPlaceHolder = strings.Join(stringx.Remove(userFieldNames, "`id`", "`create_at`", "`created_at`", "`create_time`", "`update_at`", "`updated_at`", "`update_time`"), "=?,") + "=?"

	cacheUserIdPrefix     = "cache:user:id:"
	cacheUserMobilePrefix = "cache:user:mobile:"
)

type (
	userModel interface {
		Insert(ctx context.Context, data *User) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*User, error)
		FindOneByMobile(ctx context.Context, mobile string) (*User, error)
		Update(ctx context.Context, data *User) error
		Delete(ctx context.Context, id int64) error
	}

	defaultUserModel struct {
		sqlc.CachedConn
		table string
	}

	User struct {
		Id         int64     `db:"id"`
		Name       string    `db:"name"`     // 用户姓名
		Gender     int64     `db:"gender"`   // 用户性别
		Mobile     string    `db:"mobile"`   // 用户电话
		Password   string    `db:"password"` // 用户密码
		CreateTime time.Time `db:"create_time"`
		UpdateTime time.Time `db:"update_time"`
	}
)

func newUserModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultUserModel {
	return &defaultUserModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`user`",
	}
}

func (m *defaultUserModel) Delete(ctx context.Context, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	userIdKey := fmt.Sprintf("%s%v", cacheUserIdPrefix, id)
	userMobileKey := fmt.Sprintf("%s%v", cacheUserMobilePrefix, data.Mobile)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, userIdKey, userMobileKey)
	return err
}

func (m *defaultUserModel) FindOne(ctx context.Context, id int64) (*User, error) {
	userIdKey := fmt.Sprintf("%s%v", cacheUserIdPrefix, id)
	var resp User
	err := m.QueryRowCtx(ctx, &resp, userIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", userRows, m.table)
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

func (m *defaultUserModel) FindOneByMobile(ctx context.Context, mobile string) (*User, error) {
	userMobileKey := fmt.Sprintf("%s%v", cacheUserMobilePrefix, mobile)
	var resp User
	err := m.QueryRowIndexCtx(ctx, &resp, userMobileKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `mobile` = ? limit 1", userRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, mobile); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserModel) Insert(ctx context.Context, data *User) (sql.Result, error) {
	userIdKey := fmt.Sprintf("%s%v", cacheUserIdPrefix, data.Id)
	userMobileKey := fmt.Sprintf("%s%v", cacheUserMobilePrefix, data.Mobile)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, userRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Name, data.Gender, data.Mobile, data.Password)
	}, userIdKey, userMobileKey)
	return ret, err
}

func (m *defaultUserModel) Update(ctx context.Context, newData *User) error {
	data, err := m.FindOne(ctx, newData.Id)
	if err != nil {
		return err
	}

	userIdKey := fmt.Sprintf("%s%v", cacheUserIdPrefix, data.Id)
	userMobileKey := fmt.Sprintf("%s%v", cacheUserMobilePrefix, data.Mobile)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, userRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, newData.Name, newData.Gender, newData.Mobile, newData.Password, newData.Id)
	}, userIdKey, userMobileKey)
	return err
}

func (m *defaultUserModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheUserIdPrefix, primary)
}

func (m *defaultUserModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", userRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultUserModel) tableName() string {
	return m.table
}
