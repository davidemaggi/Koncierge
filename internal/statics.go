package internal

const (
	ForwardPod     = "📦 Pod"
	ForwardService = "🌐 Service"

	ConfigTypeSecret = "🔑 Secret"
	ConfigTypeMap    = "🔧 ConfigMap"

	BoolYes = "✅ Yes"
	BoolNo  = "🛑 No"

	FORWARD_ADD_SHORT = "Add a new port-forward"

	FORWARD_ADD_DESCRIPTION = `Add a new port-forward to your list, you can even store related secrets and config map...
As a dev one of the most tedious activities is to find out that a password changed and do the walk of shame on the cluster...`

	FORWARD_SHORT       = "Just start a new forward"
	FORWARD_DESCRIPTION = `A wizard driven port-forward nobody wants to remember commands.`

	FORWARD_START_SHORT       = "Start one or more of your saved forwards"
	FORWARD_START_DESCRIPTION = `Well, you saved your forwards, here you can start them easily`

	FORWARD_DELETE_SHORT       = "Delete one or more of your saved forwards"
	FORWARD_DELETE_DESCRIPTION = `There is no better sensation than closing a project, let's remove also some forwards'`

	CONTEXT_SHORT       = "Change the Current Context"
	CONTEXT_DESCRIPTION = `Here you can change the current ctx for the desired KubeConfig`

	NAMESPACE_SHORT       = "Change the current namespace"
	NAMESPACE_DESCRIPTION = `Change the current namespace for the desired KubeConfig`

	CONTEXT_MERGE_SHORT       = "Copy Context to one config to another"
	CONTEXT_MERGE_DESCRIPTION = `Too many projects, too many contexts, put them all together`

	CONTEXT_DELETE_SHORT       = "Remove one or more contexts from you config file"
	CONTEXT_DELETE_DESCRIPTION = `Too many projects, too many contexts, remove some`
)
