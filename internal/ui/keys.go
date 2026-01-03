package ui

// KeySet represents a set of keys that trigger the same action.
type KeySet []string

// Contains checks if the key is in the set.
func (k KeySet) Contains(key string) bool {
	for _, v := range k {
		if v == key {
			return true
		}
	}
	return false
}

// Keys contains all keybindings for the TUI.
var Keys = struct {
	Quit     KeySet
	Up       KeySet
	Down     KeySet
	Left     KeySet
	Right    KeySet
	Enter    KeySet
	Back     KeySet
	NextPage KeySet
	PrevPage KeySet
	Refresh  KeySet
	Comments KeySet
	Help     KeySet
}{
	Quit:     KeySet{"q", "ctrl+c"},
	Up:       KeySet{"k", "up"},
	Down:     KeySet{"j", "down"},
	Left:     KeySet{"h", "left"},
	Right:    KeySet{"l", "right"},
	Enter:    KeySet{"enter", " "},
	Back:     KeySet{"esc", "backspace"},
	NextPage: KeySet{"pgdown", "ctrl+d"},
	PrevPage: KeySet{"pgup", "ctrl+u"},
	Refresh:  KeySet{"r"},
	Comments: KeySet{"c"},
	Help:     KeySet{"?"},
}
