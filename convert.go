package main

import "strconv"

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
