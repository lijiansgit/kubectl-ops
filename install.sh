#!/bin/bash
yum install epel-release -y
yum install golang -y
echo "export GOPATH=/usr/share/gocode" >> /etc/profile.d/golang.sh
source /etc/profile.d/golang.sh
export GOPATH=/usr/share/gocode
go get -v github.com/lijiansgit/go/libs
go get -v k8s.io/apimachinery/pkg
go get -v github.com/google/gofuzz
cd $GOPATH/src/k8s.io
git clone -v git@github.com:kubernetes/klog.git
git clone -v git@github.com:kubernetes/client-go.git
git clone -v git@github.com:kubernetes/api.git
cd $GOPATH/src/github.com/lijiansgit/kubectl-ops
go install