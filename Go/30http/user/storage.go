package User

type Storage map[int]*User

func (u *Storage) Search(id int) (us *User, ok bool) {
	if us, ok = (*u)[id]; ok {
		return
	}
	return
}
