package command

import (
	domainCommand "proj/internal/user/domain/command"
	"proj/internal/user/domain/entity"
	"proj/internal/user/domain/repository"
	"proj/pkg/logger"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"proj/pkg/tracer"
)

type UserCommandHandler struct {
	userRepo repository.UserRepository
}

func NewUserCommandHandler(repo repository.UserRepository) *UserCommandHandler {
	return &UserCommandHandler{
		userRepo: repo,
	}
}

func (h *UserCommandHandler) HandleCreate(cmd domainCommand.CreateUserCommand) (*entity.User, error) {
	ctx, span := tracer.StartSpan(cmd.Context(), "create_user_command")
	defer span.End()

	start := time.Now()
	logger.Info("Executing create user command",
		zap.String("username", cmd.Username),
		zap.String("email", cmd.Email),
	)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(cmd.Password), bcrypt.DefaultCost)
	if err != nil {
		logger.Error("Failed to hash password", zap.Error(err))
		return nil, err
	}

	user := &entity.User{
		Username:  cmd.Username,
		Email:     cmd.Email,
		Password:  string(hashedPassword),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := h.userRepo.Create(ctx, user); err != nil {
		logger.Error("Failed to create user",
			zap.Error(err),
			zap.String("email", user.Email),
		)
		return nil, err
	}

	logger.Info("User created successfully",
		zap.String("email", user.Email),
		zap.Duration("duration", time.Since(start)),
	)
	return user, nil
}

func (h *UserCommandHandler) HandleUpdate(cmd domainCommand.UpdateUserCommand) (*entity.User, error) {
	ctx, span := tracer.StartSpan(cmd.Context(), "update_user_command")
	defer span.End()
	objectID, err := primitive.ObjectIDFromHex(cmd.ID)
	if err != nil {
		return nil, err
	}

	fetchedUser, err := h.userRepo.GetByID(objectID)
	if err != nil {
		return nil, err
	}

	if cmd.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(cmd.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}
		cmd.Password = string(hashedPassword)
	}

	user := &entity.User{
		ID:        fetchedUser.ID,
		CreatedAt: fetchedUser.CreatedAt,
		Username:  cmd.Username,
		Email:     cmd.Email,
		Password:  cmd.Password,
		UpdatedAt: time.Now(),
	}

	if err := h.userRepo.Update(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

func (h *UserCommandHandler) HandleDelete(cmd domainCommand.DeleteUserCommand) error {
	objectID, err := primitive.ObjectIDFromHex(cmd.ID)
	if err != nil {
		return err
	}
	return h.userRepo.Delete(objectID)
}
