package views

type AdminGetUserResponse struct {
	AdminUserReadView
}

type AdminUserReadView struct {
	ID       uint
	Username string
}

type AdminUserCreateRequest struct {
	Username string
	Password string
}

type AdminUserCreateResponse struct {
	AdminUserReadView
}
