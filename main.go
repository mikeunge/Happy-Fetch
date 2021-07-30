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
const MotdPath = "/tmp/hpy.json"
const CachePath = "/tmp/.hpy"

var config u.Config

func getHome() string {
	homedir, err := os.UserHomeDir()
	if err != nil {
		u.Error(config, err)
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
		u.Warn(config, "config file does not exist")
		config = u.ConfigDefault()
	} else {
		// Read the configfile.
		config, err = u.ConfigReader(f)
		if err != nil {
			u.Error(config, err)
			fmt.Printf("ERROR: %+v\n", err)
			os.Exit(-1)
		}
	}
	return config
}

func writeCacheAndMotd(c u.Config, quote q.Quote, quotes []q.Quote) {
	if q.CheckMotd(c.Motd, MotdPath) {
		u.Debug(config, "writing motd")
		err := q.WriteMotd(MotdPath, quote)
		if err != nil {
			log.Printf("ERROR: %+v\n", err)
		}
	}
	if q.CheckCache(c.Cache, c.CachePath, c.CacheRefresh) {
		u.Debug(config, "writing cache")
		err := q.WriteCache(c.CachePath, quotes)
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
		quote, err = q.GetMotd(MotdPath)
		if err != nil {
			u.Error(config, err)
			fmt.Printf("ERROR: %+v\n", err)
			os.Exit(1)
		}
	case "cache":
		quotes, err := q.GetCache(CachePath)
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
	if config.Logging {
		u.StartLogger()
	}
}

func main() {
	var quote q.Quote
	var method string

	// Get the users homedir and construct the fullConfigPath.
	fullConfigPath := getHome()
	config = loadConfig(fullConfigPath)

	// Define what to do/where to get the data from
	if q.CheckMotd(config.Motd, MotdPath) {
		method = "motd"
	} else if q.CheckCache(config.Cache, config.CachePath, config.CacheRefresh) {
		method = "cache"
	} else if q.CheckCache(config.Cache, CachePath, config.CacheRefresh) {
		// check a different path, if works, set the config.CachePath
		config.CachePath = CachePath
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
