package main


import "errors"

type command struct {
	name string
	args []string
}

type commands struct {
	namedHandlers map[string]func(*state, *command) error
}

func (cmds *commands) register(name string, handler func(*state, *command) error) {
	if cmds.namedHandlers == nil {
		cmds.namedHandlers = make(map[string]func(*state, *command) error)
	}
	cmds.namedHandlers[name] = handler
}

func (cmds *commands) run(s *state, cmd *command) error {
	handler, ok := cmds.namedHandlers[cmd.name]
	if !ok {
		return errors.New("No such command registered: " + cmd.name)
	}
	return handler(s, cmd)
}
