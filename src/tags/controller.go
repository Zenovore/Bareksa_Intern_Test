package tags

import (
	"bareksaIntern/src/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

func SearchOneTag() echo.HandlerFunc {
	return func(ctx echo.Context) (err error) {
		var payload RequestSearchTag

		if err = ctx.Bind(&payload); err != nil {
			return ctx.JSON(http.StatusBadRequest,
				utils.ErrorPayloadResponse())
		}

		if err = ctx.Validate(&payload); err != nil {
			return ctx.JSON(http.StatusBadRequest,
				utils.ErrorValidationResponse(err.Error()))
		}

		data, err := agent.SearchOneTag(ctx.Request().Context(), payload)

		if err != nil {
			return ctx.JSON(http.StatusBadRequest,
				utils.ErrorProcessingDataResponse(err.Error()))
		}

		return ctx.JSON(http.StatusOK, utils.SuccessResponse(data))
	}
}

func CreateTag() echo.HandlerFunc {
	return func(ctx echo.Context) (err error) {
		var payload RequestCreateTag

		if err = ctx.Bind(&payload); err != nil {
			return ctx.JSON(http.StatusBadRequest,
				utils.ErrorPayloadResponse())
		}

		if err = ctx.Validate(&payload); err != nil {
			return ctx.JSON(http.StatusBadRequest,
				utils.ErrorValidationResponse(err.Error()))
		}

		data, err := agent.CreateTag(ctx.Request().Context(), payload)

		if err != nil {
			return ctx.JSON(http.StatusBadRequest,
				utils.ErrorProcessingDataResponse(err.Error()))
		}

		return ctx.JSON(http.StatusOK, utils.SuccessResponse(data))
	}
}

func UpdateTag() echo.HandlerFunc {
	return func(ctx echo.Context) (err error) {
		var payload RequestUpdateTag

		if err = ctx.Bind(&payload); err != nil {
			return ctx.JSON(http.StatusBadRequest,
				utils.ErrorPayloadResponse())
		}

		if err = ctx.Validate(&payload); err != nil {
			return ctx.JSON(http.StatusBadRequest,
				utils.ErrorValidationResponse(err.Error()))
		}

		data, err := agent.UpdateTag(ctx.Request().Context(), payload)

		if err != nil {
			return ctx.JSON(http.StatusBadRequest,
				utils.ErrorProcessingDataResponse(err.Error()))
		}

		return ctx.JSON(http.StatusOK, utils.SuccessResponse(data))
	}
}

func DeleteTag() echo.HandlerFunc {
	return func(ctx echo.Context) (err error) {
		var payload RequestDeleteTag

		if err = ctx.Bind(&payload); err != nil {
			return ctx.JSON(http.StatusBadRequest,
				utils.ErrorPayloadResponse())
		}

		if err = ctx.Validate(&payload); err != nil {
			return ctx.JSON(http.StatusBadRequest,
				utils.ErrorValidationResponse(err.Error()))
		}

		err = agent.DeleteTag(ctx.Request().Context(), payload)

		if err != nil {
			return ctx.JSON(http.StatusBadRequest,
				utils.ErrorProcessingDataResponse(err.Error()))
		}

		return ctx.JSON(http.StatusOK, utils.SuccessResponse(true))
	}
}
