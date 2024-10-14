package iterator

type UserIterator struct {
	Users []*User
	Index int
}

func (u *UserIterator) HasNext() bool {
	return u.Index < len(u.Users)

}

func (u *UserIterator) GetNext() *User {
	if u.HasNext() {
		user := u.Users[u.Index]
		u.Index++
		return user
	}
	return nil
}
