package formatter

type citizen struct{}

type Citizen struct {
	Firstname string
	firstname string
}

func (c *Citizen) PrintIDCard() {}

func (c *Citizen) crintNameCard() {}