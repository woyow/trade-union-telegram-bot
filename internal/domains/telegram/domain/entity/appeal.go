package entity

type DeleteDraftAppealRepoDTO struct {
	ChatID int64
}

type CreateAppealRepoDTO struct {
	ChatID  int64
	IsDraft bool
}

type CreateAppealOut struct {
	ID string
}

type UpdateAppealBase struct {
	Fname   *string
	Lname   *string
	Mname   *string
	IsDraft *bool
}

type UpdateAppealRepoDTO struct {
	UpdateAppealBase
	ChatID int64
}

type GetDraftAppealRepoDTO struct {
	ChatID int64
}

type GetDraftAppealRepoOut struct {
	Fname string `bson:"fname"`
	Lname string `bson:"lname"`
	Mname string `bson:"mname"`
}
