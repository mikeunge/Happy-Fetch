package main

import (
	"fmt"
	"log"
	"os"

	q "github.com/mikeunge/Happy-Fetch/quote"
	u "github.com/mikeunge/Happy-Fetch/utils"
)

const ApiEndpoint = "https://type.fit/api/quotes"
const ConfigPath = "/.config/happy_fetch/init.toml"

// const MotdPath = "/tmp/hpy.json"
// const CachePath = "/tmp/.hpy"

var config u.Config

func getHome() string {
	homedir, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("ERROR: %+v\n", err)
		os.Exit(1)
	}
	return homedir + ConfigPath
}

func loadConfig(f string) u.Config {
	var config u.Config
	var err error

	ok := u.FileExists(f)
	if !ok {
		// Get the default configuration.
		fmt.Printf("Config not found, using default config.")
		config = u.ConfigDefault()
	} else {
		// Read the configfile.
		config, err = u.ConfigReader(f)
		if err != nil {
			fmt.Printf("ERROR: %+v\n", err)
			os.Exit(-1)
		}
	}
	return config
}

func writeCacheAndMotd(config u.Config, quote q.Quote, quotes []q.Quote) {
	if q.CheckMotd(config, true) {
		u.Debug(config, "writing motd")
		err := q.WriteMotd(config.MotdPath, quote)
		if err != nil {
			log.Printf("ERROR: %+v\n", err)
		}
	}
	if q.CheckCache(config, true) {
		u.Debug(config, "writing cache")
		err := q.WriteCache(config.CachePath, quotes)
		if err != nil {
			u.Error(config, err)
		}
	}
}

func do(m string, c u.Config) q.Quote {
	var quote q.Quote
	var quotes []q.Quote
	var err error

	switch m {
	case "motd":
		// load motd
		quote, err = q.GetMotd(config.MotdPath)
		if err != nil {
			u.Error(config, err)
			fmt.Printf("ERROR: %+v\n", err)
			os.Exit(1)
		}
	case "cache":
		quotes, err := q.GetCache(config.CachePath)
		if err != nil {
			u.Error(config, err)
			fmt.Printf("ERROR: %+v", err)
			os.Exit(1)
		}
		quote = q.GetQuoteFromQuotes(quotes)
	case "api":
		// Fetch the quotes from the server.
		quotes, err = q.GetQuotesFromApi(ApiEndpoint)
		if err != nil {
			u.Error(config, err)
			fmt.Printf("ERROR: %+v", err)
			os.Exit(1)
		}
		quote = q.GetQuoteFromQuotes(quotes)
	default:
		// this case cannot happen, but yeah, just exit I guess.
		os.Exit(125)
	}

	writeCacheAndMotd(c, quote, quotes)
	return quote
}

func init() {
	// Get the users homedir and construct the fullConfigPath.
	fullConfigPath := getHome()
	config = loadConfig(fullConfigPath)

	// check if we should start/create a logger object
	if config.Logging {
		ok := u.StartLogger(config.LoggingPath)
		// make sure the logger was created, if not, check if we should contiue or not
		if !ok && !config.IgnoreLogging {
			log.Fatalf("error opening file: %s", config.LoggingPath)
			os.Exit(-1)
		}
	}
}

func main() {
	var quote q.Quote
	var method string

	// Define what to do/where to get the data from
	if q.CheckMotd(config, false) {
		method = "motd"
	} else if q.CheckCache(config, false) {
		method = "cache"
	} else {
		method = "api"
	}
	u.Debug(config, fmt.Sprintf("method -> %s", method))

	quote = do(method, config)
	// q.WriteMotd(MotdPath, quote)
	msg, err := q.FormatQuote(quote)
	if err != nil {
		u.Error(config, err)
		fmt.Printf("ERROR: %+v\n", err)
		os.Exit(1)
	}

	fmt.Println(msg)
	os.Exit(0)
}
