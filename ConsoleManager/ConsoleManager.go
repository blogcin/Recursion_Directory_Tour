package ConsoleManager

import "bufio"

type ConsoleManager struct {
	Reader *bufio.Reader
}

func (cs *ConsoleManager) ReadFromConsole(text *string) {
	buffer, _, _ := cs.Reader.ReadLine()

	*text = string(buffer[:])
}