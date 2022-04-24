package exception

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/webtoor/go-fiber/model/web"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {

	_, validator := err.(validator.ValidationErrors)

	if validator {
		return ctx.Status(fiber.StatusBadRequest).JSON(web.JsonResponse{
			Code:   400,
			Status: "BAD REQUEST",
			Data:   validationResponse(err),
		})
	}

	if err.Error() == "record not found" {
		return ctx.Status(fiber.StatusNotFound).JSON(web.JsonResponse{
			Code:   404,
			Status: "RECORD NOT FOUND",
			Data:   err.Error(),
		})
	}

	return ctx.Status(fiber.StatusInternalServerError).JSON(web.JsonResponse{
		Code:   500,
		Status: "INTERNAL SERVER ERROR",
		Data:   err.Error(),
	})
}

func validationResponse(err error) []*web.ValidationMessage {
	var errors []*web.ValidationMessage
	for _, err := range err.(validator.ValidationErrors) {
		var element web.ValidationMessage
		element.FailedField = err.StructNamespace()
		element.Tag = err.Tag()
		element.Value = err.Param()
		errors = append(errors, &element)
	}
	return errors
}
