package models

//Login model is the model used to log in.
//Its contains the information used to login is more info is added to log in
//you need to add it here.
type Login struct {
	Username string
	Password string
}

//TokenClaims are the private claims stored in the token.
//If private claims are added you need to add them here.
type TokenClaims struct {
	Username string
	Admin    bool
}
