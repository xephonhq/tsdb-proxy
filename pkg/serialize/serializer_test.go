package serialize

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/xephonhq/tsdb-proxy/pkg/common"

	"time"
)

// test implementation satisfies the interface
func TestSerializerInterface(t *testing.T) {
	t.Parallel()
	var _ Serializer = (*DebugSerializer)(nil)
	var _ Serializer = (*JsonSerializer)(nil)
}

type SerializeTestSuite struct {
	suite.Suite
	iP *common.IntPointWithSeries
	dP *common.DoublePointWithSeries
	ts int64
}

func TestSerializeTestSuite(t *testing.T) {
	suite.Run(t, new(SerializeTestSuite))
}

func (suite *SerializeTestSuite) SetupTest() {
	name := "cpu.idle"
	s := common.NewSeries(name)
	s.AddTag("os", "ubuntu")
	s.AddTag("arch", "amd64")
	ts := time.Now().UnixNano()
	suite.ts = ts
	suite.iP = &common.IntPointWithSeries{Series: s}
	suite.iP.V = 123
	suite.iP.TimeNano = ts
	suite.dP = &common.DoublePointWithSeries{Series: s}
	suite.dP.V = 12.03
	suite.dP.TimeNano = ts
}

func (suite *SerializeTestSuite) TestDebugSerializer() {
	assert := assert.New(suite.T())
	ds := DebugSerializer{}

	w, _ := ds.WriteInt(suite.iP)
	s := string(w)
	assert.Contains(s, "cpu.idle", "123", "os=ubuntu", "arch=amd64")

	w, _ = ds.WriteDouble(suite.dP)
	s = string(w)
	assert.Contains(s, "cpu.idle", "12.03", "os=ubuntu", "arch=amd64")
}

func (suite *SerializeTestSuite) TestJsonSerializer() {
	assert := assert.New(suite.T())
	js := JsonSerializer{}
	w, err := js.WriteInt(suite.iP)
	//o := fmt.Sprintf("{\"v\":123,\"t\":%d,\"name\":\"cpu.idle\",\"tag\":{\"arch\":\"amd64\",\"os\":\"ubuntu\"}}", suite.ts )
	assert.Nil(err)
	suite.T().Log(string(w))
	//assert.Equal(o, string(w))
	// NOTE: \n has no effect on json deserialization
	dI, err := js.ReadInt(w)
	assert.Nil(err)
	suite.T().Log(dI)
	suite.T().Log(dI.V)
	suite.T().Log(dI.TimeNano)

	w, err = js.WriteDouble(suite.dP)
	//o = fmt.Sprintf("{\"v\":12.03,\"t\":%d,\"name\":\"cpu.idle\",\"tag\":{\"arch\":\"amd64\",\"os\":\"ubuntu\"}}", suite.ts)
	assert.Nil(err)
	suite.T().Log(string(w))
	//assert.Equal(o, string(w))
	dd, err := js.ReadDouble(w)
	assert.Nil(err)
	suite.T().Log(dd)
	suite.T().Log(dd.V)
}
