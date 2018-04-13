package main

import (
	"fmt"
	"github.com/blang/semver"
	"github.com/rhysd/go-github-selfupdate/selfupdate"
	"os"
)

const version = "1.0.0"
const update_url = "Erachter/go-self-update"

func self_update() error {
	selfupdate.EnableLog()

	fmt.Println("Current version is:", version)

	previous := semver.MustParse(version)
	latest, err := selfupdate.UpdateSelf(previous, update_url)
	if err != nil {
		return err
	}

    fmt.Println("The latest version is:", latest.Version)
	if previous.Equals(latest.Version) {
		fmt.Println("No necessity to update.")
	} else {
		fmt.Println("Update successful.", latest.Version)
		fmt.Println("Release note:\n", latest.ReleaseNotes)
	}
	return nil
}

func main() {
	err := self_update();
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Println("Hello World")
	os.Exit(0)
}
