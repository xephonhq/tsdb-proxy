package opentsdb

import (
	"github.com/dyweb/gommon/requests"
	"github.com/pkg/errors"
	"github.com/xephonhq/tsdb-proxy/pkg/tsdb"
	"github.com/xephonhq/tsdb-proxy/pkg/tsdb/config"
	"github.com/xephonhq/tsdb-proxy/pkg/util"
)

// Short name use in OpenTSDB client package
var log = util.Logger.NewEntryWithPkg("x.tsdb.opentsdb")

type OpenTSDBHTTPClient struct {
	Config config.TSDBClientConfig
}

type OpenTSDBTelnetClient struct {
}

func (client *OpenTSDBHTTPClient) Ping() error {
	versionURL := client.Config.Host.HostURL() + "/api/version"
	res, err := requests.GetJSON(versionURL)
	if err != nil {
		return errors.Wrapf(err, "can't reach OpenTSDB via %s", versionURL)
	}
	log.Info("OpenTSDB version is " + res["version"])
	return nil
}

func (client *OpenTSDBHTTPClient) Put(p tsdb.TSDBPayload) error {
	return nil
}
