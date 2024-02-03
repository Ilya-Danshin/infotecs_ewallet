package walletRepository

import (
	"context"
	"time"

	"github.com/gofrs/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"

	"EWallet/internal/models"
)

type Postgres struct {
	conn *pgxpool.Pool
}

func New(conn *pgxpool.Pool) (*Postgres, error) {
	return &Postgres{conn: conn}, nil
}

const insertWallet = `INSERT INTO wallet_balance(id, balance)
				VALUES ($1, $2) ON CONFLICT (id) DO UPDATE 
				SET balance=excluded.balance`

func (db *Postgres) InsertWallet(ctx context.Context, walletId uuid.UUID, balance float32) error {
	row, err := db.conn.Query(ctx, insertWallet, walletId, balance)
	if err != nil {
		return err
	}
	defer row.Close()

	return nil
}

const selectWallet = `SELECT id, balance 
						FROM wallet_balance 
						WHERE id = $1`

func (db *Postgres) SelectWallet(ctx context.Context, walletId uuid.UUID) (*models.Wallet, error) {
	res := &models.Wallet{}
	err := db.conn.QueryRow(ctx, selectWallet, walletId).Scan(
		&res.Id,
		&res.Balance)
	if err == pgx.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return res, nil
}

const walletFromUpdate = `UPDATE wallet_balance
							SET balance=balance-$1
							WHERE id=$2`

const walletToUpdate = `UPDATE wallet_balance
							SET balance=balance+$1
							WHERE id=$2`

func (db *Postgres) UpdateWalletBalance(ctx context.Context, from, to uuid.UUID, amount float32) error {
	tx, err := db.conn.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	_, err = tx.Exec(ctx, walletFromUpdate, amount, from)
	if err != nil {
		return err
	}

	_, err = tx.Exec(ctx, walletToUpdate, amount, to)
	if err != nil {
		return err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (db *Postgres) DeleteWallet(ctx context.Context, walletId uuid.UUID) error {
	return nil
}

func (db *Postgres) InsertTransaction(ctx context.Context, time time.Time, from uuid.UUID, to uuid.UUID, money float32) error {
	return nil
}

func (db *Postgres) SelectTransactionsByWallet(ctx context.Context, walletId uuid.UUID) ([]*models.Transaction, error) {
	return nil, nil
}
