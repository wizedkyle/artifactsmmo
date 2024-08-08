package utils

import "errors"

type Error struct {
	InternalError error         `json:"internalError"`
	ExternalError ExternalError `json:"externalError"`
}

type ExternalError struct {
	Id            string `json:"id"`
	Message       string `json:"message"`
	Code          int    `json:"code"`
	TransactionId string `json:"transactionId"`
}

const (
	CharacterAtDestination            = "Character already at destination."
	CharacterCooldown                 = "Character in cooldown."
	CharacterInventoryFull            = "Character inventory full."
	CharacterLockedActionInProgress   = "Character is locked. Action is already in progress."
	CharacterLimitReached             = "Maximum characters reached on your account."
	CharacterNameInUse                = "Name already used."
	CharacterNotFound                 = "Character not found."
	CharacterNotAtSkillLevel          = "Character not at required skill level."
	GenericError                      = "Unknown error, server could be unavailable."
	GenericInternalServerErrorMessage = "internal server error"
	InvalidRequestBody                = "invalid request body"
	ItemNotFound                      = "Item not found."
	ItemsNotFound                     = "Items not found."
	ItemMissingOrInsufficientQuantity = "Missing item or insufficient quantity in your inventory."
	MapNotFound                       = "Map not found."
	NotAuthenticated                  = "Not authenticated."
	ResourceNotFound                  = "Resource not found on this map."
	TaskNotFound                      = "Task not found."
	TasksNotFound                     = "Tasks not found."
)

var (
	ErrCharacterAtDestination            = errors.New(CharacterAtDestination)
	ErrCharacterCooldown                 = errors.New(CharacterCooldown)
	ErrCharacterInventoryFull            = errors.New(CharacterInventoryFull)
	ErrCharacterLockedActionInProgress   = errors.New(CharacterLockedActionInProgress)
	ErrCharacterLimitReached             = errors.New(CharacterLimitReached)
	ErrCharacterNameInUse                = errors.New(CharacterNameInUse)
	ErrCharacterNotFound                 = errors.New(CharacterNotFound)
	ErrCharacterNotAtSkillLevel          = errors.New(CharacterNotAtSkillLevel)
	ErrGenericError                      = errors.New(GenericError)
	ErrItemNotFound                      = errors.New(ItemNotFound)
	ErrItemsNotFound                     = errors.New(ItemsNotFound)
	ErrItemMissingOrInsufficientQuantity = errors.New(ItemMissingOrInsufficientQuantity)
	ErrMapNotFound                       = errors.New(MapNotFound)
	ErrNotAuthenticated                  = errors.New(NotAuthenticated)
	ErrResourceNotFound                  = errors.New(ResourceNotFound)
	ErrTaskNotFound                      = errors.New(TaskNotFound)
	ErrTasksNotFound                     = errors.New(TasksNotFound)
)

// GenerateError
// Generates an error message that is used when processing API requests.
func GenerateError(id string, message string, code int, transactionId string, err error) *Error {
	return &Error{
		InternalError: err,
		ExternalError: ExternalError{
			Id:            id,
			Message:       message,
			Code:          code,
			TransactionId: transactionId,
		},
	}
}
