package main

import "errors"

func commandExit(_ *Config, _ []string) error {
	return errors.New("Exiting")
}
