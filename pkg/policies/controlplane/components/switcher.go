package components

import (
	"go.uber.org/fx"

	policylangv1 "github.com/fluxninja/aperture/v2/api/gen/proto/go/aperture/policy/language/v1"
	"github.com/fluxninja/aperture/v2/pkg/config"
	"github.com/fluxninja/aperture/v2/pkg/notifiers"
	"github.com/fluxninja/aperture/v2/pkg/policies/controlplane/iface"
	"github.com/fluxninja/aperture/v2/pkg/policies/controlplane/runtime"
	"github.com/fluxninja/aperture/v2/pkg/policies/controlplane/runtime/tristate"
)

// Switcher switches between two inputs based on third input.
type Switcher struct{}

// Name implements runtime.Component.
func (*Switcher) Name() string { return "Switcher" }

// Type implements runtime.Component.
func (*Switcher) Type() runtime.ComponentType { return runtime.ComponentTypeSignalProcessor }

// ShortDescription implements runtime.Component.
func (*Switcher) ShortDescription() string { return "" }

// IsActuator implements runtime.Component.
func (*Switcher) IsActuator() bool { return false }

// Make sure Switcher complies with Component interface.
var _ runtime.Component = (*Switcher)(nil)

// NewSwitcherAndOptions creates a new Switcher Component.
func NewSwitcherAndOptions(_ *policylangv1.Switcher, _ runtime.ComponentID, _ iface.Policy) (runtime.Component, fx.Option, error) {
	switcher := Switcher{}
	return &switcher, fx.Options(), nil
}

// Execute implements runtime.Component.Execute.
func (dec *Switcher) Execute(inPortReadings runtime.PortToReading, circuitAPI runtime.CircuitAPI) (runtime.PortToReading, error) {
	onSignal := inPortReadings.ReadSingleReadingPort("on_signal")
	offSignal := inPortReadings.ReadSingleReadingPort("off_signal")
	switchValue := inPortReadings.ReadSingleReadingPort("switch")

	var output runtime.Reading

	if tristate.FromReading(switchValue).IsTrue() {
		output = onSignal
	} else {
		output = offSignal
	}

	return runtime.PortToReading{
		"output": []runtime.Reading{output},
	}, nil
}

// DynamicConfigUpdate is a no-op for Switcher.
func (dec *Switcher) DynamicConfigUpdate(event notifiers.Event, unmarshaller config.Unmarshaller) {}
