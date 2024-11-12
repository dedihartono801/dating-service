package helper

import (
	"context"
	"database/sql"
	"dating-service/pkg/constant"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func HandleError(err error) error {

	switch {
	case err == context.DeadlineExceeded:
		return Error(http.StatusRequestTimeout, constant.MsgErrorTimeout, err)
	case err == sql.ErrNoRows:
		return Error(http.StatusNotFound, constant.MsgErrorNotFound, err)
	default:
		return Error(http.StatusInternalServerError, constant.MsgErrorInternal, err)
	}
}

func Error(statusCode int, msg string, err error) error {
	//logger.LogError(constant.Err, strconv.Itoa(statusCode), fmt.Sprintf(msg+err.Error()))
	return fmt.Errorf("%d | %s | %w", statusCode, msg, err)
}

func TrimMesssage(err error) (statusCode int, customError, originalError string) {
	errs := strings.Split(err.Error(), "|")
	statusCode, _ = strconv.Atoi(strings.TrimSpace(errs[0]))
	customError = strings.TrimSpace(errs[1])
	originalError = strings.TrimSpace(errs[2])
	return
}
