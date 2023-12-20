package day20

import (
	"aoc/util"
	"fmt"
	"regexp"
	"slices"
	"strings"
)

func Run() {
	lines := util.ReadInput("day20.txt")
	result := partA(lines)
	fmt.Printf("partA: %d\n", result)

	result = partB(lines)
	fmt.Printf("partB: %d\n", result)
}

func partA(lines []string) int {
	pulseChan := make(chan Pulse, 200)
	defer close(pulseChan)

	moduleMap := BuildModules(lines, pulseChan)

	// send initial pulse
	pulseChan <- Pulse{"button", "broadcaster", false}
	buttonPresses := 1

	lowPulseCount := 0
	highPulseCount := 0

PulseLoop:
	for {
		// loop while pulseChan is not empty
		select {
		case pulse := <-pulseChan:
			if pulse.dest == "rx" && !pulse.strength {
				return buttonPresses
			}
			if pulse.strength {
				highPulseCount++
			} else {
				lowPulseCount++
			}
			if _, exists := moduleMap[pulse.dest]; !exists {
				continue
			}
			moduleMap[pulse.dest].Receive(pulse)
		default:
			if buttonPresses < 1000 {
				pulseChan <- Pulse{"button", "broadcaster", false}
				buttonPresses++
				continue PulseLoop
			}
			break PulseLoop
		}
	}

	result := lowPulseCount * highPulseCount
	return result
}

func partB(lines []string) int {
	pulseChan := make(chan Pulse, 200)
	defer close(pulseChan)

	moduleMap := BuildModules(lines, pulseChan)
	monitor := getMonitor(moduleMap)

	// send initial pulse
	pulseChan <- Pulse{"button", "broadcaster", false}
	buttonPresses := 1

	for {
		// loop while pulseChan is not empty
		select {
		case pulse := <-pulseChan:
			if pulse.dest == "nc" && pulse.strength {
				monitor.AddInput(pulse.source, buttonPresses)
				if monitor.Check(buttonPresses, moduleMap[monitor.name]) {
					return monitor.GetLCM()
				}
			}
			if _, exists := moduleMap[pulse.dest]; !exists {
				continue
			}
			moduleMap[pulse.dest].Receive(pulse)
		default:
			// fmt.Printf("buttonPresses: %d\n", buttonPresses)
			// fmt.Printf("monitor: %s\n", monitor.String())
			pulseChan <- Pulse{"button", "broadcaster", false}
			buttonPresses++

		}
	}
}

type Pulse struct {
	source   string
	dest     string
	strength bool
}

func (p Pulse) String() string {
	var level string
	if p.strength {
		level = "high"
	} else {
		level = "low"
	}
	return fmt.Sprintf("%s -%s-> %s", p.source, level, p.dest)
}

type Module interface {
	Receive(input Pulse)
	OriginalState() bool
	GetOutput() []string
	AddInput(input string)
	GetInputs() map[string]bool
}

type FlipFlop struct {
	name   string
	state  bool
	output []string
	send   chan Pulse
}

func NewFlipFlop(name string, send chan Pulse, output []string) *FlipFlop {
	return &FlipFlop{name: name, output: output, send: send}
}

func (f *FlipFlop) Receive(input Pulse) {
	if input.strength {
		return
	}
	f.state = !f.state
	for _, o := range f.output {
		f.send <- Pulse{f.name, o, f.state}
	}
}

func (f *FlipFlop) OriginalState() bool {
	return !f.state
}

func (f *FlipFlop) GetOutput() []string {
	return f.output
}

func (f *FlipFlop) AddInput(input string) {}

func (f *FlipFlop) GetInputs() map[string]bool {
	return nil
}

type Conjunction struct {
	name   string
	inputs map[string]bool
	output []string
	send   chan Pulse
}

func NewConjunction(name string, send chan Pulse, output []string) *Conjunction {
	return &Conjunction{name: name, inputs: map[string]bool{}, output: output, send: send}
}

func (c *Conjunction) Receive(input Pulse) {
	c.inputs[input.source] = input.strength
	state := c.State()
	for _, o := range c.output {
		c.send <- (Pulse{c.name, o, state})
	}
}

func (c *Conjunction) State() bool {
	for _, v := range c.inputs {
		if !v {
			return true
		}
	}
	return false
}

func (c *Conjunction) OriginalState() bool {
	for _, v := range c.inputs {
		if v {
			return false
		}
	}
	return true
}

func (c *Conjunction) GetOutput() []string {
	return c.output
}

func (c *Conjunction) AddInput(input string) {
	c.inputs[input] = false
}

func (c *Conjunction) GetInputs() map[string]bool {
	return c.inputs
}

type Broadcast struct {
	name   string
	output []string
	send   chan Pulse
}

func NewBroadcast(name string, send chan Pulse, output []string) *Broadcast {
	return &Broadcast{name: name, output: output, send: send}
}

func (b *Broadcast) Receive(input Pulse) {
	for _, o := range b.output {
		b.send <- Pulse{b.name, o, input.strength}
	}
}

func (b *Broadcast) OriginalState() bool {
	return true
}

func (b *Broadcast) GetOutput() []string {
	return b.output
}

func (b *Broadcast) AddInput(input string) {}

func (b *Broadcast) GetInputs() map[string]bool {
	return nil
}

func BuildModules(lines []string, send chan Pulse) map[string]Module {
	modules := make(map[string]Module)
	conjunctionModules := make([]string, 0)
	for _, line := range lines {
		moduleDef := strings.Split(line, " -> ")
		moduleOutput := strings.Split(moduleDef[1], ", ")

		reModule := regexp.MustCompile(`(broadcaster|%|&)(\w*)`)
		matches := reModule.FindStringSubmatch(moduleDef[0])

		switch matches[1] {
		case "broadcaster":
			modules["broadcaster"] = NewBroadcast("broadcaster", send, moduleOutput)
		case "%":
			modules[matches[2]] = NewFlipFlop(matches[2], send, moduleOutput)
		case "&":
			conjunctionModules = append(conjunctionModules, matches[2])
			modules[matches[2]] = NewConjunction(matches[2], send, moduleOutput)
		}
	}

	for _, c := range conjunctionModules {
		for k, m := range modules {
			if slices.Contains(m.GetOutput(), c) {
				modules[c].AddInput(k)
			}
		}
	}
	return modules
}

// part b functions

func getMonitor(modules map[string]Module) Monitor {
	var monitor Monitor
	for k, m := range modules {
		if slices.Contains(m.GetOutput(), "rx") {
			monitor = Monitor{k, make(map[string]int)}
		}
	}

	for k, m := range modules {
		if slices.Contains(m.GetOutput(), monitor.name) {
			monitor.inputs[k] = 0
		}
	}
	return monitor
}

type Monitor struct {
	name   string
	inputs map[string]int
}

func (m Monitor) Check(buttonPresses int, module Module) bool {
	for _, v := range m.inputs {
		if v == 0 {
			return false
		}
	}
	return true
}

func (m Monitor) AddInput(source string, buttonPresses int) {
	if v, exists := m.inputs[source]; exists {
		if v == 0 {
			m.inputs[source] = buttonPresses
		}
	}
}

func (m Monitor) GetLCM() int {
	buttonPresses := make([]int, 0)
	for _, v := range m.inputs {
		buttonPresses = append(buttonPresses, v)
	}
	return util.LcmMultiple(buttonPresses...)
}

func (m Monitor) String() string {
	return fmt.Sprintf("Monitor(%s, %+v)", m.name, m.inputs)
}
