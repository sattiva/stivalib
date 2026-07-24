package auth

type Scope uint64

const (
	ScopeRead Scope = 1 << iota
	ScopeWrite
	ScopeAdmin
	ScopeDeploy
	ScopeExecute
)

func HasScope(userScopes, requiredScope Scope) bool {
	return (userScopes & requiredScope) == requiredScope
}

func AddScope(userScopes, newScope Scope) Scope {
	return userScopes | newScope
}

func RemoveScope(userScopes, remScope Scope) Scope {
	return userScopes &^ remScope
}
