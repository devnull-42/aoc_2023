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

	lowPulseCount := 0
	highPulseCount := 0
	cycleCount := 1
	cycleMap := make(map[int][2]int)

PulseLoop:
	for {
		// loop while pulseChan is not empty
		select {
		case pulse := <-pulseChan:
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
			for _, m := range moduleMap {
				if !m.OriginalState() {
					// if a module is not in it's original state then we need to send another pulse
					pulseChan <- Pulse{"button", "broadcaster", false}
					cycleMap[cycleCount] = [2]int{lowPulseCount, highPulseCount}
					cycleCount++
					if cycleCount > 10 {
						break PulseLoop
					}
					continue PulseLoop
				}
			}
			break PulseLoop
		}
	}

	cycles := 1000 / cycleCount
	remainder := 1000 % cycleCount

	totalLowPulses := cycles * lowPulseCount
	totalHighPulses := cycles * highPulseCount

	if remainder > 0 {
		totalLowPulses += cycleMap[remainder][0]
		totalHighPulses += cycleMap[remainder][1]
	}

	fmt.Printf("cycleMap: %+v\n", cycleMap)
	fmt.Printf("cycleCount: %d\n", cycleCount)
	fmt.Printf("cycles: %d, remainder: %d\n", cycles, remainder)
	fmt.Printf("lowPulseCount: %d, highPulseCount: %d\n", lowPulseCount, highPulseCount)
	fmt.Printf("totalLowPulses: %d, totalHighPulses: %d\n", totalLowPulses, totalHighPulses)

	result := totalLowPulses * totalHighPulses
	return result
}

func partB(lines []string) int {
	return 0
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
