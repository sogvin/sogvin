package spaceflight

// Service provides logic for reading and writing spaceflight related
// resources.
type Service struct {
	// e.g. database, sync mutexes
}

// Use provides the given role access to the service.
func (me *Service) Use(role Role) {
	var u user
	u.setService(me)
	role.setUser(&u)
}
