#!/bin/bash
CPWD=`pwd`
yum install epel-release -y
yum install golang -y
echo "export GOPATH=/usr/share/gocode" >> /etc/profile.d/golang.sh
export GOPATH=/usr/share/gocode
go get -v github.com/lijiansgit/go/libs
go get -v github.com/google/gofuzz
go get -v github.com/hashicorp/consul/api
mkdir -pv $GOPATH/src/k8s.io
cd $GOPATH/src/k8s.io
git clone -v https://github.com/kubernetes/apimachinery
git clone -v https://github.com/kubernetes/klog.git
git clone -v https://github.com/kubernetes/client-go.git
git clone -v https://github.com/kubernetes/api.git
git clone -v https://github.com/kubernetes/utils.git
cd $CPWD
go build -o /usr/local/bin/kubectl-ops