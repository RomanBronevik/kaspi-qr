package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/spf13/viper"
	"kaspi-qr/pkg/utils"
	"log"
	"os"
	"time"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
}

//func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
//	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
//		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
//
//	if err != nil {
//		return nil, err
//	}
//
//	err = db.Ping()
//	if err != nil {
//		return nil, err
//	}
//
//	return db, nil
//}

type StorageConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Database string `json:"database"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Client interface {
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Begin(ctx context.Context) (pgx.Tx, error)
}

func NewClient(ctx context.Context) (pool *pgxpool.Pool, err error) {
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s",
		viper.GetString("db.username"),
		os.Getenv("DB_PASSWORD"),
		viper.GetString("db.host"),
		viper.GetString("db.port"),
		viper.GetString("dbname"))

	err = utils.DoWithTries(func() error {
		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		pool, err = pgxpool.Connect(ctx, dsn)

		if err != nil {
			fmt.Println("failed to connect to postgresql")
			return err
		}

		return nil
	}, viper.GetInt("db.maxAttempts"), 5*time.Second)

	if err != nil {
		log.Fatal("error do with tries postgresql")
		return nil, err
	}

	return pool, nil
}

func GetConfig() *StorageConfig {
	return &StorageConfig{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Database: viper.GetString("db.dbname"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
	}
}
