package main

import (
	"github.com/learning/api/server"
	"flag"
	log "github.com/Sirupsen/logrus"
	"os"
	"time"
	"github.com/asaskevich/govalidator"
	"fmt"
	"errors"
)

var (
	bindAddress = flag.String(
		"bind-address",
		"127.0.0.1",
		"The address to which to bind.")
	bindPort = flag.String(
		"bind-port",
		"8080",
		"The port to which to bind.")
)

func main() {
	flag.Parse()

	logConfig()

	errs := []error{}
	errs = ValidateAsHostname("bind-adress", *bindAddress, errs)
	errs = ValidateAsPort("bind-port", *bindPort, errs)

	ExitIfNotEmpty(errs, "API")

	apiURL := fmt.Sprintf("%s:%s", *bindAddress, *bindPort)
	log.Infof("The API is binding to : [address: http://%v:%v]", *bindAddress, *bindPort)
	server.StartAPI(apiURL)
}

func logConfig(){
	log.SetOutput(os.Stdout)
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
		TimestampFormat: time.UnixDate,
		ForceColors: true,
	})
}

func ValidateAsHostname(flagName string, flagValue string, errs []error) []error {
	if !govalidator.IsHost(flagValue) {
		errs = append(errs, errors.New(fmt.Sprintf(
			"Expected a hostname for -%s but got %#v", flagName, flagValue)))
	}
	return errs
}

func ValidateAsPort(flagName string, flagValue string, errs []error) []error {
	if !govalidator.IsPort(flagValue) {
		errs = append(errs, errors.New(fmt.Sprintf(
			"Expected a port for -%s but got %#v", flagName, flagValue)))
	}
	return errs
}

func ExitIfNotEmpty(errs []error, cmdName string) {
	if len(errs) > 0 {
		for _, err := range errs {
			fmt.Fprintf(os.Stdout, "ERROR: %v\n", err)
		}
		fmt.Fprintln(os.Stdout)
		ExitWithUsage(cmdName)
	}
}

func ExitWithUsage(cmdName string) {
	fmt.Fprint(os.Stdout, "Usage:\n")
	fmt.Fprintf(os.Stdout, "  %s [options]\n\n", cmdName)
	fmt.Fprint(os.Stdout, "Options:\n")
	flag.PrintDefaults()
	os.Exit(1)
}
