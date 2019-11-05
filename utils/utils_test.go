package utils

import (
	"testing"
)

func TestResExist(t *testing.T) {
	_, err := ResExist(ProjectRoot)

	if err != nil {
		t.Error(err)
	}
}

func TestFullPath(t *testing.T) {
	res := FullPath(ProjectRoot)
	if len(res) == 0 {
		t.Errorf("%v zero length", ProjectRoot)
	}
}

func TestParameters(t *testing.T) {
	if exist, err := ResExist(ProjectRoot); !exist {
		t.Error(err)
	}

	if exist, err := ResExist(EnvFile); !exist {
		t.Error(err)
	}

	if exist, err := ResExist(StaticPath); !exist {
		t.Error(err)
	}
}
