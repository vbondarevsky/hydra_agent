package platform

type InstalledVersion struct {
	Name     string `json:"name"`
	Version  string `json:"version"`
	Location string `json:"location"`
	Size     int    `json:"size"`
}
