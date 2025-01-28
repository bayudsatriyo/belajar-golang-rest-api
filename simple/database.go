package simple

type Database struct {
	Name string
}

type DatabasePostgresql Database
type DatabaseMysql Database

type DatabaseRepository struct {
	DatabasePostgresql *DatabasePostgresql
	DatabaseMysql      *DatabaseMysql
}

func NewDatabaseRepository(postgresql *DatabasePostgresql, mysqlDatabase *DatabaseMysql) *DatabaseRepository {
	return &DatabaseRepository{DatabasePostgresql: postgresql, DatabaseMysql: mysqlDatabase}
}

func NewDatabaseMysql() *DatabaseMysql {
	return (*DatabaseMysql)(&Database{Name: "Mysql"})
}

func NewDatabasePostgresql() *DatabasePostgresql {
	return (*DatabasePostgresql)(&Database{Name: "Postgresql"})
}
