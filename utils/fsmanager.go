package utils

import (
	"fmt"
	"os"
)

// ResExist check is file or dir exists on system
func ResExist(inRes string) (bool, error) {
	if _, err := os.Stat(inRes); os.IsNotExist(err) {
		return false, err
	}
	return true, nil
}

// FullPath builds global path, from relative
func FullPath(partialPath string) string {
	return fmt.Sprintf("%v%v", ProjectRoot, partialPath)
}
