package k8s

import (
	"fmt"
	"github.com/davidemaggi/koncierge/internal"
	"github.com/davidemaggi/koncierge/internal/container"
	"k8s.io/client-go/tools/portforward"
	"k8s.io/client-go/transport/spdy"
	"net/http"
	"os"
)

func StartPortForward(fwd internal.ForwardDto) (stopChan chan struct{}, readyChan chan struct{}, err error) {

	logger := container.App.Logger
	/*
		if params.Out == nil {
			params.Out = io.Discard
		}
		if params.ErrOut == nil {
			params.ErrOut = io.Discard
		}
	*/
	// Build the port forward URL
	req := k8sClient.CoreV1().RESTClient().Post().
		Resource("pods").
		Namespace(fwd.Namespace).
		Name(fwd.PodName).
		SubResource("portforward")

	transport, upgrader, err := spdy.RoundTripperFor(k8sConfig)
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
	forwarder, err := portforward.New(dialer, []string{portMapping}, stopChan, readyChan, os.Stdout, os.Stderr)
	if err != nil {
		return nil, nil, fmt.Errorf("Failed to create port forwarder: %w", err)
	}

	// Run it in background
	go func() {
		if err := forwarder.ForwardPorts(); err != nil {
			logger.Error("Error Port Forward")
		}
	}()

	return stopChan, readyChan, nil
}
