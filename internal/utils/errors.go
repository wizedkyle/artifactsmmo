package utils

import "errors"

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
	ItemNotFound                      = "Item not found."
	ItemMissingOrInsufficientQuantity = "Missing item or insufficient quantity in your inventory."
	MapNotFound                       = "Map not found."
	NotAuthenticated                  = "Not authenticated."
	ResourceNotFound                  = "Resource not found on this map."
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
	ErrItemMissingOrInsufficientQuantity = errors.New(ItemMissingOrInsufficientQuantity)
	ErrMapNotFound                       = errors.New(MapNotFound)
	ErrNotAuthenticated                  = errors.New(NotAuthenticated)
	ErrResourceNotFound                  = errors.New(ResourceNotFound)
)
