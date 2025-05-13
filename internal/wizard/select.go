package wizard

import (
	"github.com/davidemaggi/koncierge/internal/config"
	"github.com/davidemaggi/koncierge/internal/container"
	"github.com/davidemaggi/koncierge/internal/k8s"
	"github.com/pterm/pterm"
)

func SelectNamespace() string {

	logger := container.App.Logger

	spaces, err := k8s.GetAllNameSpaces(config.KubeConfigFile)

	if err != nil {
		return ""
	}

	selectedOption := ""
	current := k8s.GetCurrentNamespaceForContext(config.KubeConfigFile, k8s.GetCurrentContextAsString(config.KubeConfigFile))

	if len(spaces) == 0 {
		logger.Info("No namespace available in " + pterm.Green(config.KubeConfigFile))

		// Display the selected option to the user with a green color for emphasis
	}

	if len(spaces) == 1 {
		selectedOption = spaces[0]
		logger.Info("Only " + pterm.Green("one") + " namespace is available")

		// Display the selected option to the user with a green color for emphasis
	}

	if len(spaces) >= 2 {
		if current == "" {
			current = spaces[0]
		}
		selectedOption, _ = pterm.DefaultInteractiveSelect.WithOptions(spaces).WithDefaultOption(current).Show()

	}

	if selectedOption == current {
		logger.Warn("Selected and Current namespace are the same " + pterm.Yellow("Skipping"))

	}
	return selectedOption
}

func SelectContext() string {

	logger := container.App.Logger

	contexts := k8s.GetAllContexts(config.KubeConfigFile)
	selectedOption := ""
	current := k8s.GetCurrentContextAsString(config.KubeConfigFile)

	if len(contexts) == 0 {
		logger.Info("No context available in " + pterm.Green(config.KubeConfigFile))

		// Display the selected option to the user with a green color for emphasis
	}

	if len(contexts) == 1 {
		selectedOption = contexts[0]
		logger.Info("Only " + pterm.Green("one") + " context is available")

		// Display the selected option to the user with a green color for emphasis
	}

	if len(contexts) >= 2 {
		if current == "" {
			current = contexts[0]
		}

		selectedOption, _ = pterm.DefaultInteractiveSelect.WithOptions(contexts).WithDefaultOption(current).Show()

	}

	if selectedOption == current {
		logger.Warn("Selected and Current context are the same " + pterm.Yellow("Skipping"))

	}

	return selectedOption

}

func SelectService(namespace string) string {

	logger := container.App.Logger

	services := k8s.GetServicesInNamespace(namespace)

	selectedOption := ""
	current := ""

	if len(services) == 0 {
		logger.Info("No namespace available in " + pterm.Green(config.KubeConfigFile))

		// Display the selected option to the user with a green color for emphasis
	}

	if len(services) == 1 {
		selectedOption = services[0]
		logger.Info("Only " + pterm.Green("one") + " namespace is available")

		// Display the selected option to the user with a green color for emphasis
	}

	if len(services) >= 2 {
		if current == "" {
			current = services[0]
		}
		selectedOption, _ = pterm.DefaultInteractiveSelect.WithOptions(services).WithDefaultOption(current).Show()

	}

	return selectedOption
}
