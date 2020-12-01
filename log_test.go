package log_test

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path"
	"testing"
	"time"

	"github.com/masterZSH/log"
	"github.com/stretchr/testify/assert"
)

var testBaseDir = "l"
var testMsg = "foo"

func TestWrite(t *testing.T) {
	err := os.RemoveAll("l")
	assert.Nil(t, err)
	l := log.New("l")
	l.Info(testMsg)

	now := time.Now()
	d := now.Format("2006-01-02")
	h := now.Format("15")
	data, err := ioutil.ReadFile(path.Join(testBaseDir, d, h+".log"))
	assert.Nil(t, err)
	var m map[string]interface{}
	json.Unmarshal(data, &m)

	c, exists := m["level"]
	assert.NotNil(t, exists)
	assert.Equal(t, "info", c)

	msg, exists := m["msg"]
	assert.NotNil(t, exists)
	assert.Equal(t, testMsg, msg)

	// l.Error("1235")
	os.RemoveAll("l")
}
