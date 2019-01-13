package main

import (
	"fmt"
	"os"

	"github.com/lijiansgit/go/libs"
	log "github.com/lijiansgit/go/libs/log4go"
)

func build() (err error) {
	if err = writeDockerFile(); err != nil {
		return err
	}

	log.Info("docker build %s", config.image)
	buildCmd := fmt.Sprintf("docker build -t %s .", config.image)
	_, err = libs.Cmd(buildCmd, config.appBuildPath)
	if err != nil {
		return err
	}

	log.Info("docker push %s", config.image)
	pushCmd := fmt.Sprintf("docker push %s", config.image)
	_, err = libs.Cmd(pushCmd, config.appBuildPath)
	if err != nil {
		return err
	}

	return nil
}

func writeDockerFile() (err error) {
	_, err = os.Stat(dockerFileName)
	if err == nil {
		log.Info("local %s file found", dockerFileName)
		return nil
	}

	log.Warn("local %s no found! read from consul", dockerFileName)
	log.Debug("consul %s content:\n%s", dockerFileName, config.dockerFile)
	f, err := os.Create(dockerFileName)
	if err != nil {
		return err
	}

	defer f.Close()

	_, err = f.WriteString(config.dockerFile)
	if err != nil {
		return err
	}

	return err
}
