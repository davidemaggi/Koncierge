package internal

type ForwardDto struct {
	KubeconfigPath string
	ContextName    string
	Namespace      string
	ForwardType    string
	TargetName     string
	PodName        string
	TargetPort     int32
	LocalPort      int32
}

type ServicePortDto struct {
	Protocol    string
	ServicePort int32
	PodPort     int32
	Podname     string
}

const (
	ForwardPod     = "📦 Pod"
	ForwardService = "🌐 Service"
)
