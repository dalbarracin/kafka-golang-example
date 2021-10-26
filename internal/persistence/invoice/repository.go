package invoice

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

const (
	host     = "postgres"
	port     = 5432
	user     = "postgres"
	password = "postgrespass"
	dbname   = "kafka_messages"
)

type InvoiceRepository interface {
	Insert(string) error
	Search(string) error
	Close()
}

type invoiceRepository struct {
	DB *sql.DB
}

func NewRepository() (InvoiceRepository, error) {

	conn, err := getConn()
	if err != nil {
		return nil, err
	}

	return &invoiceRepository{DB: conn}, nil
}

func getConn() (*sql.DB, error) {

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil

}

func (r *invoiceRepository) Close() {
	r.DB.Close()
}

func (r *invoiceRepository) Insert(message string) error {

	query := `insert into invoice (message, created_at) values ($1,$2)`

	fmt.Printf("inserting...")
	_, err := r.DB.Exec(query, message, time.Now())
	if err != nil {
		return err
	}

	fmt.Println("Row inserted successfuly")

	return nil
}

func (r *invoiceRepository) Search(message string) error {

	return nil
}
