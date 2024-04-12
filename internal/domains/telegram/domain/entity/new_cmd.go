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

type NewCommandConfirmationStateServiceDTO struct {
	HandleCallback
}
