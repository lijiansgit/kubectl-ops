# kubectl-ops

[![Go Report](https://goreportcard.com/badge/github.com/lijiansgit/kubectl-ops)](https://goreportcard.com/report/github.com/lijiansgit/kubectl-ops)
[![Build Status](https://travis-ci.org/lijiansgit/kubectl-ops.svg?branch=master)](https://travis-ci.org/lijiansgit/kubectl-ops)

----
kubernetes 发布客户端工具，结合Jenkins使用，支持发布Deployment、Service、HPA，发布动作支持上线、灰度、回滚等功能
----


* [Pre Requisites](#pre-requisites)
* [Quick Start](#quick-start)
* [Installation](#installation)
* [Configuration](#configuration)
* [Usage](#usage)
* [Contribute](#contribute)
* [TODO](#todo)


## Pre Requisites
Kubernetes: 
* Server 1.10+

System: 
* CentOS 7.0+

Software: 
* Golang
* Consul
* Jenkins
* Harbor

## Quick Start

```bash
yum install git -y
git clone -v https://github.com/lijiansgit/kubectl-ops
bash kubectl-ops/install.sh
kubectl-ops -h
```

## Installation

### Consul

安装程序

```bash
yum install wget unzip -y
wget https://releases.hashicorp.com/consul/1.4.2/consul_1.4.2_linux_amd64.zip
unzip consul_1.4.2_linux_amd64.zip
mv consul /usr/local/bin/
cp kubectl-ops/service/consul.service /usr/lib/systemd/system/
mkdir -pv /data/consul
systemctl enable consul && systemctl start consul
```

导入kubectl-ops初始化配置

```bash
echo "export CONSUL_HTTP_ADDR=127.0.0.1:8500" >> /etc/profile.d/consul.sh
source /etc/profile.d/consul.sh
consul kv import @kubectl-ops/conf/kubernetes.json
```

### Harbor

安装程序

此处略，官方文档：https://github.com/goharbor/harbor/blob/master/docs/installation_guide.md

默认Harbor地址为test.hub.com，参数设置见下文

### Jenkins

安装程序

```bash
yum install java-1.8.0-openjdk -y
mkdir -pv /data/jenkins/data
wget http://mirrors.jenkins.io/war-stable/latest/jenkins.war -O /data/jenkins/jenkins.war
cp kubectl-ops/service/jenkins.service /usr/lib/systemd/system/
echo "JENKINS_HOME=/data/jenkins/data" >> /etc/sysconfig/jenkins
echo "GOPATH=/usr/share/gocode" >> /etc/sysconfig/jenkins
systemctl enable jenkins && systemctl start jenkins
# jenkins启动初始化较慢，需耐心等待...
```

## Configuration

Install Plugin

* Git Parameter

Config jenkins

创建任务

![1](https://github.com/lijiansgit/kubectl-ops/raw/master/png/1.png)

参数化构建过程参数设置，GIT源：https://github.com/lijiansgit/test.git

![2](https://github.com/lijiansgit/kubectl-ops/raw/master/png/2.png)
![3](https://github.com/lijiansgit/kubectl-ops/raw/master/png/3.png)
![4](https://github.com/lijiansgit/kubectl-ops/raw/master/png/4.png)

## Usage

查看命令帮助
```bash
kubectl-ops -h
```

-c: 连接kubernetes apiserver的配置文件，默认为$HOME/.kube/config

-cp: consul 配置路径，默认为kubernetes/v1

Dockerfile从consul读取默认值，如果发现代码根目录存在此文件，则使用代码自定义Dockerfile

发布对象资源参数：[详解](./README_ENV.md)

## Contribute

欢迎提交问题及反馈

## TODO

* 比例和动态灰度
* Ingress


------------------------
