package views

type AdminGetUserResponse struct {
	AdminUserReadView
}

type AdminUserReadView struct {
	ID        uint
	Email     string
	FirstName string
	LastName  string
	Role      string
}

type AdminUserCreateRequest struct {
	Email    string
	Role     string
	Password string
}

type AdminUserCreateResponse struct {
	AdminUserReadView
}
