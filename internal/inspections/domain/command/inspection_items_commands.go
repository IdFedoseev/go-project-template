package command

import (
	"context"
	"time"
)

type CreateInspectionItemCommand struct {
	Ctx       context.Context
	Question  string
	Answer    string
	PhotoUrls []string
	Score     int
	Comment   string
}

func (c *CreateInspectionItemCommand) Context() context.Context { return c.Ctx }

type UpdateInspectionItemCommand struct {
	Ctx       context.Context
	ID        string
	Question  string
	Answer    string
	PhotoUrls []string
	Score     int
	Comment   string
	UpdatedAt time.Time
}

func (c *UpdateInspectionItemCommand) Context() context.Context { return c.Ctx }

type DeleteUserCommand struct {
	ID string `validate:"required,hexadecimal"`
}
