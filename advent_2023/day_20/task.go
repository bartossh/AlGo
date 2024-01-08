package day_20

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

type module struct {
	name   string
	memory map[string]bool
	dests  []string
	cat    rune
	on     bool
}

func (m module) isSwitch() bool {
	return m.cat == '%'
}

func (m module) isConjunction() bool {
	return m.cat == '&'
}

func (m module) isFinal(name string) bool {
	return m.name == name
}

func (m module) isBroadcast() bool {
	return m.cat == 0
}

func (m module) isSwitchOn() (bool, error) {
	if !m.isSwitch() {
		return false, errors.New("it is not a switch module")
	}
	return m.on, nil
}

func (m *module) actWithSwitch(signalRec bool) (action, signalSend bool, err error) {
	if !m.isSwitch() {
		return false, false, errors.New("it is not a switch module")
	}
	if signalRec {
		return false, false, nil
	}
	m.on = !m.on
	switch m.on {
	case true:
		return true, true, nil
	default:
		return true, false, nil
	}
}

func (m *module) addSenderToConjuntion(n string) error {
	if !m.isConjunction() {
		return errors.New("it is not a conjunction module")
	}
	m.memory[n] = false
	return nil
}

func (m module) isAllConjunctionSendersHigh() bool {
	for _, v := range m.memory {
		if !v {
			return false
		}
	}
	return true
}

func (m *module) updateConjunctionSender(n string, signal bool) error {
	if _, ok := m.memory[n]; !ok {
		return errors.New("sender not found in the conjunction module")
	}
	m.memory[n] = signal
	return nil
}

func (m *module) actWithConjunction(n string, signal bool) (bool, error) {
	if !m.isConjunction() {
		return false, errors.New("it is not a conjunction module")
	}

	if err := m.updateConjunctionSender(n, signal); err != nil {
		return false, err
	}

	if m.isAllConjunctionSendersHigh() {
		return false, nil
	}
	return true, nil
}

type grid struct {
	modules            map[string]module
	lowSignalsCounter  int
	highSignalsCounter int
}

func newGrid() *grid {
	return &grid{
		modules: make(map[string]module),
	}
}

func (g *grid) addModule(m module) {
	g.modules[m.name] = m
}

func (g grid) prepare() {
	for _, m := range g.modules {
		for _, dest := range m.dests {
			inner, ok := g.modules[dest]
			if !ok {
				continue
			}
			if inner.isConjunction() {
				inner.addSenderToConjuntion(m.name)
				g.modules[dest] = inner
			}
		}
	}
}

func (g *grid) updateSignal(s bool) {
	switch s {
	case true:
		g.highSignalsCounter++
	default:
		g.lowSignalsCounter++
	}
}

func (g *grid) countSignals() (low, high int) {
	low = g.lowSignalsCounter
	high = g.highSignalsCounter
	return
}

type next struct {
	sender string
	name   string
	signal bool
}

func (g *grid) emulate(start string, signal bool) error {
	g.updateSignal(signal)
	broadcaster := g.modules[start]
	var nextDestinations []next
	for _, name := range broadcaster.dests {
		g.updateSignal(signal)
		m, ok := g.modules[name]
		if !ok {
			continue
		}
		var proceed, sig bool
		var err error
		if m.isConjunction() {
			sig, err = m.actWithConjunction(start, signal)
			if err != nil {
				return err
			}
			proceed = true
			g.modules[name] = m
		}
		if m.isSwitch() {
			proceed, sig, err = m.actWithSwitch(signal)
			if err != nil {
				return err
			}
			g.modules[name] = m
		}

		if proceed {
			nexts := make([]next, 0, len(m.dests))
			for _, n := range m.dests {
				nexts = append(nexts, next{name: n, sender: name, signal: sig})
			}

			nextDestinations = append(nextDestinations, nexts...)
		}
	}

	for len(nextDestinations) > 0 {
		var queuedDestinations []next

		for _, nx := range nextDestinations {
			g.updateSignal(nx.signal)
			m, ok := g.modules[nx.name]
			if !ok {
				continue
			}
			var proceed, sig bool
			var err error
			if m.isConjunction() {
				sig, err = m.actWithConjunction(nx.sender, nx.signal)
				if err != nil {
					return err
				}
				proceed = true
				g.modules[nx.name] = m
			}
			if m.isSwitch() {
				proceed, sig, err = m.actWithSwitch(nx.signal)
				if err != nil {
					return err
				}
				g.modules[nx.name] = m
			}

			if proceed {
				nexts := make([]next, 0, len(m.dests))
				for _, n := range m.dests {
					nexts = append(nexts, next{name: n, sender: nx.name, signal: sig})
				}

				queuedDestinations = append(queuedDestinations, nexts...)
			}
		}
		nextDestinations = queuedDestinations
	}

	return nil
}

