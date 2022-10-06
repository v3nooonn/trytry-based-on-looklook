// Code generated by goctl. DO NOT EDIT!

package car

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
	productionCarFieldNames          = builder.RawFieldNames(&ProductionCar{})
	productionCarRows                = strings.Join(productionCarFieldNames, ",")
	productionCarRowsExpectAutoSet   = strings.Join(stringx.Remove(productionCarFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), ",")
	productionCarRowsWithPlaceHolder = strings.Join(stringx.Remove(productionCarFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), "=?,") + "=?"

	cacheTrytryProductionCarIdPrefix = "cache:trytry:productionCar:id:"
)

type (
	productionCarModel interface {
		Insert(ctx context.Context, data *ProductionCar) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*ProductionCar, error)
		Update(ctx context.Context, data *ProductionCar) error
		Delete(ctx context.Context, id int64) error
	}

	defaultProductionCarModel struct {
		sqlc.CachedConn
		table string
	}

	ProductionCar struct {
		Id        int64     `db:"id"`
		Brand     string    `db:"brand"`
		Serie     string    `db:"serie"`
		CreatedAt time.Time `db:"created_at"`
		UpdatedAt time.Time `db:"updated_at"`
	}
)

func newProductionCarModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultProductionCarModel {
	return &defaultProductionCarModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`production_car`",
	}
}

func (m *defaultProductionCarModel) Delete(ctx context.Context, id int64) error {
	trytryProductionCarIdKey := fmt.Sprintf("%s%v", cacheTrytryProductionCarIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, trytryProductionCarIdKey)
	return err
}

func (m *defaultProductionCarModel) FindOne(ctx context.Context, id int64) (*ProductionCar, error) {
	trytryProductionCarIdKey := fmt.Sprintf("%s%v", cacheTrytryProductionCarIdPrefix, id)
	var resp ProductionCar
	err := m.QueryRowCtx(ctx, &resp, trytryProductionCarIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", productionCarRows, m.table)
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

func (m *defaultProductionCarModel) Insert(ctx context.Context, data *ProductionCar) (sql.Result, error) {
	trytryProductionCarIdKey := fmt.Sprintf("%s%v", cacheTrytryProductionCarIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, productionCarRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Brand, data.Serie, data.CreatedAt, data.UpdatedAt)
	}, trytryProductionCarIdKey)
	return ret, err
}

func (m *defaultProductionCarModel) Update(ctx context.Context, data *ProductionCar) error {
	trytryProductionCarIdKey := fmt.Sprintf("%s%v", cacheTrytryProductionCarIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, productionCarRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.Brand, data.Serie, data.CreatedAt, data.UpdatedAt, data.Id)
	}, trytryProductionCarIdKey)
	return err
}

func (m *defaultProductionCarModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheTrytryProductionCarIdPrefix, primary)
}

func (m *defaultProductionCarModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", productionCarRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultProductionCarModel) tableName() string {
	return m.table
}
