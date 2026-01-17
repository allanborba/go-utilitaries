package profiler

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var profiler *Profiler

func SetProfiler(rootName string) {
	profiler = New(rootName, true)
}

type measurement struct {
	Name     string         `json:"name"`
	Duration time.Duration  `json:"-"`
	Children []*measurement `json:"children,omitempty"`

	startedAt time.Time
	finished  bool
}

type outputMeasurement struct {
	Name       string              `json:"name"`
	DurationMs float64             `json:"duration_ms"`
	Children   []outputMeasurement `json:"children,omitempty"`
}

type Profiler struct {
	enabled   bool
	root      *measurement
	stack     []*measurement
	startedAt time.Time
	finished  bool
	mu        sync.Mutex
}

func New(rootName string, enabled bool) *Profiler {
	if !enabled {
		return &Profiler{enabled: false}
	}

	root := &measurement{
		Name:      rootName,
		startedAt: time.Now(),
	}

	return &Profiler{
		enabled:   true,
		root:      root,
		stack:     []*measurement{root},
		startedAt: root.startedAt,
	}
}

func (p *Profiler) Enabled() bool {
	return p != nil && p.enabled
}

func (p *Profiler) Start(name string) func() {
	if p == nil || !p.enabled {
		return func() {}
	}

	m := &measurement{Name: name, startedAt: time.Now()}

	p.mu.Lock()
	if len(p.stack) == 0 {
		p.stack = []*measurement{p.root}
	}
	parent := p.stack[len(p.stack)-1]
	parent.Children = append(parent.Children, m)
	p.stack = append(p.stack, m)
	p.mu.Unlock()

	return func() {
		if p == nil || !p.enabled {
			return
		}

		p.mu.Lock()
		if !m.finished {
			m.Duration = time.Since(m.startedAt)
			m.finished = true
		}

		if l := len(p.stack); l > 0 && p.stack[l-1] == m {
			p.stack = p.stack[:l-1]
		} else {
			for i := len(p.stack) - 1; i >= 0; i-- {
				if p.stack[i] == m {
					p.stack = append(p.stack[:i], p.stack[i+1:]...)
					break
				}
			}
		}
		p.mu.Unlock()
	}
}

func (p *Profiler) Track(name string, fn func()) {
	if p == nil || !p.enabled {
		fn()
		return
	}

	stop := p.Start(name)
	defer stop()
	fn()
}

func (p *Profiler) Finish() {
	if p == nil || !p.enabled {
		return
	}

	p.mu.Lock()
	if p.finished {
		p.mu.Unlock()
		return
	}

	now := time.Now()

	for len(p.stack) > 1 {
		pending := p.stack[len(p.stack)-1]
		if !pending.finished {
			pending.Duration = now.Sub(pending.startedAt)
			pending.finished = true
		}
		p.stack = p.stack[:len(p.stack)-1]
	}

	if p.root != nil && !p.root.finished {
		p.root.Duration = now.Sub(p.startedAt)
		p.root.finished = true
	}

	p.finished = true
	p.mu.Unlock()
}

func (p *Profiler) WriteToFile(dir, prefix string) (string, error) {
	if p == nil || !p.enabled {
		return "", nil
	}

	p.Finish()

	if err := os.MkdirAll(dir, 0o755); err != nil {
		return "", err
	}

	timestamp := time.Now().UTC().Format("20060102T150405.000Z")
	filename := fmt.Sprintf("%s_%s.json", timestamp, prefix)
	path := filepath.Join(dir, filename)

	payload := p.buildOutput()
	content, err := json.MarshalIndent(payload, "", "  ")
	if err != nil {
		return "", err
	}

	if err := os.WriteFile(path, content, 0o644); err != nil {
		return "", err
	}

	return path, nil
}

func (p *Profiler) buildOutput() outputMeasurement {
	if p == nil || p.root == nil {
		return outputMeasurement{}
	}

	return toOutput(p.root)
}

func toOutput(m *measurement) outputMeasurement {
	if m == nil {
		return outputMeasurement{}
	}

	children := make([]outputMeasurement, len(m.Children))
	for i, child := range m.Children {
		children[i] = toOutput(child)
	}

	return outputMeasurement{
		Name:       m.Name,
		DurationMs: float64(m.Duration.Microseconds()) / 1000.0,
		Children:   children,
	}
}

func Track(name string, fn func()) {
	if profiler == nil || !profiler.Enabled() {
		fn()
		return
	}

	stop := profiler.Start(name)
	defer stop()
	fn()
}

func WriteFile(dir, prefix string) {
	profiler.WriteToFile(dir, prefix)
}
