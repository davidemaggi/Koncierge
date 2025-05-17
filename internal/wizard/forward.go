package wizard

import (
	"fmt"
	"github.com/davidemaggi/koncierge/internal"
	"github.com/davidemaggi/koncierge/internal/config"
	"github.com/davidemaggi/koncierge/internal/container"
	"github.com/davidemaggi/koncierge/internal/k8s"
	"github.com/pterm/pterm"
	"os"
	"strconv"
)

func BuildForward() internal.ForwardDto {

	var ret internal.ForwardDto
	logger := container.App.Logger

	ret.KubeconfigPath = config.KubeConfigFile
	ret.ContextName = config.KubeContext
	kubeService, _ := k8s.ConnectToClusterAndContext(config.KubeConfigFile, config.KubeContext)
	spaces, err := kubeService.GetAllNameSpaces()

	if err != nil {
		logger.Error("Error retrieving namespaces")
	}

	current := k8s.GetCurrentNamespaceForContext(config.KubeConfigFile, config.KubeContext)

	selNamespace, ok := SelectOne(spaces, "Select a namespace", func(f string) string {
		return f
	}, current)
	if !ok {
		os.Exit(1)
	}
	//selectedOption, _ := pterm.DefaultInteractiveSelect.WithOptions(options).Show()
	ret.Namespace = selNamespace
	var fwdtypes []string

	fwdtypes = append(fwdtypes, internal.ForwardService)
	fwdtypes = append(fwdtypes, internal.ForwardPod)

	ret.ForwardType, _ = pterm.DefaultInteractiveSelect.WithOptions(fwdtypes).Show()
	var ports []internal.ServicePortDto
	if ret.ForwardType == internal.ForwardService {

		services := kubeService.GetServicesInNamespace(ret.Namespace)

		selTarget, ok := SelectOne(services, "Select a service", func(s string) string { return s }, "")
		if !ok {
			os.Exit(1)
		}

		ret.TargetName = selTarget
		ret.PodName, err = kubeService.GetFirstPodForService(ret.Namespace, ret.TargetName)

		if err != nil {
			logger.Error("Error retrieving pod")
			os.Exit(1)
		}

		ports = kubeService.GetServicePorts(ret.Namespace, ret.TargetName)

	}

	if ret.ForwardType == internal.ForwardPod {
		// TODO: get Pod ports
		//ret.TargetName, _ = pterm.DefaultInteractiveSelect.WithOptions(k8s.GetPodsInNamespace(ret.Namespace)).Show()
		//ports = k8s.GetServicePorts(ret.Namespace, ret.TargetName)
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

	ret.TargetPort = selectedPort.ServicePort

	localPortTxt, _ := pterm.DefaultInteractiveTextInput.WithDefaultValue(fmt.Sprintf("%d", ret.TargetPort)).Show()

	if val, err := strconv.ParseInt(localPortTxt, 10, 32); err == nil {
		ret.LocalPort = int32(val)

	} else {
		logger.Error("Failed to parse local port number")
	}

	return ret
}
