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
	FORWARD_DELETE_DESCRIPTION = `There is no better sensation than closing a project, let's remove also some forwards`

	FORWARD_LIST_SHORT       = "List all the known forwards"
	FORWARD_LIST_DESCRIPTION = `Just in case you forgot about them... leets have a look`

	FORWARD_MOVE_SHORT       = "Move some forwards to a different context"
	FORWARD_MOVE_DESCRIPTION = `Context changed but the services are the same? copy Forwards instead of recreating them`

	FORWARD_EDIT_SHORT       = "Update an existing forward"
	FORWARD_EDIT_DESCRIPTION = `You regret your life choices? I cannot do much about it but for sure i can help you editing an existing Port Forward`

	FORWARD_COPY_SHORT       = "Copy one or more forwards to a different context"
	FORWARD_COPY_DESCRIPTION = `Multiple enviroments with the same configuration? Copy the forwards on other contexts`

	CONTEXT_SHORT       = "Change the Current Context"
	CONTEXT_DESCRIPTION = `Here you can change the current ctx for the desired KubeConfig`

	NAMESPACE_SHORT       = "Change the current namespace"
	NAMESPACE_DESCRIPTION = `Change the current namespace for the desired KubeConfig`

	CONTEXT_MERGE_SHORT       = "Copy Context to one config to another"
	CONTEXT_MERGE_DESCRIPTION = `Too many projects, too many contexts, put them all together`

	CONTEXT_DELETE_SHORT       = "Remove one or more contexts from you config file"
	CONTEXT_DELETE_DESCRIPTION = `Too many projects, too many contexts, remove some`
)
