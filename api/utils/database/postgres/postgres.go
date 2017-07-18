package postgres

import (
	"database/sql"
	"errors"

	"sync"

	"strings"

	"github.com/JoaoEymard/ingressoscariri/api/utils/settings"
	_ "github.com/lib/pq"
)

var (
	postgres *sql.DB
	mutex    sync.Mutex
)

type Host struct {
	ID     string
	ExamID string
	Path   string
	IP     string
	Config string
}

// Open realizar a conexão com o banco de dados postgres.
func Open() error {
	var (
		err error
	)

	mutex.Lock()
	criptPass := settings.GetSettings().Database.ConnectionRw.Pass
	postgres, err = sql.Open("postgres", "host="+settings.GetSettings().Database.ConnectionRw.Host+" user="+settings.GetSettings().Database.ConnectionRw.User+" password="+string(criptPass[0:2])+string(criptPass[len(criptPass)-2:])+" dbname="+settings.GetSettings().Database.ConnectionRw.Database+" sslmode=disable")
	mutex.Unlock()

	if err != nil {
		return err
	}
	if err := postgres.Ping(); err != nil {
		return err
	}
	postgres.SetMaxOpenConns(settings.GetSettings().Database.ConnectionRw.MaxOpenConn)
	return nil
}

// Close t
func Close() {
	// conexao.Commit()
	postgres.Close()
}

// ExecuteQuery recebir query para serem executadas no banco postgres
func ExecuteQuery(query string) error {
	res, err := postgres.Exec(query)
	if err != nil {
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

// Select Select Universal para o Banco Postgres
func Select(attributes []string, table string, join, where, order []string, limit string) (*sql.Rows, error) {

	var query []string

	query = []string{"SELECT"}

	if attributes != nil {
		query = append(query, strings.Join(attributes, ","))
	} else {
		query = append(query, "*")
	}

	query = append(query, "FROM")

	if table != "" {
		query = append(query, table)
	} else {
		return nil, errors.New("Parametro do nome da tabela está vazio")
	}

	if join != nil {
		query = append(query, strings.Join(join, " "))
	}

	if where != nil {
		query = append(query, "WHERE")
		query = append(query, strings.Join(where, " "))
	}

	if order != nil {
		query = append(query, "ORDER BY")
		query = append(query, strings.Join(order, " "))
	}

	if limit != "" {
		query = append(query, "LIMIT")
		query = append(query, limit)
	}

	query = append(query, ";")

	stmt, err := postgres.Prepare(strings.Join(query, " "))
	if err != nil {
		return nil, err
	}

	return stmt.Query()
}

// Update Update Universal para o Banco Postgres
func Update() {

}

// Insert Insert Universal para o Banco Postgres
func Insert() {

}

// Delete Delete Universal para o Banco Postgres
func Delete() {

}
