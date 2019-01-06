CURRENT_DIR=$(shell pwd)
GOPATH_DIR=$(shell echo ${CURRENT_DIR}/../../)
DATE_TIME=`date '+%Y-%m-%d %H:%M:%S'`

common:
	go build -o dev

dev: common
	echo "dev"
	./dev
