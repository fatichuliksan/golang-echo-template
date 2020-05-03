package postgre

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// postgreHelper ...
type postgreHelper struct {
	Username string
	Password string
	Host     string
	Port     int
	Name     string
}

// Method ...
type Method interface {
	Connect() (*gorm.DB, error)
}

// NewPostgre ...
func NewPostgre(username string, password string, host string, port int, name string) Method {
	return &postgreHelper{
		Username: username,
		Password: password,
		Host:     host,
		Port:     port,
		Name:     name,
	}
}

// Connect ...
func (t *postgreHelper) Connect() (*gorm.DB, error) {
	db, err := gorm.Open("postgres", fmt.Sprintf("host=%s port=%v user=%s dbname=%s password=%s sslmode=disable", t.Host, t.Port, t.Username, t.Name, t.Password))
	if err != nil {
		panic(err)
	}

	fmt.Println("DB '" + t.Name + "' Connected!")
	return db, err
}
