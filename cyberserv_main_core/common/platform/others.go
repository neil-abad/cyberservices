// +build !windows

package platform

import (
	"os"
	"path/filepath"
)

func ExpandEnv(s string) string {
	return os.ExpandEnv(s)
}

func LineSeparator() string {
	return "\n"
}

func GetToolLocation(file string) string {
	const name = "cyberservices.location.tool"
	toolPath := EnvFlag{Name: name, AltName: NormalizeEnvName(name)}.GetValue(getExecutableDir)
	return filepath.Join(toolPath, file)
}

// GetAssetLocation search for `file` in certain locations
func GetAssetLocation(file string) string {
	const name = "cyberservices.location.asset"
	assetPath := NewEnvFlag(name).GetValue(getExecutableDir)
	defPath := filepath.Join(assetPath, file)
	for _, p := range []string{
		defPath,
		filepath.Join("/usr/local/share/Project CS/", file),
		filepath.Join("/usr/share/Project CS/", file),
	} {
		if _, err := os.Stat(p); os.IsNotExist(err) {
			continue
		}

		// asset found
		return p
	}

	// asset not found, let the caller throw out the error
	return defPath
}
