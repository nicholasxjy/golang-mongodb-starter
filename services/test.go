package services

import (
	"edenedu/conf"
	"edenedu/models"

	mgo "gopkg.in/mgo.v2"
)

// TestService service
type TestService struct {
	DBSession *mgo.Session
}

// InsertTests insert some test data
func (s *TestService) InsertTests(tests []string) error {
	var data []interface{}
	session := s.DBSession.Copy()
	defer session.Close()
	c := session.DB(conf.DBName).C("test")
	for _, val := range tests {
		data = append(data, &models.Test{Foo: val})
	}
	return c.Insert(data...)
}

// FindTests find all test
func (s *TestService) FindTests() ([]models.Test, error) {
	var tests = []models.Test{}
	session := s.DBSession.Copy()
	defer session.Close()
	c := session.DB(conf.DBName).C("test")
	err := c.Find(nil).All(&tests)
	return tests, err
}
