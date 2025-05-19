package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/ypxd99/yandex-diplom-56/internal/model"
	"github.com/ypxd99/yandex-diplom-56/internal/repository"
)

func (p *PostgresRepo) CreateUser(ctx context.Context, user *model.User) error {
	_, err := p.db.NewInsert().Model(user).Exec(ctx)
	return err
}

func (p *PostgresRepo) GetUserByLogin(ctx context.Context, login string) (*model.User, error) {
	user := new(model.User)
	err := p.db.NewSelect().Model(user).Where("login = ?", login).Scan(ctx)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return user, err
}

func (p *PostgresRepo) GetUserByID(ctx context.Context, id uuid.UUID) (*model.User, error) {
	user := new(model.User)
	err := p.db.NewSelect().Model(user).Where("id = ?", id).Scan(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, repository.ErrUserNotFound
		}
		return nil, err
	}
	return user, nil
}

func (p *PostgresRepo) CreateOrder(ctx context.Context, order *model.Order) error {
	_, err := p.db.NewInsert().Model(order).Exec(ctx)
	return err
}

func (p *PostgresRepo) GetOrderByNumber(ctx context.Context, number string) (*model.Order, error) {
	order := new(model.Order)
	err := p.db.NewSelect().Model(order).Where("number = ?", number).Scan(ctx)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return order, err
}

func (p *PostgresRepo) GetOrderByID(ctx context.Context, id uuid.UUID) (*model.Order, error) {
	order := new(model.Order)
	err := p.db.NewSelect().Model(order).Where("id = ?", id).Scan(ctx)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return order, err
}

func (p *PostgresRepo) GetUserOrders(ctx context.Context, userID uuid.UUID) ([]*model.Order, error) {
	var orders []*model.Order
	err := p.db.NewSelect().
		Model(&orders).
		Where("user_id = ?", userID).
		Order("uploaded_at DESC").
		Scan(ctx)
	return orders, err
}

func (p *PostgresRepo) UpdateOrderStatus(ctx context.Context, orderID uuid.UUID, status model.OrderStatus, accrual float64) error {
	_, err := p.db.NewUpdate().
		Model((*model.Order)(nil)).
		Set("status = ?", status).
		Set("accrual = ?", accrual).
		Where("id = ?", orderID).
		Exec(ctx)
	return err
}

func (p *PostgresRepo) GetUserBalance(ctx context.Context, userID uuid.UUID) (*model.UserBalance, error) {
	balance := new(model.UserBalance)
	err := p.db.NewSelect().
		Model(balance).
		Where("user_id = ?", userID).
		Scan(ctx)
	if err == sql.ErrNoRows {
		balance = &model.UserBalance{
			UserID:    userID,
			Current:   0,
			Withdrawn: 0,
		}
		_, err = p.db.NewInsert().Model(balance).Exec(ctx)
		if err != nil {
			return nil, err
		}
		return balance, nil
	}
	return balance, err
}

func (p *PostgresRepo) UpdateUserBalance(ctx context.Context, userID uuid.UUID, current, withdrawn float64) error {
	_, err := p.db.NewUpdate().
		Model((*model.UserBalance)(nil)).
		Set("current = ?", current).
		Set("withdrawn = ?", withdrawn).
		Where("user_id = ?", userID).
		Exec(ctx)
	return err
}

func (p *PostgresRepo) CreateWithdrawal(ctx context.Context, withdrawal *model.Withdrawal) error {
	_, err := p.db.NewInsert().Model(withdrawal).Exec(ctx)
	return err
}

func (p *PostgresRepo) GetUserWithdrawals(ctx context.Context, userID uuid.UUID) ([]*model.Withdrawal, error) {
	var withdrawals []*model.Withdrawal
	err := p.db.NewSelect().
		Model(&withdrawals).
		Where("user_id = ?", userID).
		Order("processed_at DESC").
		Scan(ctx)
	return withdrawals, err
}
