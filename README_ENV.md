# kubectl-ops 资源参数

----
由于资源对象较多，且更改较为频繁，所以设置为取环境变量的方式读取，设置方式：
```bash
export K8S_DEPLOY_NAME=“$JOB_NAME” K8S_DEPLOY_REPLICAS=“2”
```
----

## 参数列表

列出的参数是可更改的，所有值为字符串

Consul配置里读取，也可单独设置：

|Name|default|explain|
|:-----:|:-------:|:-------:|
|K8S_DEPLOY_NAME|go|Deployment 名称|
|K8S_DEPLOY_NAMESPACE|default|Deployment 命名空间|
|K8S_DEPLOY_REPLICAS|1|Deployment 副本数量|
|K8S_DEPLOY_MRS|0|Pod 在可用之前的最小准备时间|
|K8S_DEPLOY_RHL|3|Deployment 历史版本保留数量|
|K8S_POD_NAMESPACE|default|Pod 命名空间|
|K8S_POD_NAME|go|Pod 名称|
|K8S_POD_TGPS|5|Pod 在删除之前的最小等待时间|
|K8S_POD_HPA_KIND|Deployment|HPA 的资源类型|
|K8S_POD_HPA_MIN|2|Deployment 的最小副本数量|
|K8S_POD_HPA_MAX|10|Deployment 的最大副本数量|
|K8S_POD_HPA_CPU|35|Pod CPU使用率在达到多少百分比时扩容|
|K8S_APP_PORT|8080|容器端口，可指定多个，用逗号分隔|
|K8S_APP_LIMIT_CPU|100m|容器可使用CPU的最大值|
|K8S_APP_LIMIT_MEM|100Mi|容器可使用内存的最大值|
|K8S_APP_REQ_CPU|30m|容器CPU初始分配值|
|K8S_APP_REQ_MEM|30Mi|容器内存初始分配值|
|K8S_APP_LIVE_PATH|/ping|容器健康监测HTTP路径|
|K8S_APP_LIVE_DELAY|2|容器健康监测延迟多少秒|
|K8S_APP_READ_PATH|/ping|容器启动健康监测HTTP路径|
|K8S_APP_READ_DELAY|2|容器启动健康监测延迟多少秒|
|K8S_APP_CMD|/test|容器启动命令|
|K8S_APP_CMD_PATH|/|容器运行默认路径|

命令行设置：

|Name|default|explain|
|:--------:|:--------:|:------------:|
|K8S_POD_HPA|1|自动扩容是否打开，0为关闭，1为打开|
|K8S_APP_HEALTH_CHECK|1|容器健康监测是否打开，0为关闭，1为打开|
|K8S_DOCKER_HUB|test.hub.com/test|镜像推送HUB地址+项目名称|
|K8S_APP_ENV|dev|容器运行环境：dev/qa/pre/prd|
|K8S_RELEASE_CHECK|1|容器发布之后运行是否成功检查，0为关闭，1为打开|
|K8S_RELEASE_ACTION|check|容器发布动作，check:检查运行状态/deploy:发布/gray:灰度/rollback:回滚|
|GIT_BRANCH|origin/master|GIT代码分支/TAG|