func (g *grid) emulateToFinal(start, final string, signal bool) (bool, error) {
	broadcaster := g.modules[start]
	var nextDestinations []next
	for _, name := range broadcaster.dests {
		m, ok := g.modules[name]
		if !ok {
			continue
		}
		if m.isFinal(final) {
			if !signal {
				return true, nil
			}
			continue
		}
		var proceed, sig bool
		var err error
		if m.isConjunction() {
			sig, err = m.actWithConjunction(start, signal)
			if err != nil {
				return false, err
			}
			proceed = true
			g.modules[name] = m
		}
		if m.isSwitch() {
			proceed, sig, err = m.actWithSwitch(signal)
			if err != nil {
				return false, err
			}
			g.modules[name] = m
		}

		if proceed {
			nexts := make([]next, 0, len(m.dests))
			for _, n := range m.dests {
				nexts = append(nexts, next{name: n, sender: name, signal: sig})
			}

			nextDestinations = append(nextDestinations, nexts...)
		}
	}

	for len(nextDestinations) > 0 {
		var queuedDestinations []next

		for _, nx := range nextDestinations {
			m, ok := g.modules[nx.name]
			if !ok {
				continue
			}
			if m.isFinal(final) {
				if !nx.signal {
					return true, nil
				}
				continue
			}
			var proceed, sig bool
			var err error
			if m.isConjunction() {
				sig, err = m.actWithConjunction(nx.sender, nx.signal)
				if err != nil {
					return false, err
				}
				proceed = true
				g.modules[nx.name] = m
			}
			if m.isSwitch() {
				proceed, sig, err = m.actWithSwitch(nx.signal)
				if err != nil {
					return false, err
				}
				g.modules[nx.name] = m
			}

			if proceed {
				nexts := make([]next, 0, len(m.dests))
				for _, n := range m.dests {
					nexts = append(nexts, next{name: n, sender: nx.name, signal: sig})
				}

				queuedDestinations = append(queuedDestinations, nexts...)
			}
		}
		nextDestinations = queuedDestinations
	}

	return false, nil
}

func lineReader(line string) (module, error) {
	arr := strings.Split(line, "->")
	if len(arr) != 2 {
		return module{}, errors.New("lack of '->' separator")
	}
	arr[0] = strings.TrimSpace(arr[0])
	if len(arr) < 2 {
		return module{}, errors.New("cannot decode name")
	}
	destinations := strings.Split(arr[1], ",")
	if len(destinations) == 0 {
		return module{}, errors.New("no destinatins found")
	}

	var name string
	var cat rune

	switch arr[0] {
	case "broadcaster":
		name = arr[0]
	default:
		name = arr[0][1:]
		cat = rune(arr[0][0])
	}

	m := module{
		cat:    cat,
		name:   name,
		on:     false,
		memory: make(map[string]bool),
		dests:  make([]string, 0, len(destinations)),
	}

	for _, dest := range destinations {
		dest = strings.TrimSpace(dest)
		m.dests = append(m.dests, dest)
	}

	return m, nil
}

func Solver1(path string) (int, error) {
	f, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	g := newGrid()

	for scanner.Scan() {
		l := scanner.Text()
		m, err := lineReader(l)
		if err != nil {
			return 0, err
		}
		g.addModule(m)
	}

	g.prepare()
	for i := 0; i < 1000; i++ {
		if err := g.emulate("broadcaster", false); err != nil {
			return 0, err
		}
	}

	low, high := g.countSignals()

	return low * high, nil
}

func Solver2(path, final string) (int, error) {
	f, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	g := newGrid()

	for scanner.Scan() {
		l := scanner.Text()
		m, err := lineReader(l)
		if err != nil {
			return 0, err
		}
		g.addModule(m)
	}
	g.addModule(module{name: final})

	g.prepare()
	var counter int
	for {
		counter++
		ok, err := g.emulateToFinal("broadcaster", final, false)
		if err != nil {
			return 0, err
		}
		if ok {
			break
		}
	}

	return counter, nil
}
