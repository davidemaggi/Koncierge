package wizard

import (
	"fmt"
	"github.com/davidemaggi/koncierge/internal"
	"github.com/davidemaggi/koncierge/internal/config"
	"github.com/davidemaggi/koncierge/internal/container"
	"github.com/davidemaggi/koncierge/internal/k8s"
	"github.com/davidemaggi/koncierge/internal/ui"
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
		logger.Error("Error retrieving namespaces", err)
	}

	current := k8s.GetCurrentNamespaceForContext(config.KubeConfigFile, config.KubeContext)

	ui.PrintCurrentStatus(ret.KubeconfigPath, ret.ContextName, current)

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
			logger.Error("Error retrieving pod", err)
			os.Exit(1)
		}

		ports = kubeService.GetServicePorts(ret.Namespace, ret.TargetName)

	}

	if ret.ForwardType == internal.ForwardPod {

		// TODO: get Pod ports
		ret.TargetName, _ = pterm.DefaultInteractiveSelect.WithOptions(kubeService.GetPodsInNamespace(ret.Namespace)).Show()
		ret.PodName = ret.TargetName
		ports = kubeService.GetPodPorts(ret.Namespace, ret.TargetName)
	}

	var portOptions []string

	portMap := make(map[string]internal.ServicePortDto)

	for _, port := range ports {
		tmpKey := fmt.Sprintf("%d (%s)", port.ServicePort, port.Protocol)
		portOptions = append(portOptions, tmpKey)
		portMap[tmpKey] = port
	}

	selectedName, _ := pterm.DefaultInteractiveSelect.
		WithOptions(portOptions).
		WithDefaultText("Select a port").
		Show()

	// Retrieve full object based on name
	selectedPort := portMap[selectedName]

	if selectedPort.PodPort == 0 {

		ret.TargetPort = selectedPort.ServicePort

	} else {

		ret.TargetPort = selectedPort.PodPort

	}

	localPortTxt, _ := pterm.DefaultInteractiveTextInput.WithDefaultValue(fmt.Sprintf("%d", ret.TargetPort)).WithDefaultText("Insert the Local Port").Show()

	if val, err := strconv.ParseInt(localPortTxt, 10, 32); err == nil {
		ret.LocalPort = int32(val)

	} else {
		logger.Error("Failed to parse local port number", err)
		os.Exit(1)
	}

	return ret
}
