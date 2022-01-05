package version

import (
	"fmt"
	"runtime"
)

var (
	Version     = "0.4.0"
	BuildNumber = "27149"
	GitCommit   = ""
	gitTreeState = ""
	buildDate = ""
)

type info struct {
	GitVersion   string `json:"gitVersion"`
	GitCommit    string `json:"gitCommit"`
	GitTreeState string `json:"gitTreeState"`
	BuildDate    string `json:"buildDate"`
	GoVersion    string `json:"goVersion"`
	Compiler     string `json:"compiler"`
	Platform     string `json:"platform"`
}

func Get() info {
	return info{
		GitVersion:   Version,
		GitCommit:    GitCommit,
		GitTreeState: gitTreeState,
		BuildDate:    buildDate,
		GoVersion:    runtime.Version(),
		Compiler:     runtime.Compiler,
		Platform:     fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
	}
}
