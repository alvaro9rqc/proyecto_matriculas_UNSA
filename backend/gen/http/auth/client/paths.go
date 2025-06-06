// Code generated by goa v3.21.0, DO NOT EDIT.
//
// HTTP request path constructors for the Auth service.
//
// Command:
// $ goa gen github.com/enrollment/design/api

package client

// MeAuthPath returns the URL path to the Auth service Me HTTP endpoint.
func MeAuthPath() string {
	return "/auth/me"
}

// GoogleLoginAuthPath returns the URL path to the Auth service Google login HTTP endpoint.
func GoogleLoginAuthPath() string {
	return "/auth/google/login"
}

// GoogleLogoutAuthPath returns the URL path to the Auth service Google logout HTTP endpoint.
func GoogleLogoutAuthPath() string {
	return "/auth/google/logout"
}

// GoogleCallbackAuthPath returns the URL path to the Auth service GoogleCallback HTTP endpoint.
func GoogleCallbackAuthPath() string {
	return "/google/callback"
}
