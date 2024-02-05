package postgres

import (
	"context"
	"fmt"
	"shifolink/config"
	"shifolink/storage"

	_ "github.com/golang-migrate/migrate/v4/database"          //database is needed for migration
	_ "github.com/golang-migrate/migrate/v4/database/postgres" //postgres is used for database
	_ "github.com/golang-migrate/migrate/v4/source/file"       //file is needed for migration url

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
)

type Store struct {
	pool *pgxpool.Pool
}

func New(ctx context.Context, cfg config.Config) (storage.IStorage, error) {
	url := fmt.Sprintf(
		`postgres://%s:%s@%s:%s/%s?sslmode=disable`,
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDB,
	)

	poolConfig, err := pgxpool.ParseConfig(url)
	if err != nil {
		fmt.Println("error while parsing config", err.Error())
		return nil, err
	}

	poolConfig.MaxConns = 100

	pool, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		fmt.Println("error while connecting to db", err.Error())
		return nil, err
	}

	// // migration
	// m, err := migrate.New("file://migrations/postgres/", url)
	// if err != nil {
	// 	fmt.Println("error while migrating", err.Error())
	// 	return nil, err
	// }

	// if err = m.Up(); err != nil {
	// 	fmt.Println("here up")
	// 	if !strings.Contains(err.Error(), "no change") {
	// 		fmt.Println("in !strings")
	// 		version, dirty, err := m.Version()
	// 		if err != nil {
	// 			fmt.Println("err in checking version and dirty", err.Error())
	// 			return nil, err
	// 		}

	// 		if dirty {
	// 			version--
	// 			if err = m.Force(int(version)); err != nil {
	// 				fmt.Println("ERR in making force", err.Error())
	// 				return nil, err
	// 			}
	// 		}
	// 		fmt.Println("ERROR in migrating", err.Error())
	// 		return nil, err
	// 	}
	// }

	return Store{
		pool: pool,
	}, nil

}

func (s Store) CloseDB() {
	s.pool.Close()
}