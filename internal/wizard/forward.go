package wizard

import (
	"github.com/davidemaggi/koncierge/internal"
	"github.com/davidemaggi/koncierge/internal/config"
	"github.com/davidemaggi/koncierge/internal/k8s"
	"github.com/pterm/pterm"
)

func BuildForward() internal.ForwardDto {

	var ret internal.ForwardDto

	ret.KubeconfigPath = config.KubeConfigFile
	ret.ContextName = k8s.GetCurrentContextAsString(config.KubeConfigFile)

	ret.Namespace = SelectNamespace()

	//selectedOption, _ := pterm.DefaultInteractiveSelect.WithOptions(options).Show()

	var fwdtypes []string

	fwdtypes = append(fwdtypes, internal.ForwardService)
	fwdtypes = append(fwdtypes, internal.ForwardPod)

	ret.ForwardType, _ = pterm.DefaultInteractiveSelect.WithOptions(fwdtypes).Show()

	if ret.ForwardType == internal.ForwardService {
		ret.TargetName, _ = pterm.DefaultInteractiveSelect.WithOptions(k8s.GetServicesInNamespace(ret.Namespace)).Show()
		//ports := k8s.GetServicePorts(ret.Namespace, ret.TargetName) // Arrivato qui

	}

	return ret
}
