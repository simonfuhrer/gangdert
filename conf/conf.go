package conf

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/golang/glog"
)

//PathSeparator sdfsd
const (
	PathSeparator     = string(os.PathSeparator)     // OS-specific path separator
	PathListSeparator = string(os.PathListSeparator) // OS-specific path list separator
)

//WideVersion d
const (
	Version = "0.0.1" // wide version
)

type conf struct {
	Listen      string // 127.0.0.1:80
	MaxProcs    int    // Go max procs
	RuntimeMode string // runtime mode (dev/prod)

}

//JSON
var GangDert conf

//Load the configurationfile
func Load() {
	bytes, _ := ioutil.ReadFile("conf/conf.json")
	err := json.Unmarshal(bytes, &GangDert)
	if err != nil {
		glog.Error(err)
		os.Exit(-1)
	}

}
