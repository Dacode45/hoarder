//
package config

import (
	"io/ioutil"
	"path/filepath"
	"strconv"

	"github.com/ghodss/yaml"
	"github.com/jcelliott/lumber"

	"github.com/nanopack/hoarder/backends"
)

//
const (
	DEFAULT_BACKEND  = "filesystem"
	DEFAULT_HOST     = "0.0.0.0"
	DEFAULT_LOGLEVEL = "info"
	DEFAULT_PORT     = ":7410"
	DEFAULT_TOKEN    = "TOKEN"
	VERSION          = "0.0.1"
)

//
var (

	//  configurable options
	Backend    = DEFAULT_BACKEND  // the pluggable backend the api will use for storage
	Driver     backends.Driver    // the actual backend driver
	GCInterval = 0                // the interval between clearning out old storage
	GCAmount   = 0                // the amount of storage to clear at interval
	Host       = DEFAULT_HOST     // the connection host
	Insecure   = false            // connect insecurly
	LogLevel   = DEFAULT_LOGLEVEL // the output log level
	Port       = DEFAULT_PORT     // the connection port
	Token      = DEFAULT_TOKEN    // the secury token used to connect with

	// internal options
	Addr = Host + Port // the host:port combo for connecting to the server
	Log  lumber.Logger // the logger to use
)

//
func Parse(path string) error {

	// if a config is provided (and found), parse the config file overwriting any
	// defaults
	if fp, err := filepath.Abs(path); err == nil {

		//
		f, err := ioutil.ReadFile(fp)
		if err != nil {
			return err
		}

		// parse config file
		options := map[string]string{}
		if err := yaml.Unmarshal(f, options); err != nil {
			return err
		}

		// override defaults
		for k, v := range options {
			switch k {
			case "backend":
				Backend = v
			case "gc_interval":
				i, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				GCInterval = i
			case "gc_amount":
				i, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				GCAmount = i
			case "host":
				Host = v
			case "insecure":
				b, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				Insecure = b
			case "log_level":
				LogLevel = v
			case "port":
				Port = v
			case "token":
				Token = v
			}
		}
	}

	// create a new logger
	Log = lumber.NewConsoleLogger(lumber.LvlInt(LogLevel))
	Log.Prefix("[hoarder]")

	return nil
}