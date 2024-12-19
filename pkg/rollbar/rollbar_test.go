package rollbar_test

import (
	"testing"

	"github.com/yogamandayu/go-boilerplate/consts"
	"github.com/yogamandayu/go-boilerplate/tests"
)

func TestRollbar(t *testing.T) {

	testSuite := tests.NewTestSuite()
	defer func() {
		t.Cleanup(testSuite.Clean)
	}()
	testSuite.LoadApp()
	testSuite.App.Rollbar.Message(consts.RollbarSeverityLevelInfo.String(), "Hello World!")
}
