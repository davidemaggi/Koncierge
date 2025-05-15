package wizard

import (
	"fmt"
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
	var ports []internal.ServicePortDto
	if ret.ForwardType == internal.ForwardService {
		ret.TargetName, _ = pterm.DefaultInteractiveSelect.WithOptions(k8s.GetServicesInNamespace(ret.Namespace)).Show()
		ports = k8s.GetServicePorts(ret.Namespace, ret.TargetName) // Arrivato qui

	}

	if ret.ForwardType == internal.ForwardPod {
		// TODO: get Pod ports

	}

	var portOptions []string

	portMap := make(map[string]internal.ServicePortDto)

	for _, port := range ports {
		tmpKey := fmt.Sprintf("%d (%s)", port.PodPort, port.Protocol)
		portOptions = append(portOptions, tmpKey)
		portMap[tmpKey] = port
	}

	selectedName, _ := pterm.DefaultInteractiveSelect.
		WithOptions(portOptions).
		WithDefaultText("Select a port").
		Show()

	// Retrieve full object based on name
	selectedPort := portMap[selectedName]

	_ = selectedPort

	return ret
}
