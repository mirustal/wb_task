package pattern


type MobileAlertStater interface {
	Alert() string
}

type MobileAlert struct {
	state MobileAlertStater
}

func (a *MobileAlert) Alert() string {
	return a.state.Alert()
}

func (a *MobileAlert) SetState(state MobileAlertStater) {
	a.state = state
}

func NewMobileAlert() *MobileAlert {
	return &MobileAlert{state: &CarAlertBeep{}}
}

type CarAlertBeep struct {
}

func (a *CarAlertBeep) Alert() string {
	return "Beep! Beep! Beep!"
}

type CarAlertVibration struct {
}

func (a *CarAlertVibration) Alert() string {
	return "rrrrrrr"
}


type CarAlertVoice struct {
}

// Alert возвращает строку оповещения.
func (a *CarAlertVoice) Alert() string {
	return "I am tipa Alisa"
}
