package internal

type ForwardDto struct {
	KubeconfigPath string
	ContextName    string
	Namespace      string
	ForwardType    string
	TargetName     string
}

const (
	ForwardPod     = "📦 Pod"
	ForwardService = "🌐 Service"
)
