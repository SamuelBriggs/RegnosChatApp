package models

type User struct {
	userId      int
	firstName   string
	lastName    string
	phoneNumber string
	userName    string
	password    string
}

func (u *User) UserId() int {
	return u.userId
}

func (u *User) SetUserId(userId int) {
	u.userId = userId
}

func (u *User) FirstName() string {
	return u.firstName
}

func (u *User) SetFirstName(firstName string) {
	u.firstName = firstName
}

func (u *User) LastName() string {
	return u.lastName
}

func (u *User) SetLastName(lastName string) {
	u.lastName = lastName
}

func (u *User) PhoneNumber() string {
	return u.phoneNumber
}

func (u *User) SetPhoneNumber(phoneNumber string) {
	u.phoneNumber = phoneNumber
}

func (u *User) UserName() string {
	return u.userName
}

func (u *User) SetUserName(userName string) {
	u.userName = userName
}

func (u *User) Password() string {
	return u.password
}

func (u *User) SetPassword(password string) {
	u.password = password
}
