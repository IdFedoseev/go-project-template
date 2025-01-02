package command

import (
    "context"
    "time"
)

type CreateUserCommand struct {
    Ctx      context.Context
    Username string
    Email    string
    Password string
}

func (c CreateUserCommand) Context() context.Context {
    return c.Ctx
}

type UpdateUserCommand struct {
    Ctx      context.Context
    ID       string
    Username string
    Email    string
    Password string
    UpdatedAt time.Time
}

func (c UpdateUserCommand) Context() context.Context {
    return c.Ctx
}

type DeleteUserCommand struct {
    ID string `validate:"required,hexadecimal"`
} 