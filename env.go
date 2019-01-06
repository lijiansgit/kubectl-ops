package main

const (
	// shell: export env=123
	// jenkins shell build parms
	//
	// default jenkins project name
	// DeployName = "K8S_DEPLOY_NAME"
	// default namespace: default
	//DeployNamespace= "K8S_DEPLOY_NAMESPACE"
	DeployReplicas = "K8S_DEPLOY_REPLICAS"
	DeployMinReadySeconds = "K8S_DEPLOY_MRS"
	DeployRevisionHistoryLimit = "K8S_DEPLOY_RHL"
	//DeploySelectorLabelApp = "K8S_DEPLOY_SLA"
	DeployRollingUpdateMaxSurge	=	"K8S_DEPLOY_RUMS"
	DeployRollingUpdateMaxUnavailable	=	"K8S_DEPLOY_RUMU"
	DeployStrategyType	=	"K8S_DEPLOY_ST"
	PodName = "K8S_POD_NAME"
	PodNamespace = "K8S_POD_NAMESPACE"
	//PodLabelApp = "K8S_POD_LA"
	PodTerminationGracePeriodSeconds = "K8S_POD_TGPS"
	PodRestartPolicy = "K8S_POD_RP"
	//AppName = "K8S_APP_NAME"
	AppImage = "K8S_APP_IMAGE"
	AppImagePullPolicy = "K8S_APP_IPP"
	AppPort = "K8S_APP_PORT"
	//AppENV = "K8S_APP_ENV"
	AppLimitsCPU = "K8S_APP_LIMIT_CPU"
	AppLimitsMemory = "K8S_APP_LIMIT_MEM"
	AppRequestsCPU = "K8S_APP_REQ_CPU"
	AppRequestsMemory = "K8S_APP_REQ_MEM"
	AppLivenessPath = "K8S_APP_LIVE_PATH"
	AppLivenessPort = "K8S_APP_LIVE_PORT"
	AppLivenessDelay = "K8S_APP_LIVE_DELAY"
	AppReadinessPath = "K8S_APP_READ_PATH"
	AppReadinessPort = "K8S_APP_READ_PORT"
	AppReadinessDelay = "K8S_APP_READ_DELAY"
)
