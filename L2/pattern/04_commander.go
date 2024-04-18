package pattern

type Command interface {
	Execute() string
}

type LightsOnCommand struct {
	receiver *Lights
}

func (c *LightsOnCommand) Execute() string {
	return c.receiver.TurnOn()
}

type LightsOffCommand struct {
	receiver *Lights
}

func (c *LightsOffCommand) Execute() string {
	return c.receiver.TurnOff()
}

type Lights struct {
}

func (l *Lights) TurnOn() string {
	return "Фары включены"
}

func (l *Lights) TurnOff() string {
	return "Фары выключены"
}

type SoundSystemVolumeUpCommand struct {
	receiver *SoundSystem
}

func (c *SoundSystemVolumeUpCommand) Execute() string {
	return c.receiver.VolumeUp()
}

type SoundSystemVolumeDownCommand struct {
	receiver *SoundSystem
}

func (c *SoundSystemVolumeDownCommand) Execute() string {
	return c.receiver.VolumeDown()
}

type SoundSystem struct {
}

func (s *SoundSystem) VolumeUp() string {
	return "Громкость увеличена"
}

func (s *SoundSystem) VolumeDown() string {
	return "Громкость уменьшена"
}

type Invoker struct {
	commands []Command
}

func (i *Invoker) StoreCommand(command Command) {
	i.commands = append(i.commands, command)
}

func (i *Invoker) UnStoreCommand() {
	if len(i.commands) > 0 {
		i.commands = i.commands[:len(i.commands)-1]
	}
}

func (i *Invoker) Execute() string {
	var result string
	for _, command := range i.commands {
		result += command.Execute() + "\n"
	}
	return result
}
