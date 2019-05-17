package api

import (
	"../packages/platform"
	"encoding/json"
	"fmt"
	"net/http"
)

func PlatformHandler(w http.ResponseWriter, r *http.Request) {
	buf, _ := json.Marshal(platform.GetInstalledVersions())
	_, _ = fmt.Fprintf(w, "%s", buf)
}
