package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
	"zavad/config"
	"zavad/internal/entities"
	_ "zavad/internal/entities"
	// protos "zavad/pkg/proto/gen/go"
)

type Repository struct {
	ctx context.Context
	log *zap.Logger
	cfg *config.ConfigModel
	DB  *pgxpool.Pool
}

func NewRepository(log *zap.Logger, cfg *config.ConfigModel, ctx context.Context) (*Repository, error) {
	return &Repository{
		ctx: ctx,
		log: log,
		cfg: cfg,
	}, nil
}

func (r *Repository) OnStart(_ context.Context) error {
	connectionUrl := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		r.cfg.Postgres.Host,
		r.cfg.Postgres.Port,
		r.cfg.Postgres.User,
		r.cfg.Postgres.Password,
		r.cfg.Postgres.DBName,
		r.cfg.Postgres.SSLMode)
	pool, err := pgxpool.Connect(r.ctx, connectionUrl)
	if err != nil {
		return err
	}
	r.DB = pool
	return nil
}

func (r *Repository) OnStop(_ context.Context) error {
	r.DB.Close()
	return nil
}

const qGetUserID = `
select 
    id 
from 
    users 
where 
    email = $2 or phone = $1
`

func (r *Repository) GetUserIDbyPhoneOrEmail(ctx context.Context, phone, email string) (id int, err error) {
	err = r.DB.QueryRow(ctx, qGetUserID, phone, email).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

const qCreateUser = `
INSERT INTO users (id, name, email, phone) VALUES ($1, $2, $3, $4)
`

func (r *Repository) CreateUser(ctx context.Context, user *entities.User) error {
	_, err := r.DB.Exec(ctx, qCreateUser, user.ID, user.Name, user.Email, user.Phone)
	if err != nil {
		return err
	}

	return nil
}

