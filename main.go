package main

import (
	"os"

	"flag"

	"github.com/murlokswarm/log"
)

func main() {
	var force bool
	var verb bool
	var sass bool

	flag.BoolVar(&force, "f", false, "Forces the copy of the whole resources directory rather than sync it.")
	flag.BoolVar(&verb, "v", false, "Verbose mode.")
	flag.BoolVar(&sass, "sass", false, "exec sass --watch resources/scss:resources/css")
	flag.Parse()

	if sass {
		launchSass()
		return
	}

	if err := build(false); err != nil {
		log.Error(err)
		return
	}

	conf, err := readConfig(confName)
	if err != nil {
		if !os.IsNotExist(err) {
			log.Error(err)
			return
		}

		conf = defaultConfig()
		err = saveConfig(conf, confName)
	}

	if err = createPackage(conf); err != nil {
		log.Error(err)
		return
	}

	if err = createExec(conf); err != nil {
		log.Error(err)
		return
	}

	if err = createPlist(conf); err != nil {
		log.Error(err)
		return
	}

	if err = syncResources(conf, force, verb); err != nil {
		log.Error(err)
		return
	}

	if len(conf.Icon) != 0 {
		if err = generateIcon(conf); err != nil {
			log.Error(err)
		}
	}
}
