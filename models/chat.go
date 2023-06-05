package models

import "time"

type Chats struct {
	chatsId        int
	members        []User
	lastMessages   []Message
	unreadMessages int
	timeStamp      time.Time
}

func (c *Chats) ChatsId() int {
	return c.chatsId
}

func (c *Chats) SetChatsId(chatsId int) {
	c.chatsId = chatsId
}

func (c *Chats) Members() []User {
	return c.members
}

func (c *Chats) SetMembers(members []User) {
	c.members = members
}

func (c *Chats) LastMessages() []Message {
	return c.lastMessages
}

func (c *Chats) SetLastMessages(lastMessages []Message) {
	c.lastMessages = lastMessages
}

func (c *Chats) UnreadMessages() int {
	return c.unreadMessages
}

func (c *Chats) SetUnreadMessages(unreadMessages int) {
	c.unreadMessages = unreadMessages
}

func (c *Chats) TimeStamp() time.Time {
	return c.timeStamp
}

func (c *Chats) SetTimeStamp(timeStamp time.Time) {
	c.timeStamp = timeStamp
}
