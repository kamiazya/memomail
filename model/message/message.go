package message

type Message []byte

func (m *Message) String() string {
	return string(*m)
}
