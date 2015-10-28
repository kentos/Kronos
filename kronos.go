package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func runner(url string) {
	response, err := http.Get(url)
	if err != nil {
		log.Println("Error setting up http.Get")
	}

	if response.StatusCode != 200 {
		log.Println("Response with wrong status code.", response.StatusCode, url)
	}
}

/*
*	Defined Kronstamps
*	10 second, 30 second, minute, 5 minute, 15 minute, hourly, 2 hour, daily (00:00)
 */

func shouldRun(t string) bool {
	s := time.Now().Second()
	m := time.Now().Minute()
	h := time.Now().Hour()

	if t == "minute" {
		if s == 0 {
			return true
		}
	} else if t == "hourly" {
		if s == 0 && m == 0 {
			return true
		}
	} else if t == "2 hour" {
		if s == 0 && h%2 == 0 {
			return true
		}
	} else if t == "daily" {
		if s == 0 && m == 0 && h == 0 {
			return true
		}
	} else if t == "10 second" {
		if s%10 == 0 {
			return true
		}
	} else if t == "30 second" {
		if s == 30 || s == 0 {
			return true
		}
	} else if t == "5 minute" {
		if s == 0 && m%5 == 0 {
			return true
		}
	} else if t == "15 minute" {
		if s == 0 && m%15 == 0 {
			return true
		}
	}

	return false
}

func main() {
	log.Println("Starting Kronos")

	cron_file := "./crons.js"

	// Should we sync to closests minute
	sec := time.Now().Second()

	if sec != 0 {
		res := 60 - sec
		log.Printf("Waiting for closets minute for %d seconds", res)

		for {
			if time.Now().Second() != 0 {
				time.Sleep(time.Millisecond * 500)
				continue
			} else {
				log.Println("A minute! Lest goo!")
				break
			}
		}
	}

	for {
		// Load feeds on loop start
		f, err := ioutil.ReadFile(cron_file)
		if err != nil {
			log.Printf("Could not read %q, waiting 30 seconds to re-try", cron_file)
			time.Sleep(time.Second * 30)
			continue
		}

		// Parse log

		// Unmarshal json
		cmds := []Kron{}

		err = json.Unmarshal(f, &cmds)
		if err != nil {
			log.Printf("Could not unmarshal %q", cron_file)
		}

		// Loop through kron cmds
		for x := range cmds {
			if shouldRun(cmds[x].When) {
				go runner(cmds[x].Url)
			}
		}

		// Waiting X seconds
		time.Sleep(time.Second * 10)
	}
}

type Kron struct {
	When string `json:"when"`
	Url  string `json:"url"`
}
