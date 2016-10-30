package controllers

import (
	"edenedu/services"
	"net/http"

	"github.com/labstack/echo"
	mgo "gopkg.in/mgo.v2"
)

// TestController controller
type TestController struct {
	DBSession *mgo.Session
}

// TestGet /api/v1/tests
func (t *TestController) TestGet(c echo.Context) error {
	testService := &services.TestService{DBSession: t.DBSession}
	tests, err := testService.FindTests()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, tests)
}

// TestInsert /api/v1/tests/insert
func (t *TestController) TestInsert(c echo.Context) error {
	testService := &services.TestService{DBSession: t.DBSession}
	var tests = []string{"a", "b", "c", "d", "f"}
	err := testService.InsertTests(tests)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, struct{ Status string }{Status: "ok"})
}
