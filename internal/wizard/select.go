package wizard

import (
	"atomicgo.dev/keyboard/keys"
	"github.com/pterm/pterm"
)

func SelectOne[T any](items []T, label string, getKey func(T) string, current string) (T, bool) {

	var options []string
	itemMap := make(map[string]T)

	for _, item := range items {
		key := getKey(item)
		options = append(options, key)
		itemMap[key] = item
	}

	if current == "" {
		current = getKey(items[0])
	}

	selectedKey, err := pterm.DefaultInteractiveSelect.
		WithOptions(options).
		WithDefaultOption(current).
		WithDefaultText(label).
		Show()

	if err != nil {
		var zero T
		return zero, false
	}

	return itemMap[selectedKey], true
}

func SelectMany[T any](items []T, label string, getKey func(T) string) ([]T, bool) {
	var options []string
	itemMap := make(map[string]T)

	for _, item := range items {
		key := getKey(item)
		options = append(options, key)
		itemMap[key] = item
	}

	selectedKeys, err := pterm.DefaultInteractiveMultiselect.
		WithOptions(options).
		WithDefaultText(label).
		WithKeyConfirm(keys.Enter).
		WithKeySelect(keys.Space).
		Show()

	if err != nil {
		return nil, false
	}

	var selectedItems []T
	for _, key := range selectedKeys {
		selectedItems = append(selectedItems, itemMap[key])
	}

	return selectedItems, true
}
