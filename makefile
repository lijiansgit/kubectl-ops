CURRENT_DIR=$(shell pwd)
GOPATH_DIR=$(shell echo ${CURRENT_DIR}/../../)
DATE_TIME=`date '+%Y-%m-%d %H:%M:%S'`

common:
	go build -o dev

dev: common
	echo "dev"
	./dev

com: common
	echo "company test"
	scp dev root@192.168.56.102:/qtt/bin/kubectl-ops
