package main

// YEAH YEAH CARGO CULT PROGRAMMING
// *HACK* *HACK* *COPY* *HACK*

import (
	"flag"
	"fmt"
	"time"
	"net/http"
	"log"
	"encoding/json"
	"os/exec"
)

type Clubinfo struct {
	Last_event string
	Club_offen bool
}

func main() {
	// WAAAAAH!
	var url = flag.String("url", "http://club.entropia.de", "Target URL for JSON data")
	var timeout = flag.Int("timeout", 1, "Timeout in hours")
	var dryrun = flag.Bool("dryrun", false, "Dry run")

	flag.Parse()

	resp, err := http.Get(*url)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	// WAAAAAAAAHAAAAAA!
	decoder := json.NewDecoder(resp.Body)

	var c Clubinfo
	err = decoder.Decode(&c)

	if err != nil {
		log.Fatal(err)
	}

	// RFC3339 without ':' for time zone foooooo
	last, _ := time.Parse("2006-01-02T15:04:05Z0700", c.Last_event)

	// WAAAAABWHARGLGLRELRLLLllmn...
	if time.Since(last).Hours() > float64(*timeout) {
		if !c.Club_offen {
			if !*dryrun {
				cmd := exec.Command("shutdown", "-h", "+5")
				cmd.Start()
			} else {
				fmt.Println("I advise against letting this machine run any longer.");
			}
		}
	}
}

