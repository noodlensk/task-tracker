package command

type EventPublisher interface {
	PayForFinishedTaskEventPublisher
	TaskPriceEstimatedEventPublisher
	ChargeForAssignedTaskEventPublisher
}
