package main

import (
	"strconv"
	"strings"

	"k8s.io/api/core/v1"
)

func strToInt32(str string) int32 {
	i, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}

	return int32(i)
}

func strToInt32p(str string) *int32 {
	i, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}

	in := int32(i)
	return &in
}

func strToInt64p(str string) *int64 {
	i, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}

	in := int64(i)
	return &in
}

func strToContainerPorts(str string) (ports []v1.ContainerPort) {
	list := strings.Split(str, ",")
	ports = make([]v1.ContainerPort, len(list))
	for k, v := range list {
		ports[k].ContainerPort = strToInt32(v)
	}

	return ports
}
