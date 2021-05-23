package main

import "os"

func GetEnvironmentVariable(name string, defaultValue string) string {
	environmentVariable, isPresent := os.LookupEnv(name)

	if !isPresent {
		return defaultValue
	}

	return environmentVariable
}
