# Version Checker for Go

This is a fork of the original go-update-checker by Chris (Christian1984)

go-update-checker is a go library for checking the version of a currently installed application or package against its latest release on github. It also enables caching and setting a minimum interval of days after which a updatecheck against the github API should be performed to prevent spamming the API.

Versions used with go-update-checker must follow [SemVer](http://semver.org/).

## Installation and Usage

Installation can be done with a normal `go get`:

```
$ go get github.com/amarillier/go-update-checker
$ go get github.com/amarillier/go-update-checker@latest
```


## Issues and Contributing

If you find an issue with this library, please report an issue. If you'd like, we welcome any contributions. Fork this library and submit a pull request.

Again, this is a fork of the original by Chris with some minor updates to suit my preferences, correct the sample code, and return slightly different output:
https://github.com/Christian1984/go-update-checker


#### Update Check Example1

```go
import (
	// "fmt"

	"fmt"
	// updatechecker "github.com/Christian1984/go-update-checker"
	updatechecker "github.com/amarillier/go-update-checker"
)

func main() {
	uc := updatechecker.New("amarillier", "go-update-checker", "Go Update Checker", "https://github.com/amarillier/go-update-checker/releases", 0, false)

	fmt.Println("Checking for 0.0.1")
	uc.CheckForUpdate("0.0.1")
	uc.PrintMessage()
	/* alternatively use uc.Message (type string) in any other context */
	// fmt.Println(uc.Message)
	fmt.Println("Update available: ", uc.UpdateAvailable)
	fmt.Println("-------------")
	dl := uc.DownloadLink
	fmt.Println(dl)
}
```

#### Example output 1
```Checking for 0.0.1
=======================================================
=== A new update is available for Go Update Checker ===
=======================================================
Your Version: v0.0.1

New Version: v0.0.3
Title: v0.0.3
Description:
v0.0.3
Added isCurrentVersionNewer capability

**Full Changelog**: https://github.com/amarillier/go-update-checker/compare/v0.0.2...v0.0.3

Download the latest version here:
https://github.com/amarillier/go-update-checker/releases
=======================================================
Update available:  true
-------------
https://github.com/amarillier/go-update-checker/releases
```

#### Update Check Example2

```go
import (
	// "fmt"

	"fmt"
	// updatechecker "github.com/Christian1984/go-update-checker"
	updatechecker "github.com/amarillier/go-update-checker"
)

func main() {
	uc := updatechecker.New("amarillier", "go-update-checker", "Go Update Checker", "https://github.com/amarillier/go-update-checker/releases", 0, false)

	fmt.Println("Checking for 0.0.3")
	uc.CheckForUpdate("0.0.3")
	uc.PrintMessage()
	/* alternatively use uc.Message (type string) in any other context */
	// fmt.Println(uc.Message)
	fmt.Println("Update available: ", uc.UpdateAvailable)
	fmt.Println("-------------")
	dl := uc.DownloadLink
	fmt.Println(dl)
}
```

#### Example output 2
```
Checking for 0.0.3
===============================================================
=== You are running the latest version of Go Update Checker ===
===============================================================
Update available:  false
-------------
https://github.com/amarillier/go-update-checker/releases
```

#### Update Check Example3

```go
import (
	// "fmt"

	"fmt"
	// updatechecker "github.com/Christian1984/go-update-checker"
	updatechecker "github.com/amarillier/go-update-checker"
)

func main() {
	uc := updatechecker.New("amarillier", "go-update-checker", "Go Update Checker", "https://github.com/amarillier/go-update-checker/releases", 0, false)

	fmt.Println("Checking for 9.0.3")
	uc.CheckForUpdate("9.0.3")
	uc.PrintMessage()
	/* alternatively use uc.Message (type string) in any other context */
	// fmt.Println(uc.Message)
	fmt.Println("Update available: ", uc.UpdateAvailable)
	fmt.Println("-------------")
	dl := uc.DownloadLink
	fmt.Println(dl)
}
```

#### Example output 3
```
Checking for 9.0.3
==================================================================
=== You are running a newer version of Go Update Checker 9.0.3 ===
==================================================================
Update available:  false
-------------
https://github.com/amarillier/go-update-checker/releases
```

