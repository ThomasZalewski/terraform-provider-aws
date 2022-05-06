package cognitoidp

import "time"

const (
	ResIdentityProvider = "Identity Provider"
	ResResourceServer   = "Resource Server"
	ResUserGroup        = "User Group"
	ResUserPoolClient   = "User Pool Client"
	ResUserPoolDomain   = "User Pool Domain"
	ResUserPool         = "User Pool"
	ResUser             = "User"
)

const (
	iamPropagationTimeout = 2 * time.Minute
)
