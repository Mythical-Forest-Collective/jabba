package command

import (
	"github.com/Mythical-Forest-Collective/jabba/cfg"
	"path/filepath"
	"runtime"
)

func Which(selector string, home bool) (string, error) {
	aliasValue := GetAlias(selector)
	if aliasValue != "" {
		selector = aliasValue
	}
	ver, err := LsBestMatch(selector)
	if err != nil {
		return "", err
	}
	path := filepath.Join(cfg.Dir(), "jdk", ver)
	if home && runtime.GOOS == "darwin" {
		path = filepath.Join(path, "Contents", "Home")
	}
	return path, nil
}