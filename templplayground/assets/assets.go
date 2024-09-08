package assets

import (
	"crypto/md5"
	"embed"
	"encoding/base32"
	"io/fs"
	"path"
	"path/filepath"
	"strings"
)

//go:embed css/* js/*
var assets embed.FS

var Assets = WrapFS(Config{Root: "/static"})

type FS struct {
	cfg    Config
	hashes map[string]string
	fs.FS
}

type Config struct {
	Root string
}

func WrapFS(cfg Config) *FS {
	fs := &FS{
		cfg:    cfg,
		FS:     assets,
		hashes: make(map[string]string),
	}
	err := fs.generateHashes(".")
	if err != nil {
		panic(err)
	}
	return fs
}

func (f *FS) generateHashes(p string) error {
	des, err := assets.ReadDir(p)
	if err != nil {
		return err
	}
	for _, de := range des {
		p := filepath.Join(p, de.Name())
		switch {
		case strings.HasPrefix(de.Name(), "."):
		case de.IsDir():
			f.generateHashes(p)
		default:
			content, err := assets.ReadFile(p)
			if err != nil {
				return err
			}
			hash := md5.Sum(content)
			f.hashes[path.Join(f.cfg.Root, p)] = strings.Trim(strings.ToLower(base32.StdEncoding.EncodeToString(hash[:])), "=")
		}
	}
	return nil
}

func (f FS) URI(asset string) string {
	return asset + "?" + f.hashes[asset]
}
