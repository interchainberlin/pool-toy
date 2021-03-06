package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrInvalid    = sdkerrors.Register(ModuleName, 1, "custom error message")
	ErrUserExists = sdkerrors.Register(ModuleName, 2, "user exists")
)
