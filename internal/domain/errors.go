package domain

import "errors"

var (
	// ErrNotFound reutilizable
	ErrNotFound = errors.New("recurso no encontrado")
	// ErrAlreadyExists reutilizable
	ErrAlreadyExists = errors.New("recurso ya existe")
	// ErrUnauthorized reutilizable
	ErrUnauthorized = errors.New("no autorizado")
	// ErrForbidden reutilizable
	ErrForbidden = errors.New("acceso prohibido")
	// ErrInvalidInput reutilizable
	ErrInvalidInput = errors.New("datos de entrada inválidos")
	// ErrInternalServer reutilizable
	ErrInternalServer = errors.New("error interno del servidor")
)
