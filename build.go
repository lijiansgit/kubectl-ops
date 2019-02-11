package main

import (
	"fmt"
	"os"

	"github.com/lijiansgit/go/libs"
	log "github.com/lijiansgit/go/libs/log4go"
)

const (
	dockerFileName   = "Dockerfile"
	dockerIgnoreName = ".dockerignore"
)

func build() (err error) {
	if err = writeDockerFile(); err != nil {
		return err
	}

	buildCmd := fmt.Sprintf("docker build -t %s .", config.image)
	log.Info("build cmd: %s", buildCmd)
	res, err := libs.Cmd(buildCmd, config.appBuildPath)
	if err != nil {
		return err
	}

	log.Debug("build cmd res: %s", res)
	pushCmd := fmt.Sprintf("docker push %s", config.image)
	log.Info("push cmd: %s", pushCmd)
	res, err = libs.Cmd(pushCmd, config.appBuildPath)
	if err != nil {
		return err
	}

	log.Debug("push cmd res: %s", res)
	return nil
}

func writeDockerFile() (err error) {
	fi, err := os.Create(dockerIgnoreName)
	if err != nil {
		return err
	}

	defer fi.Close()

	_, err = fi.WriteString(config.dockerIgnore)
	if err != nil {
		return err
	}

	log.Debug("%s write ok", dockerIgnoreName)

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

	return nil
}
