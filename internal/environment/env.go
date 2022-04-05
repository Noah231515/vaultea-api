package environment

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

var env map[string]interface{}

func SetEnv() {
	env = make(map[string]interface{})
	envData, err := ioutil.ReadFile("env.json") // TODO: make absolute
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(envData, &env)
	fmt.Println(env["SECRET_KEY"])
}
