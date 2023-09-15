package travel2

import (
	"fmt"
	"sync"
	"time"

	"tugboat/components/common"
	"tugboat/defs"
	"tugboat/logger"
)

func init() {
	name := "comp_travel2"
	defs.ComponentAvailable[name] = CreateComponent(name)
}

type ComponentTravel2 struct {
	common.Component
}

func CreateComponent(name string) *ComponentTravel2 {
	comp := &ComponentTravel2{}
	comp.Active = true
	comp.Name = name
	comp.Options = defs.OpCmd{}
	comp.SignalChan = nil
	comp.Data = []string{}
	comp.Timeout = defs.MAX_TIMEOUT
	comp.Modules = []defs.CModule{}
	comp.Tlog = nil
	return comp
}

func (component *ComponentTravel2) SetLogger(tlog *logger.LogManager) {
	component.Tlog = tlog
}
func (component *ComponentTravel2) GetName() string {
	return component.Name
}

func (component *ComponentTravel2) GetModules() []defs.CModule {
	return component.Modules
}

func (component *ComponentTravel2) SetModules(modules []defs.CModule) {
	component.Modules = make([]defs.CModule, len(modules))
	copy(component.Modules, modules)
}
func (component *ComponentTravel2) SetCmdOptions(op *defs.OpCmd) {
	component.Options.Directive = op.Directive
	component.Options.DirectiveOpts = op.DirectiveOpts
}
func (component *ComponentTravel2) SetCmdDir(opDir string) {
	component.Options.Directive = opDir
}
func (component *ComponentTravel2) SetCmdDirOpt(opDirOpts []string) {
	component.Options.DirectiveOpts = opDirOpts
}

func (component *ComponentTravel2) SetActive(active bool) {
	component.Active = active
}
func (component *ComponentTravel2) SetData(data []string) {
	component.Data = data
}
func (component *ComponentTravel2) SetTimeout(timeoutIn uint) {
	if timeoutIn > 0 {
		component.Timeout = timeoutIn
	}
}
func (component *ComponentTravel2) SetSignalChan(signal chan bool) {
	component.SignalChan = signal
}

func (component *ComponentTravel2) InvokeComponent(
	wg *sync.WaitGroup, msg defs.Message, handoffTo []string) {
	defer wg.Done()

	resCh := make(chan bool, 1)
	go func() {
		res := component.invoke(msg, handoffTo)
		resCh <- res
	}()

	var message string
	select {
	case res := <-resCh:
		message = fmt.Sprintf("Processing returned on its own: %t", res)
		component.Tlog.Log(component.Name, "INFO", message)

	case <-time.After(time.Duration(component.Timeout * uint(time.Second))):
		message = fmt.Sprintf("Processing out of time internally")
		component.Tlog.Log(component.Name, "INFO", message)

	case signalC := <-component.SignalChan:
		message = fmt.Sprintf("Processing received signal: %v", signalC)
		component.Tlog.Log(component.Name, "INFO", message)
		message = fmt.Sprintf("Partial output so far: \nsOut: %s\n sErr: %s", component.Sout, component.Serr)
		component.Tlog.Log(component.Name, "INFO", message)
	}

}
