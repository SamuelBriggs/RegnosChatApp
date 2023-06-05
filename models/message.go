package models

import "time"

type Message struct {
	chat      Chats
	messageId int
	sender    User
	receiver  User
	text      string
	read      bool
	timestamp time.Time
}

func (m *Message) MessageId() int {
	return m.messageId
}

func (m *Message) SetMessageId(messageId int) {
	m.messageId = messageId
}

func (m *Message) Sender() User {
	return m.sender
}

func (m *Message) SetSender(sender User) {
	m.sender = sender
}

func (m *Message) Receiver() User {
	return m.receiver
}

func (m *Message) SetReceiver(receiver User) {
	m.receiver = receiver
}

func (m *Message) Text() string {
	return m.text
}

func (m *Message) SetText(text string) {
	m.text = text
}

func (m *Message) Read() bool {
	return m.read
}

func (m *Message) SetRead(read bool) {
	m.read = read
}

func (m *Message) Timestamp() time.Time {
	return m.timestamp
}

func (m *Message) SetTimestamp(timestamp time.Time) {
	m.timestamp = timestamp
}
