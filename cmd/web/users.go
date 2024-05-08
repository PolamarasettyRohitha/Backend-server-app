package main

type user struct {
	name     string
	password string
	admin    bool
}

var regularUsers = map[string]user{
	"user":  user{name: "user", password: "user@123"},
	"user1": user{name: "user1", password: "user@123"},
}

var adminUsers = map[string]user{
	"admin":  user{name: "admin", password: "admin@123", admin: true},
	"admin1": user{name: "admin1", password: "admin1@123", admin: true},
}

func isValidUser(username, password string) (user, bool) {
	if val, ok := regularUsers[username]; ok && val.password == password {
		return val, true
	}

	if val, ok := adminUsers[username]; ok && val.password == password {
		return val, true
	}

	return user{}, false
}
