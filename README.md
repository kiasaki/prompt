# prompt

_A simplistic command line prompt library in pure Go._

### intro

`prompt` is small (< 250 lines) command line prompt library. It does not have
all the fancy support and keybindings and config files that `readline` has but
it's still very functional. If you where looking for something light or small
enough to embed in a bigger application and modify to your own needs this is it.

### features

- Showing a custom prompt message
- Basic line editing and keybindings (<key>Backspace</key>, <key>Ctrl-U</key>, <key>Ctrl-L</key>, ...)
- History loading and exporting support + <key>Up</key>/<key>Down</key> keybindings
- Simple completion callback bound to <key>Tab</key>

### api

```
var ErrorPromptAborted = errors.New("Prompt aborted")
var ErrorPromptEnded = errors.New("Prompt ended")
func NewPrompt() *Prompt
func (p *Prompt) History() string
func (p *Prompt) LoadHistory(history string)
func (p *Prompt) AppendHistory(line string)
func (p *Prompt) Prompt(prompt string) (string, error)
```

### license

MIT. See `LICENSE` file.
