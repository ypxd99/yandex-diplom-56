package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/ypxd99/yandex-diplom-56/internal/model"
	"github.com/ypxd99/yandex-diplom-56/internal/repository"
)

func (r *PostgresRepo) CreateUser(ctx context.Context, user *model.User) error {
	_, err := r.db.NewInsert().Model(user).Exec(ctx)
	return err
}

func (r *PostgresRepo) GetUserByLogin(ctx context.Context, login string) (*model.User, error) {
	user := new(model.User)
	err := r.db.NewSelect().Model(user).Where("login = ?", login).Scan(ctx)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return user, err
}

func (r *PostgresRepo) GetUserByID(ctx context.Context, id uuid.UUID) (*model.User, error) {
	user := new(model.User)
	err := r.db.NewSelect().Model(user).Where("id = ?", id).Scan(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, repository.ErrUserNotFound
		}
		return nil, err
	}
	return user, nil
}

func (r *PostgresRepo) CreateOrder(ctx context.Context, order *model.Order) error {
	_, err := r.db.NewInsert().Model(order).Exec(ctx)
	return err
}

func (r *PostgresRepo) GetOrderByNumber(ctx context.Context, number string) (*model.Order, error) {
	order := new(model.Order)
	err := r.db.NewSelect().Model(order).Where("number = ?", number).Scan(ctx)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return order, err
}

func (r *PostgresRepo) GetOrderByID(ctx context.Context, id uuid.UUID) (*model.Order, error) {
	order := new(model.Order)
	err := r.db.NewSelect().Model(order).Where("id = ?", id).Scan(ctx)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return order, err
}

func (r *PostgresRepo) GetUserOrders(ctx context.Context, userID uuid.UUID) ([]*model.Order, error) {
	var orders []*model.Order
	err := r.db.NewSelect().
		Model(&orders).
		Where("user_id = ?", userID).
		Order("uploaded_at DESC").
		Scan(ctx)
	return orders, err
}

func (r *PostgresRepo) UpdateOrderStatus(ctx context.Context, orderID uuid.UUID, status model.OrderStatus, accrual float64) error {
	_, err := r.db.NewUpdate().
		Model((*model.Order)(nil)).
		Set("status = ?", status).
		Set("accrual = ?", accrual).
		Where("id = ?", orderID).
		Exec(ctx)
	return err
}

func (r *PostgresRepo) GetUserBalance(ctx context.Context, userID uuid.UUID) (*model.UserBalance, error) {
	balance := new(model.UserBalance)
	err := r.db.NewSelect().
		Model(balance).
		Where("user_id = ?", userID).
		Scan(ctx)
	if err == sql.ErrNoRows {
		balance = &model.UserBalance{
			UserID:    userID,
			Current:   0,
			Withdrawn: 0,
		}
		_, err = r.db.NewInsert().Model(balance).Exec(ctx)
		if err != nil {
			return nil, err
		}
		return balance, nil
	}
	return balance, err
}

func (r *PostgresRepo) UpdateUserBalance(ctx context.Context, userID uuid.UUID, current, withdrawn float64) error {
	_, err := r.db.NewUpdate().
		Model((*model.UserBalance)(nil)).
		Set("current = ?", current).
		Set("withdrawn = ?", withdrawn).
		Where("user_id = ?", userID).
		Exec(ctx)
	return err
}

func (r *PostgresRepo) CreateWithdrawal(ctx context.Context, withdrawal *model.Withdrawal) error {
	_, err := r.db.NewInsert().Model(withdrawal).Exec(ctx)
	return err
}

func (r *PostgresRepo) GetUserWithdrawals(ctx context.Context, userID uuid.UUID) ([]*model.Withdrawal, error) {
	var withdrawals []*model.Withdrawal
	err := r.db.NewSelect().
		Model(&withdrawals).
		Where("user_id = ?", userID).
		Order("processed_at DESC").
		Scan(ctx)
	return withdrawals, err
}

func (r *PostgresRepo) CleanupTables(ctx context.Context) error {
	_, err := r.db.NewDelete().Model((*model.Withdrawal)(nil)).Where("1=1").Exec(ctx)
	if err != nil {
		return err
	}
	_, err = r.db.NewDelete().Model((*model.Order)(nil)).Where("1=1").Exec(ctx)
	if err != nil {
		return err
	}
	_, err = r.db.NewDelete().Model((*model.UserBalance)(nil)).Where("1=1").Exec(ctx)
	if err != nil {
		return err
	}
	_, err = r.db.NewDelete().Model((*model.User)(nil)).Where("1=1").Exec(ctx)
	return err
}
