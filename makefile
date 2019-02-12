CURRENT_DIR=$(shell pwd)
GOPATH_DIR=$(shell echo ${CURRENT_DIR}/../../)
DATE_TIME=`date '+%Y-%m-%d %H:%M:%S'`

common:
	go build -o dev
	dos2unix install.sh

dev: common
	echo "dev"
	./dev

com: common
	echo "company test"
	scp dev root@10.0.2.36:/bin/kubectl-ops

home: common
	echo "home test"
	scp dev root@192.168.1.103:/bin/kubectl-ops
