package main

import (
	"database/sql"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-test/deep"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	datamanager "github.com/whitejokeer/clevertitest/database"
)

type Suite struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock

	dataManage datamanager.DataManager
	beerItems  *datamanager.BeerItems
}

func (s *Suite) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)

	s.DB, err = gorm.Open("postgres", db)
	require.NoError(s.T(), err)

	s.DB.LogMode(true)

	s.dataManage = datamanager.BeerItemManager(s.DB)
}

func (s *Suite) AfterTest(_, _ string) {
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func TestInit(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (s *Suite) Test_DB_GetByID() {
	var (
		id       = 1
		name     = "Golden"
		brewery  = "Kross"
		country  = "Chile"
		price    = 10.5
		currency = "EUR"
	)

	s.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "beer_items" WHERE (id = $1)`)).
		WithArgs(id).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "brewey", "country", "price", "currency"}).
			AddRow(id, name, brewery, country, price, currency))

	res, err := s.dataManage.GetBeerItemByID(id)

	require.NoError(s.T(), err)
	require.Nil(s.T(), deep.Equal(&datamanager.BeerItems{
		Id:       id,
		Name:     name,
		Brewery:  brewery,
		Country:  country,
		Price:    price,
		Currency: currency,
	}, res))
}
