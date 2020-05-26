// +build !release
//go:generate go run github.com/UnnoTed/fileb0x b0x.yaml

package supervisord

import (
	"net/http"
)

//HTTP auto generated
var HTTP http.FileSystem = http.Dir("./webgui")
