package model

type InsertUserReq struct {
	FullName    string
	PhoneNumber string
	Password    string
	Salt        string
}

type User struct {
	UserID      int64
	FullName    string
	PhoneNumber string
	LoginCount  int64
}

type GetUserPasswordReq struct {
	PhoneNumber string
}

type GetUserPasswordRes struct {
	UserID   int64
	Password string
	Salt     string
}

type AddUserLoginCountReq struct {
	PhoneNumber string
}

type GetUserByUserIDReq struct {
	UserID int64
}

type UpdateUserByUserIDReq struct {
	UserID      int64
	PhoneNumber string
	FullName    string
}
