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

func (s Store) Author() storage.IAuthorRepo {
	return NewAuthorRepo(s.pool)
}

func (s Store) ClinicAdmin() storage.IClinicAdminRepo {
	return NewClinicAdminRepo(s.pool)
}

func (s Store) ClinicBranch() storage.IClinicBranchRepo {
	return NewClinicBranchRepo(s.pool)
}

func (s Store) Clinic() storage.IClinicRepo {
	return NewClinicRepo(s.pool)
}

func (s Store) Customer() storage.ICustomerRepo {
	return NewCustomerRepo(s.pool)
}

func (s Store) DoctorType() storage.IDoctorTypeRepo {
	return NewDoctorTypeRepo(s.pool)
}

func (s Store) Doctor() storage.IDoctorRepo {
	return NewDoctorRepo(s.pool)
}

func (s Store) DrugStoreBranch() storage.IDrugStoreBranchRepo {
	return NewDrugStoreBranchRepo(s.pool)
}

func (s Store) DrugStore() storage.IDrugStoreRepo {
	return NewDrugStoreRepo(s.pool)
}

func (s Store) Drug() storage.IDrugRepo {
	return NewDrugRepo(s.pool)
}

func (s Store) Journal() storage.IJournalRepo {
	return NewJournalRepo(s.pool)
}

func (s Store) OrderDrug() storage.IOrderDrugRepo {
	return NewOrderDrugRepo(s.pool)
}

func (s Store) Orders() storage.IOrdersRepo {
	return NewOrdersRepo(s.pool)
}

func (s Store) Pharmacist() storage.IPharmacistRepo {
	return NewPharmacistRepo(s.pool)
}

func (s Store) Queue() storage.IQueueRepo {
	return NewQueueRepo(s.pool)
}

func (s Store) SuperAdmin() storage.ISuperAdminRepo {
	return NewSuperAdminRepo(s.pool)
}
