package prompt

import (
	"bufio"
	"errors"
	"os"
	"strings"

	"github.com/kiasaki/term"
)

// ErrorPromptAborted is returned by `Prompt()` when CtrlC is pressed
var ErrorPromptAborted = errors.New("Prompt aborted")

// ErrorPromptEnded is returned by `Prompt()` when CtrlD is pressed
var ErrorPromptEnded = errors.New("Prompt ended")

// Prompt represents a current instance of a prompt and it's associated
// history and completion callbacks
type Prompt struct {
	in       *bufio.Reader
	history  []string
	terminal *term.Terminal
}

// NewPrompt creates a new instance of a prompt
func NewPrompt() *Prompt {
	return &Prompt{
		in:       bufio.NewReader(os.Stdin),
		history:  []string{},
		terminal: term.NewTerminal(),
	}
}

// History returns the current history as a `\n` separated string
func (p *Prompt) History() string {
	return strings.Join(p.history, "\n")
}

// LoadHistory replaces the current history with the contents of the
// provided `\n` separated string
func (p *Prompt) LoadHistory(history string) {
	buf := ""
	for r := range history {
		if r == '\n' {
			p.history = append(p.history, buf)
			buf = ""
		}
		buf += string(r)
	}
}

// AppendHistory adds a line to the prompt's history
func (p *Prompt) AppendHistory(line string) {
	p.history = append(p.history, line)
}

// Prompt puts the terminal in line editing mode and waits for the user to
// enter some text followed by the `Enter` key.
func (p *Prompt) Prompt(prompt string) (string, error) {
	p.terminal.Start()
	defer p.terminal.Stop()
	p.terminal.Puts(prompt)

	line := ""
	for {
		ev := <-p.terminal.Events()
		if ev.Type == term.EventKey {
			if ev.Key == term.KeyCtrlC {
				return "", ErrorPromptAborted
			}
			if ev.Key == term.KeyCtrlD {
				return "", ErrorPromptEnded
			}
			if ev.Key == term.KeyCr {
				p.terminal.Puts("\n")
				p.terminal.SetCursorColumn(0)
				break
			}
			if ev.Key == term.KeyBackspace {
				if len(line) > 0 {
					line = line[:len(line)-1]
				}
			} else if ev.Key == term.KeyCtrlU {
				line = ""
			} else if ev.Key == term.KeyCtrlL {
				p.terminal.Clear()
			} else if ev.Key == term.KeyRune {
				line += string(ev.Rune)
			} else {
				// TODO remove debug
				// p.terminal.Puts("\n")
				// p.terminal.SetCursorColumn(0)
				// p.terminal.Puts("unknown key: ", int(ev.Key), ev.Rune)
			}
		}
		p.terminal.SetCursorColumn(0)
		p.terminal.Puts(strings.Repeat(" ", p.terminal.Width))
		p.terminal.SetCursorColumn(0)
		p.terminal.Puts(prompt + line)
	}

	return string(line), nil
}
