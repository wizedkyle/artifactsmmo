package utils

import "errors"

const (
	CharacterAtDestination          = "Character already at destination."
	CharacterCooldown               = "Character in cooldown."
	CharacterInventoryFull          = "Character inventory full."
	CharacterLockedActionInProgress = "Character is locked. Action is already in progress."
	CharacterNotFound               = "Character not found."
	CharacterNotAtSkillLevel        = "Character not at required skill level."
	GenericError                    = "Unknown error, server could be unavailable."
	MapNotFound                     = "Map not found."
	NotAuthenticated                = "Not authenticated."
	ResourceNotFound                = "Resource not found on this map."
)

var (
	ErrCharacterAtDestination          = errors.New(CharacterAtDestination)
	ErrCharacterCooldown               = errors.New(CharacterCooldown)
	ErrCharacterInventoryFull          = errors.New(CharacterInventoryFull)
	ErrCharacterLockedActionInProgress = errors.New(CharacterLockedActionInProgress)
	ErrCharacterNotFound               = errors.New(CharacterNotFound)
	ErrCharacterNotAtSkillLevel        = errors.New(CharacterNotAtSkillLevel)
	ErrGenericError                    = errors.New(GenericError)
	ErrMapNotFound                     = errors.New(MapNotFound)
	ErrNotAuthenticated                = errors.New(NotAuthenticated)
	ErrResourceNotFound                = errors.New(ResourceNotFound)
)
