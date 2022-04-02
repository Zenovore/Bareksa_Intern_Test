package news

import (
	"bareksaIntern/src/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

func SearchOneNews() echo.HandlerFunc {
	return func(ctx echo.Context) (err error) {
		var payload RequestSearchNews

		if err = ctx.Bind(&payload); err != nil {
			return ctx.JSON(http.StatusBadRequest,
				utils.ErrorPayloadResponse())
		}

		if err = ctx.Validate(&payload); err != nil {
			return ctx.JSON(http.StatusBadRequest,
				utils.ErrorValidationResponse(err.Error()))
		}

		data, err := agent.SearchOneNews(ctx.Request().Context(), payload)

		if err != nil {
			return ctx.JSON(http.StatusBadRequest,
				utils.ErrorProcessingDataResponse(err.Error()))
		}

		return ctx.JSON(http.StatusOK, utils.SuccessResponse(data))
	}
}

func SearchManyNews() echo.HandlerFunc {
	return func(ctx echo.Context) (err error) {
		var payload RequestSearchNews

		if err = ctx.Bind(&payload); err != nil {
			return ctx.JSON(http.StatusBadRequest,
				utils.ErrorPayloadResponse())
		}

		if err = ctx.Validate(&payload); err != nil {
			return ctx.JSON(http.StatusBadRequest,
				utils.ErrorValidationResponse(err.Error()))
		}

		data, err := agent.SearchManyNews(ctx.Request().Context(), payload)

		if err != nil {
			return ctx.JSON(http.StatusBadRequest,
				utils.ErrorProcessingDataResponse(err.Error()))
		}

		return ctx.JSON(http.StatusOK, utils.SuccessResponse(data))
	}
}

func CreateNews() echo.HandlerFunc {
	return func(ctx echo.Context) (err error) {
		var payload RequestCreateNews

		if err = ctx.Bind(&payload); err != nil {
			return ctx.JSON(http.StatusBadRequest,
				utils.ErrorPayloadResponse())
		}

		if err = ctx.Validate(&payload); err != nil {
			return ctx.JSON(http.StatusBadRequest,
				utils.ErrorValidationResponse(err.Error()))
		}

		data, err := agent.CreateNews(ctx.Request().Context(), payload)

		if err != nil {
			return ctx.JSON(http.StatusBadRequest,
				utils.ErrorProcessingDataResponse(err.Error()))
		}

		return ctx.JSON(http.StatusOK, utils.SuccessResponse(data))
	}
}

func UpdateNews() echo.HandlerFunc {
	return func(ctx echo.Context) (err error) {
		var payload RequestUpdateNews

		if err = ctx.Bind(&payload); err != nil {
			return ctx.JSON(http.StatusBadRequest,
				utils.ErrorPayloadResponse())
		}

		if err = ctx.Validate(&payload); err != nil {
			return ctx.JSON(http.StatusBadRequest,
				utils.ErrorValidationResponse(err.Error()))
		}

		data, err := agent.UpdateNews(ctx.Request().Context(), payload)

		if err != nil {
			return ctx.JSON(http.StatusBadRequest,
				utils.ErrorProcessingDataResponse(err.Error()))
		}

		return ctx.JSON(http.StatusOK, utils.SuccessResponse(data))
	}
}

func DeleteNews() echo.HandlerFunc {
	return func(ctx echo.Context) (err error) {
		var payload RequestDeleteNews

		if err = ctx.Bind(&payload); err != nil {
			return ctx.JSON(http.StatusBadRequest,
				utils.ErrorPayloadResponse())
		}

		if err = ctx.Validate(&payload); err != nil {
			return ctx.JSON(http.StatusBadRequest,
				utils.ErrorValidationResponse(err.Error()))
		}

		err = agent.DeleteNews(ctx.Request().Context(), payload)

		if err != nil {
			return ctx.JSON(http.StatusBadRequest,
				utils.ErrorProcessingDataResponse(err.Error()))
		}

		return ctx.JSON(http.StatusOK, utils.SuccessResponse(true))
	}
}

func FilterNewsByTopic() echo.HandlerFunc {
	return func(ctx echo.Context) (err error) {
		var payload RequestSearchNews

		if err = ctx.Bind(&payload); err != nil {
			return ctx.JSON(http.StatusBadRequest,
				utils.ErrorPayloadResponse())
		}

		if err = ctx.Validate(&payload); err != nil {
			return ctx.JSON(http.StatusBadRequest,
				utils.ErrorValidationResponse(err.Error()))
		}

		data, err := agent.FilterNewsByTopic(ctx.Request().Context(), payload)

		if err != nil {
			return ctx.JSON(http.StatusBadRequest,
				utils.ErrorProcessingDataResponse(err.Error()))
		}

		return ctx.JSON(http.StatusOK, utils.SuccessResponse(data))
	}
}
