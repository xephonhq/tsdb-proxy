package test

import (
	"os"

	st "github.com/dyweb/Ayi/common/structure"
	"github.com/xephonhq/xephon-b/pkg/tsdb/common"
	"github.com/xephonhq/xephon-b/pkg/tsdb/config"
	"github.com/xephonhq/xephon-b/pkg/tsdb/kairosdb"
)

var pinged st.Set
var running st.Set

// KairosDB determines if we should run test that requires KairosDB up and running
// - if the environment variable is set, we test
// TODO: this may mess up the running database, but sicne we use docker, the data can lost
// - if we can ping the database using provided config, we test
func KairosDB() bool {
	if running.Contains(common.KairosDB) {
		return true
	}
	// we pinged and it is not running
	if pinged.Contains(common.KairosDB) {
		return false
	}
	// env var goes before ping, pinged is empty at first, so env var is triggered first
	// TODO: get environment variable name from the `common` package instead hardcoded here
	if os.Getenv("TEST_KAIROSDB") == "1" {
		running.Add(common.KairosDB)
		return true
	}
	pinged.Add(common.KairosDB)
	// TODO: allow get test config from config file instead of just using default
	h, err := config.NewDefaultHost(common.KairosDB)
	if err != nil {
		log.Warn(err)
		return false
	}
	c := config.TSDBClientConfig{Host: h}
	client := &kairosdb.KairosDBHTTPClient{Config: c}
	err = client.Ping()
	if err != nil {
		log.Warn(err)
		return false
	}
	running.Add(common.KairosDB)
	return true
}
