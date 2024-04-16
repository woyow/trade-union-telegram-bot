package entity

type NewCommandServiceDTO struct {
	HandleCommand
}

type NewCommandFullNameStateServiceDTO struct {
	HandleMessage
}

type NewCommandSubjectStateServiceDTO struct {
	HandleCallback
}

type NewCommandConfirmationStateServiceDTO struct {
	HandleCallback
}
