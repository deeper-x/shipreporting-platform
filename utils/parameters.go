package utils

import (
	"fmt"
	"os"
)

// ProjectName global name
var ProjectName = "shipreporting-platform"

// ProjectRoot base directory
var ProjectRoot = fmt.Sprintf("%s/src/github.com/deeper-x/%s", os.Getenv("GOPATH"), ProjectName)

// EnvFile .env location
var EnvFile = fmt.Sprintf("%s/.env", ProjectRoot)
