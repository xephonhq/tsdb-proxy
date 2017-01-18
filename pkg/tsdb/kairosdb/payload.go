package kairosdb

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/xephonhq/tsdb-proxy/pkg/common"
)

// KairosDBPayload is NOT thread safe
type KairosDBPayload struct {
	firstPoint           bool
	end                  bool
	buffer               *bytes.Buffer
	bufferedSeries       []*common.Series
	bufferedIntPoints    []*common.IntPoint
	bufferedDoublePoints []*common.DoublePoint
}

func NewKairosDBPayload() *KairosDBPayload {
	p := KairosDBPayload{}
	p.buffer = bytes.NewBufferString("[")
	p.firstPoint = true
	p.end = false
	return &p
}

// AddIntPoint turns point into bytes without any grouping
func (p *KairosDBPayload) AddIntPoint(sp *common.IntPointWithSeries) {
	if !p.firstPoint {
		p.buffer.WriteString(",{")
	} else {
		p.buffer.WriteString("{")
		p.firstPoint = false
	}
	p.buffer.WriteString(fmt.Sprintf("\"name\":\"%s\",", sp.Name))
	p.buffer.WriteString(fmt.Sprintf("\"datapoints\":[[%d, %d]],\"tags\":", sp.TimeNano, sp.V))
	t, _ := json.Marshal(sp.Tags)
	p.buffer.Write(t)
	p.buffer.WriteString("}")
}

func (p *KairosDBPayload) AddPointToBuffer() {
	// this store the struct and merge into one series when get the string ([]byte) , actually it's a group by
}

func (p *KairosDBPayload) groupBySeries() {

}

func (p *KairosDBPayload) DataSize() int {
	// the real data size,
	// TODO: count series data several times?
	// TODO: the payload size, they are all different
	return 0
}

func (p *KairosDBPayload) Bytes() ([]byte, error) {
	if !p.end {
		p.buffer.WriteString("]")
		p.end = true
	}
	return p.buffer.Bytes(), nil
}
