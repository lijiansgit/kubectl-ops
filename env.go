package main

const (
	// shell: export env=123
	// jenkins shell build parms
	//
	// consul addr, default:127.0.0.1:8500
	// "CONSUL_HTTP_ADDR"
	// DeployName and PodName = jenkins job name, JOB_NAME
	// export K8S_DEPLOY_NAME=$JOB_NAME K8S_POD_NAME=$JOB_NAME
	DeployName = "K8S_DEPLOY_NAME"
	// default namespace: default
	DeployNamespace                  = "K8S_DEPLOY_NAMESPACE"
	DeployReplicas                   = "K8S_DEPLOY_REPLICAS"
	DeployMinReadySeconds            = "K8S_DEPLOY_MRS"
	DeployRevisionHistoryLimit       = "K8S_DEPLOY_RHL"
	PodNamespace                     = "K8S_POD_NAMESPACE"
	PodName                          = "K8S_POD_NAME"
	PodTerminationGracePeriodSeconds = "K8S_POD_TGPS"
	// hpa
	// K8S_POD_HPA == "0" hpa off
	PodHPA     = "K8S_POD_HPA"
	PodHPAKind = "K8S_POD_HPA_KIND"
	PodHPAMin  = "K8S_POD_HPA_MIN"
	PodHPAMax  = "K8S_POD_HPA_MAX"
	PodHPACPU  = "K8S_POD_HPA_CPU"
	// ContainerPort,first post must 8080, eg: "8080,8081"
	AppPort           = "K8S_APP_PORT"
	AppLimitsCPU      = "K8S_APP_LIMIT_CPU"
	AppLimitsMemory   = "K8S_APP_LIMIT_MEM"
	AppRequestsCPU    = "K8S_APP_REQ_CPU"
	AppRequestsMemory = "K8S_APP_REQ_MEM"
	AppLivenessPath   = "K8S_APP_LIVE_PATH"
	AppLivenessDelay  = "K8S_APP_LIVE_DELAY"
	AppReadinessPath  = "K8S_APP_READ_PATH"
	AppReadinessDelay = "K8S_APP_READ_DELAY"
	// K8S_APP_HEALTH_CHECK == "0" liveness or readiness off
	AppHealthCheck = "K8S_APP_HEALTH_CHECK"
	AppCmd         = "K8S_APP_CMD"
	AppCmdPath     = "K8S_APP_CMD_PATH"
	AppBuildCmd    = "K8S_APP_BUILD_CMD"
	AppBuildPath   = "K8S_APP_BUILD_PATH"
	DockerHub      = "K8S_DOCKER_HUB"
	//DeploySelectorLabelApp = "K8S_DEPLOY_SLA"
	//DeployRollingUpdateMaxSurge	=	"K8S_DEPLOY_RUMS"
	//DeployRollingUpdateMaxUnavailable	=	"K8S_DEPLOY_RUMU"
	//DeployStrategyType	=	"K8S_DEPLOY_ST"
	//PodLabelApp = "K8S_POD_LA"
	//PodRestartPolicy = "K8S_POD_RP"
	//AppName = "K8S_APP_NAME"
	//AppImage = "K8S_APP_IMAGE"
	//AppImagePullPolicy = "K8S_APP_IPP"
	//AppPortName = "K8S_APP_PORT"
	//AppENV = "K8S_APP_ENV"
	//AppLivenessPort   = "K8S_APP_LIVE_PORT"
	//AppReadinessPort  = "K8S_APP_READ_PORT"
	//AppCopyDir					= "K8S_APP_COPY_DIR"

	// Jenkins variabes
	AppGitBranch = "GIT_BRANCH"
	AppEnv       = "K8S_APP_ENV"
	ReleaseCheck = "K8S_RELEASE_CHECK"
	// release action: check,deploy,gray,rollback
	ReleaseAction = "K8S_RELEASE_ACTION"
)
