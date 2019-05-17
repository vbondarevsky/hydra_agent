package platform

import (
	"golang.org/x/sys/windows/registry"
)

func isVendor1C(vendor string) bool {
	return vendor == "1С-Софт" || vendor == "1C-Soft" || vendor == "1C" || vendor == "1С"
}

func GetInstalledVersions() []InstalledVersion {
	result := make([]InstalledVersion, 0)

	key := `SOFTWARE\Microsoft\Windows\CurrentVersion\Uninstall`
	keys := make([]registry.Key, 2)

	if key, err := registry.OpenKey(registry.LOCAL_MACHINE, key, registry.ENUMERATE_SUB_KEYS|registry.WOW64_32KEY); err == nil {
		keys = append(keys, key)
		defer key.Close()
	}
	if key, err := registry.OpenKey(registry.LOCAL_MACHINE, key, registry.ENUMERATE_SUB_KEYS|registry.WOW64_64KEY); err == nil {
		keys = append(keys, key)
		defer key.Close()
	}

	for _, key := range keys {
		subKeys, _ := key.ReadSubKeyNames(-1)
		for _, subKey := range subKeys {
			subK, _ := registry.OpenKey(key, subKey, registry.READ)

			vendor, _, _ := subK.GetStringValue("Publisher")

			if !isVendor1C(vendor) {
				continue
			}
			name, _, _ := subK.GetStringValue("DisplayName")
			version, _, _ := subK.GetStringValue("DisplayVersion")
			location, _, _ := subK.GetStringValue("InstallLocation")

			result = append(result, InstalledVersion{name, version, location, 0})
		}
	}

	return result
}
