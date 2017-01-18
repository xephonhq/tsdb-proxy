// +build !race

package integration_test

import (
	"testing"

	"github.com/stretchr/testify/suite"
	tutil "github.com/xephonhq/tsdb-proxy/pkg/util/test"
)

type KairosDBTestSuite struct {
	suite.Suite
}

func TestKairosDBTestSuite(t *testing.T) {
	if tutil.KairosDB() {
		suite.Run(t, new(KairosDBTestSuite))
	}
}
