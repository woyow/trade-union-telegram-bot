package entity

type NewCommandServiceDTO struct {
	HandleCommand
}

type NewCommandFirstNameStateServiceDTO struct {
	HandleMessage
}

type NewCommandLastNameStateServiceDTO struct {
	HandleMessage
}

type NewCommandMiddleNameStateServiceDTO struct {
	HandleMessage
}

type NewCommandSubjectStateServiceDTO struct {
	HandleCallback
}

type NewCommandConfirmationStateServiceDTO struct {
	HandleCallback
}
