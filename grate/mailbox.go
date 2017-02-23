package grate

type Message struct {
	X, Y, R float64
	Data string
}

type Mailbox struct {
	Messages []Message
	Incoming []Message
}

func (mailbox *Mailbox) Load() {
	
}

func (mailbox *Mailbox) Update() bool {
	mailbox.Messages, mailbox.Incoming = mailbox.Incoming, mailbox.Messages
	mailbox.Incoming = mailbox.Incoming[:0]
	
	return false
}

func (mailbox *Mailbox) Draw() {
	
}

func (mailbox *Mailbox) SendMessage(data string, x, y, r float64) {
	mailbox.Incoming = append(mailbox.Incoming, Message{Data:data, X:x, Y:y, R:r})
}

func (mailbox *Mailbox) GetMessage(x, y, r float64) Message {
	for _, mail := range mailbox.Messages {
		if (mail.X-x)*(mail.X-x) + (mail.Y-y)*(mail.Y-y) < r*r {
			return mail
		}
	}
	return Message{}
}
