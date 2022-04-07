package environment

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

var env map[string]string

func SetEnv() {
	env = make(map[string]string)
	envData, err := ioutil.ReadFile("env.json") // TODO: make absolute
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(envData, &env)
}

func GetEnv() map[string]string {
	return env
}
