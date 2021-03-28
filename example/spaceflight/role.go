package spaceflight

type Role struct {
	*user
}

func (me *Role) setUser(v *user) { me.user = v }

type Pilot Role
type Passenger Role

// ----------------------------------------

// user implements domain logic for operations on the service and
// should be used through a role.
type user struct {
	srv *Service
}

func (me *user) setService(v *Service) { me.srv = v }

// ----------------------------------------

// PlanRoute stores the given route. Fails if route already exists.
func (me *Pilot) PlanRoute(v Route) error { return me.planRoute(v) }

func (me *user) planRoute(route Route) error {
	// implement...
	return nil
}

// ----------------------------------------

// ListRoutes returns all routes in the system.
func (me *Pilot) ListRoutes() []Route     { return me.listRoutes() }
func (me *Passenger) ListRoutes() []Route { return me.listRoutes() }

func (me *user) listRoutes() []Route {
	// implement list routes
	return nil
}
