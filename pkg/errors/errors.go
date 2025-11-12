package errors

import (
	"fmt"
	"net/http"
)

type ErrorType string

const (
    ErrorTypeValidation   ErrorType = "VALIDATION_ERROR"
    ErrorTypeNotFound     ErrorType = "NOT_FOUND"
    ErrorTypeConflict     ErrorType = "CONFLICT"

    ErrorTypeInvalidToken       ErrorType = "INVALID_TOKEN"
    ErrorTypeSessionNotFound    ErrorType = "SESSION_NOT_FOUND"
    ErrorTypeForbidden          ErrorType = "FORBIDDEN"
    
    ErrorTypeSuspiciousActivity ErrorType = "SUSPICIOUS_ACTIVITY"
    ErrorTypeRateLimitExceeded  ErrorType = "RATE_LIMIT_EXCEEDED"

    ErrorTypeInternal     ErrorType = "INTERNAL_ERROR"
)

// Кастомная ошибка приложения
type AppError struct {
    Type    ErrorType `json:"type"`
    Message string    `json:"message"`
    Details any       `json:"details,omitempty"`
    Err     error     `json:"-"`
}

func MapAppErrorToHTTPStatus(appErr *AppError) int {
    switch appErr.Type {
    case ErrorTypeValidation:
        return http.StatusBadRequest
    case ErrorTypeNotFound:
        return http.StatusNotFound
    case ErrorTypeConflict:
        return http.StatusConflict
    case ErrorTypeInvalidToken, ErrorTypeSessionNotFound:
        return http.StatusUnauthorized
    case ErrorTypeForbidden:
        return http.StatusForbidden
    case ErrorTypeSuspiciousActivity:
        return http.StatusForbidden 
    case ErrorTypeRateLimitExceeded:
        return http.StatusTooManyRequests
    case ErrorTypeInternal:
        return http.StatusInternalServerError
    default:
        return http.StatusInternalServerError
    }
}

func Is(err error, targetType ErrorType) bool {
    appErr, ok := err.(*AppError)
    return ok && appErr.Type == targetType
}

func GetType(err error) (ErrorType, bool) {
    appErr, ok := err.(*AppError)
    if !ok {
        return "", false
    }
    return appErr.Type, true
}


func (e *AppError) Error() string {
    if e.Err != nil {
        return fmt.Sprintf("%s: %s (%v)", e.Type, e.Message, e.Err)
    }
    return fmt.Sprintf("%s: %s", e.Type, e.Message)
}

func (e *AppError) Unwrap() error {
    return e.Err
}

//ошибки токена и сессии

func NewInvalidCredentialsError() *AppError {
    return &AppError{
        Type:    ErrorTypeForbidden,
        Message: "Invalid email or password",
    }
}

func NewInvalidTokenError() *AppError {
    return &AppError{
        Type:    ErrorTypeInvalidToken,
        Message: "Invalid token",
        Details: map[string]interface{}{
            "action_required": "reauthenticate",
        },
    }
}

func NewSessionNotFoundError() *AppError {
    return &AppError{
        Type:    ErrorTypeSessionNotFound,
        Message: "Session not found or expired",
        Details: map[string]interface{}{
            "action_required": "reauthenticate",
        },
    }
}

func NewSuspiciousActivityError(reason string) *AppError {
    return &AppError{
        Type:    ErrorTypeSuspiciousActivity,
        Message: "Suspicious activity detected",
        Details: map[string]interface{}{
            "action_required": "confirm_identity",
            "reason": reason,
        },
    }
}

func NewFordidenErrror(msg string) *AppError {
    return &AppError{
        Type:    ErrorTypeForbidden,
        Message: fmt.Sprintf("access denied: %s", msg),
    }
}


//валидация invalide entity

func NewValidationError(message string, details any) *AppError {
    return &AppError{
        Type:    ErrorTypeValidation,
        Message: message,
        Details: details,
    }
}

// БД ошибки 
func NewNotFoundError(resource, id string) *AppError {
    return &AppError{
        Type:    ErrorTypeNotFound,
        Message: fmt.Sprintf("%s with id %s not found", resource, id),
    }
}


func NewConflictError(resource, field, value string) *AppError {
    return &AppError{
        Type:    ErrorTypeConflict,
        Message: fmt.Sprintf("%s with %s '%s' already exists", resource, field, value),
    }
}


//ограничение запросов
func NewRateLimitError(retryAfter int) *AppError {
    return &AppError{
        Type:    ErrorTypeRateLimitExceeded,
        Message: "Too many requests",
        Details: map[string]interface{}{
            "retry_after_seconds": retryAfter,
        },
    }
}

//наша, внутренняя ошибка - 500
func NewInternalError(err error) *AppError {
    return &AppError{
        Type:    ErrorTypeInternal,
        Message: "Internal server error",
        Err:     err,
    }
}
