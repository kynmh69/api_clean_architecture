package interactor

import (
	"ca/v2/usecase/port"
	"context"
)

type User struct {
	OutputPort port.UserOutputPort
	UserRepo   port.UserRepository
}

func NewUserInputPort(outputPort port.UserOutputPort, userRepository port.UserRepository) port.UserInputPort {
	return &User{
		OutputPort: outputPort,
		UserRepo:   userRepository,
	}
}

func (u *User) GetUserByID(ctx context.Context, userID string) {
	user, err := u.UserRepo.GetUserByID(ctx, userID)
	if err != nil {
		u.OutputPort.RenderError(err)
		return
	}
	u.OutputPort.Render(user)
}
