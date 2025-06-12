package updatechecker_test

import (
	"fmt"
	"strings"
	"testing"

	// updatechecker "github.com/amarillier/go-update-checker"
	updatechecker "github.com/amarillier/go-update-checker"
	// "tanium.local/amarillier/client"
)

func TestCurrent(t *testing.T) {
	uc := updatechecker.New("amarillier", "go-update-checker", "Go Update Checker", "https://github.com/amarillier/go-update-checker/releases", 0, false)
	uc.CheckForUpdate("0.0.3")
	fmt.Println(uc.Message)
	if uc.UpdateAvailable {
		t.Fatal("Update should not be available for current latest version 0.0.3")
	}
	if !strings.Contains(uc.Message, "You are running the latest version") {
		t.Fatal("Update should not be available for current latest version 0.0.3")
	}
}

func TestNewer(t *testing.T) {
	uc := updatechecker.New("amarillier", "go-update-checker", "Go Update Checker", "https://github.com/amarillier/go-update-checker/releases", 0, false)
	uc.CheckForUpdate("9.0.1")
	fmt.Println(uc.Message)
	if uc.UpdateAvailable {
		t.Fatal("Update should not be available when testing for a newer version")
	}
	if !strings.Contains(uc.Message, "You are running a newer version") {
		t.Fatal("Update message should indicate a newer version is running")
	}
}

func TestOlder(t *testing.T) {
	uc := updatechecker.New("amarillier", "go-update-checker", "Go Update Checker", "https://github.com/amarillier/go-update-checker/releases", 0, false)
	uc.CheckForUpdate("0.0.2")
	fmt.Println(uc.Message)
	if !uc.UpdateAvailable {
		t.Fatal("Update should not be available when testing for current latest version 0.0.3")
	}
	if !strings.Contains(uc.Message, "A new update is available") {
		t.Fatal("Update message should indicate a new update is available")
	}
}
