package ssh

type msgExitStatus struct {
	Status uint32
}

type msgExitSignal struct {
	Signal       Signal
	CoreDumped   bool
	ErrorMessage string
	LanguageTag  string
}

type msgPtyReq struct {
	Term   string
	Cols   uint32
	Rows   uint32
	Width  uint32
	Height uint32
	Mode   string
}

func (msg *msgPtyReq) PTY() Pty {
	return Pty{
		Term: msg.Term,
		Window: Window{
			Cols:   msg.Cols,
			Rows:   msg.Rows,
			Width:  msg.Width,
			Height: msg.Height,
		},
		Modes: []byte(msg.Mode),
	}
}

type msgExec struct {
	Command string
}

type msgEnv struct {
	Key   string
	Value string
}

type msgSignal struct {
	Signal Signal
}

type msgForwarding struct {
	DestAddr   string
	DestPort   uint32
	OriginAddr string
	OriginPort uint32
}

type msgRemoteForward struct {
	BindAddr string
	BindPort uint32
}

type msgRemoteForwardSuccess struct {
	BindPort uint32
}
