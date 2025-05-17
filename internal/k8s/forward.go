package k8s

import (
	"fmt"
	"github.com/davidemaggi/koncierge/internal"
	"github.com/davidemaggi/koncierge/internal/container"
	"github.com/pterm/pterm"
	"io"
	"k8s.io/client-go/tools/portforward"
	"k8s.io/client-go/transport/spdy"
	"net/http"
)

func (k *KubeService) StartPortForward(fwd internal.ForwardDto) (stopChan chan struct{}, readyChan chan struct{}, err error) {

	logger := container.App.Logger

	// Build the port forward URL
	req := k.client.CoreV1().RESTClient().Post().
		Resource("pods").
		Namespace(fwd.Namespace).
		Name(fwd.PodName).
		SubResource("portforward")

	transport, upgrader, err := spdy.RoundTripperFor(k.config)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create round tripper: %w", err)
	}

	// Set up the dialer
	dialer := spdy.NewDialer(upgrader, &http.Client{Transport: transport}, "POST", req.URL())

	// Format port mapping
	portMapping := fmt.Sprintf("%d:%d", fwd.LocalPort, fwd.TargetPort)

	stopChan = make(chan struct{}, 1)
	readyChan = make(chan struct{})

	// Create the port forwarder
	forwarder, err := portforward.New(dialer, []string{portMapping}, stopChan, readyChan, io.Discard, pterm.Error.Writer)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create port forwarder: %w", err)
	}

	// Run it in background
	go func() {
		if err := forwarder.ForwardPorts(); err != nil {
			logger.Error("Error Port Forward", err)
		}
	}()

	return stopChan, readyChan, nil
}
