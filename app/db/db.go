package db

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func ConnectionDB() (*sqlx.DB, error) {
	dns := "host=localhost user=postgres_fraud_score password=1234 dbname=fraud_score port=5432 sslmode=disable"
	db, err := sqlx.Connect("postgres", dns)
	if err != nil {
		log.Fatalln("Error on database conection. Error: ", err)
	}

	if err := createTables(db); err != nil {
		log.Fatalln("Error on tables database. Error: ", err)
	}

	fmt.Println("tables created")
	fmt.Println("Database connected")
	return db, nil
}

func createTables(db *sqlx.DB) error {
	schema := `
		DO $$
		BEGIN
			IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'risk_level') THEN
				CREATE TYPE risk_level AS ENUM ('low', 'medium', 'high', 'critical');
			END IF;

			IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'risk_action') THEN
				CREATE TYPE risk_action AS ENUM ('approve', 'review', 'block');
			END IF;
		END $$;

		CREATE TABLE IF NOT EXISTS tb_transactions (
			id          UUID          PRIMARY KEY,
			account_id  TEXT          NOT NULL,
			amount      NUMERIC(15,2) NOT NULL,
			currency    TEXT          NOT NULL DEFAULT 'BRL',
			country     TEXT          NOT NULL,
			merchant    TEXT          NOT NULL,
			ip_address  TEXT          NOT NULL,
			occurred_at TIMESTAMPTZ   NOT NULL
		);

		CREATE TABLE IF NOT EXISTS tb_scoring_results (
			id             UUID        PRIMARY KEY,
			transaction_id UUID        NOT NULL REFERENCES tb_transactions(id),
			score          SMALLINT    NOT NULL CHECK (score >= 0 AND score <= 100),
			risk           risk_level  NOT NULL,
			action         risk_action NOT NULL,
			reasons        TEXT[]      NOT NULL DEFAULT '{}'
		);

		CREATE INDEX IF NOT EXISTS idx_transactions_account_id   ON tb_transactions(account_id);
		CREATE INDEX IF NOT EXISTS idx_transactions_occurred_at  ON tb_transactions(occurred_at);
		CREATE INDEX IF NOT EXISTS idx_scoring_results_action    ON tb_scoring_results(action);
		CREATE INDEX IF NOT EXISTS idx_scoring_results_tx_id     ON tb_scoring_results(transaction_id);
	`
	_, err := db.Exec(schema)
	return err
}
