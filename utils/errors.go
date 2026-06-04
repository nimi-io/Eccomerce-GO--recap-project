package utils

type ErrorCode string

const (
	ErrEmailTaken    ErrorCode = "email_taken"
	ErrEmailInvalid  ErrorCode = "invalid_email"
	ErrPasswordWeak  ErrorCode = "password_weak"
	ErrUserNotFound  ErrorCode = "user_not_found"
	ErrInvalidCreds  ErrorCode = "invalid_credentials"
	ErrDatabaseError ErrorCode = "database_error"
	ErrValidation    ErrorCode = "validation_error"
)

type AppError struct {
	Code    ErrorCode `json:"code"`
	Message string    `json:"message"`
	Status  int       `json:"-"`
}

func (e *AppError) Error() string {
	return e.Message
}

func NewError(code ErrorCode, message string, status int) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		Status:  status,
	}
}

// Common error constructors
func EmailTakenError() *AppError {
	return NewError(ErrEmailTaken, "Email already registered", 409)
}

func InvalidEmailError() *AppError {
	return NewError(ErrEmailInvalid, "Invalid email format", 400)
}

func PasswordWeakError() *AppError {
	return NewError(ErrPasswordWeak, "Password must be at least 8 characters", 400)
}

func UserNotFoundError() *AppError {
	return NewError(ErrUserNotFound, "User not found", 404)
}

func InvalidCredentialsError() *AppError {
	return NewError(ErrInvalidCreds, "Invalid email or password", 401)
}

func DatabaseError() *AppError {
	return NewError(ErrDatabaseError, "Database operation failed", 500)
}

func ValidationError(message string) *AppError {
	return NewError(ErrValidation, message, 400)
}
