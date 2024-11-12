package helper

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"

	"dating-service/internal/entity"
	"dating-service/pkg/constant"
)

func CreateContextWithTimeout() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 5*time.Second)
}

func CreateContextWithCustomTimeout(timeout int) (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
}

func SetValueToContext(ctx context.Context, c *fiber.Ctx) context.Context {
	var userId, email, gender string

	userId, ok := c.Locals("user-id").(string)
	if !ok {
		userId = "0"
	}
	email, ok = c.Locals("email").(string)
	if !ok {
		email = "0"
	}
	gender, ok = c.Locals("gender").(string)
	if !ok {
		userId = "0"
	}

	isVerified, ok := c.Locals("is-verified").(bool)
	if !ok {
		isVerified = false
	}

	isPremium, ok := c.Locals("is-premium").(bool)
	if !ok {
		isPremium = false
	}

	ctx = context.WithValue(ctx, constant.HeaderContext, entity.ValueContext{
		UserId:     userId,
		Email:      email,
		Gender:     gender,
		IsVerified: isVerified,
		IsPremium:  isPremium,
	})

	return context.WithValue(ctx, constant.FiberContext, c)
}

func GetValueContext(ctx context.Context) (valueCtx entity.ValueContext) {
	return ctx.Value(constant.HeaderContext).(entity.ValueContext)
}
