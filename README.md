# Kronos
A simple program for running cron-like url get task

#### The purpose
_Kronos_, the youngest of the greek titans, is a simple program written in Go for handling time-based url tasks.

#### Setup
* cd /your/go/src/kronos
* go install
* on ubuntu do: "sudo ./your/go/bin/kronos > /dev/null &" for a simple daemon. Else set up as you want to.

#### How it runs
_Kronos_ runs every 10 second. This means it's the lowest time period an url can be GET'd. _Kronos_ currently supports GET requests and logging if response status code is not 200 (it's enough for me atm).

Built in params: **10 second**, **30 second**, **minute**, **5 minute**, **15 minute**, **hourly**, **2 hour**, **daily _(at 00:00)_**.

Just setup a json file (./crons.js) which contains an array with json objects with the attributes "when" and "url".

    [
    	{
    		"when": "minute",
    		"url": "http://example.com/sync"
    	},
    	{
    		"when": "hourly",
    		"url": "http://example.com/sync_hourly"
    	}
    ]
