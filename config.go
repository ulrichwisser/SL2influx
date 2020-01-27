package main

import (
	"flag"
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v2"
)

// Constants for TravelType
const (
	Metros uint = iota // Metros = 0
	Buses              // Buses = 1
	Trains             // Trains = 2
	Trams              // Trams = 3
	Ships              // Ships = 4
)

// Configuration holds all config data
type Configuration struct {
	Dryrun  bool
	Verbose bool

	ServerRoot    string
	APIKey        string
	SiteID        string
	SiteDirection int
	TravelType    uint

	InfluxServer string
	InfluxDB     string
	InfluxUser   string
	InfluxPasswd string
}

func getConfig() *Configuration {
	var conffilename string
	var dryrun bool
	var verbose bool

	// define and parse command line arguments
	flag.StringVar(&conffilename, "conf", "", "Filename to read configuration from")
	flag.BoolVar(&dryrun, "dryrun", false, "Print results instead of writing to InfluxDB")
	flag.BoolVar(&verbose, "verbose", false, "print more information while running")
	flag.Parse()

	if conffilename == "" {
		log.Fatal("Config file must be given.")
	}

	// read configuration
	config, err := readConfigFile(conffilename)
	if err != nil {
		log.Fatal(err)
	}

	// Dryrun and verbose are only accepted from command line
	config.Dryrun = dryrun
	config.Verbose = verbose

	// done
	return checkConfiguration(config)
}

func readConfigFile(filename string) (*Configuration, error) {
	source, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	config := &Configuration{}
	err = yaml.Unmarshal(source, config)
	if err != nil {
		return nil, err
	}
	return config, nil
}

func checkConfiguration(config *Configuration) *Configuration {

	// ServerRoot Config
	if len(config.ServerRoot) == 0 {
		log.Fatal("ServerRoot must be given.")
	}

	// APIKey Config
	if len(config.APIKey) == 0 {
		log.Fatal("APIKey must be given.")
	}

	// SiteID Config
	if len(config.SiteID) == 0 {
		log.Fatal("SiteID must be given.")
	}

	// SiteDirection Config
	if config.SiteDirection != 1 && config.SiteDirection != 2 {
		log.Fatal("SiteDirection must be 1 or 2")
	}

	// TravelType Config
	if config.TravelType < Metros || Ships < config.TravelType {
		log.Fatal("TravelType must be in range 0-4 (0=Metros,1=Buses,2=Trains,3=Trams,4=Ships)")
	}

	// Influx config
	if !config.Dryrun {
		if len(config.InfluxServer) == 0 {
			log.Fatal("Influx server address must be given.")
		}
		if len(config.InfluxDB) == 0 {
			log.Fatal("Influx database must be given.")
		}
		if (len(config.InfluxUser) == 0) && (len(config.InfluxPasswd) > 0) {
			log.Fatal("Influx user and password must be given (not only one).")
		}
		if (len(config.InfluxUser) > 0) && (len(config.InfluxPasswd) == 0) {
			log.Fatal("Influx user and password must be given (not only one).")
		}
	}

	// Done
	return config
}
