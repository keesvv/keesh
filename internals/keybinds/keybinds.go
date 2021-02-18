package keybinds

import "fmt"

var bindsMap = map[byte]string{}

func IsKeybinding(key byte) bool {
	_, ok := bindsMap[key]
	return ok
}

func GetCommand(key byte) string {
	return bindsMap[key]
}

func RegisterBinding(key byte, cmd string) {
	bindsMap[key] = cmd
	fmt.Println(bindsMap)
}
