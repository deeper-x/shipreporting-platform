package utils

import (
	"fmt"
	"os"
)

// ProjectName global name
var ProjectName = "shipreporting-platform"

// ProjectRoot base directory
var ProjectRoot = fmt.Sprintf("%s/src/github.com/deeper-x/%s", os.Getenv("GOPATH"), ProjectName)

// TemplateAssets stuff registration
var TemplateAssets = fmt.Sprintf("%v/templates/layout/theme/classic/assets/", ProjectRoot)

// GridAssets grid data tables
var GridAssets = fmt.Sprintf("%v/grid/", ProjectRoot)

// DotenvFile .env location
var DotenvFile = fmt.Sprintf("%v/.env", ProjectRoot)
