package main

import (
	"os"
)

func onRegistryWrite(path, tempfile string) error {
	var e error

	// Explicitly Remove .logstash-forwarder if it exists
	if _, err := os.Stat(path); err == nil {
		if e = os.Remove(path); e != nil {
			emit("registry rotate: removal of %s - %s\n", path, e)
			return e
		}
	}

	// Rename .logstash-forwarder.new -> .logstash-forwarder
	if e = os.Rename(tempfile, path); e != nil {
		emit("registry rotate: rename of %s to %s - %s\n", tempfile, path, e)
		return e
	}

	return nil
}
