package version

import (
	"fmt"
	"testing"
)

func TestGetVersion(t *testing.T) {
	fmt.Println("#### version is ", GetVersion())
}
