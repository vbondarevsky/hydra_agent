package api

import (
	"encoding/json"
	"fmt"
	"github.com/matishsiao/goInfo"
	"net/http"
)

type System struct {
	Kernel   string `json:"kernel"`
	Core     string `json:"core"`
	Platform string `json:"platform"`
	OS       string `json:"os"`
	Hostname string `json:"hostname"`
	CPUs     int    `json:"cpus"`
}

type Health struct {
	Version string `json:"version"`
	System  `json:"system"`
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	gi := goInfo.GetInfo()
	health := Health{"alpha", System{gi.Kernel, gi.Core, gi.Platform, gi.OS, gi.Hostname, gi.CPUs}}
	buf, _ := json.Marshal(health)

	_, _ = fmt.Fprintf(w, "%s", buf)
}
