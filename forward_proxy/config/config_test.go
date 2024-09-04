package config_test

import (
	"os"
	"scratch/random_practice/forward_proxy/config"
	"testing"

	"github.com/reedobrien/checkers"

	"gopkg.in/yaml.v3"
)

func TestNew(t *testing.T) {
	t.Parallel()

	cfg := config.Config{
		PassRequestHeaders: []string{
			"test-req-header",
		},
		PassResponseHeaders: []string{
			"test-resp-header",
		},
	}

	b, err := yaml.Marshal(cfg)
	checkers.OK(t, err)

	td := t.TempDir()

	f, err := os.Create(td + "/" + "test.yaml")
	checkers.OK(t, err)

	_, err = f.Write(b)
	checkers.OK(t, err)

	err = f.Sync()
	checkers.OK(t, err)

	tut, err := config.New(td + "/" + "test.yaml")
	checkers.OK(t, err)

	checkers.Equals(t, tut.PassRequestHeaders[0], "test-req-header")
}

func TestNewNoFile(t *testing.T) {
	t.Parallel()

	_, err := config.New("/nada/mas")
	checkers.Equals(t, err.Error(), "error opening config : open /nada/mas: no such file or directory")
}
