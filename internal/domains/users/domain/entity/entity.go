package entity

type CreateUserServiceDTO struct {
	Roles    []string
	Fname    string
	Lname    string
	Mname    string
	Position string
	ChatID   int64
}

type CreateUserRepoDTO struct {
	Roles    []string
	Fname    string
	Lname    string
	Mname    string
	Position string
	ChatID   int64
}

type CreateUserOut struct {
	ID string `json:"id"`
}

type GetUserServiceDTO struct {
	ID     *string
	ChatID *int64
}

type GetUserRepoDTO struct {
	ID     *string
	ChatID *int64
}

type GetUserOut struct {
	Roles    []string `json:"roles" bson:"roles"`
	ID       string   `json:"id" bson:"_id"`
	Fname    string   `json:"fname" bson:"fname"`
	Lname    string   `json:"lname" bson:"lname"`
	Mname    string   `json:"mname" bson:"mname"`
	Position string   `json:"position" bson:"position"`
	ChatID   int64    `json:"chat_id" bson:"chatId"`
}