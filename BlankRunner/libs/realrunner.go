package libs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

type RealRunner struct {
	UniqueId   string
	Jsonstring string
}

func (this *RealRunner) QueueName() string {
	return "download_RealQueue"
}

func (this *RealRunner) QueueSetName() string {
	return "download_checkSet"
}

func (this *RealRunner) Init(jsonstring string, uniqueId string) {
	this.UniqueId = uniqueId
	this.Jsonstring = jsonstring
}

func (this *RealRunner) RealDoHandler() (bool, string) {
	return true, ""
}
