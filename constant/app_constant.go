package constant

type ResponseStatus int
type Header int
type General int

// Constantes de la api
const (
	Success ResponseStatus = iota + 1 // Constante desde 1 - Success 1 , DataNotFound 2 ...
	DataNotFound
	UknownError
	InvalidRequest
	Unauthorized
)

func (r ResponseStatus) GetResponseStatus() string {
	return [...]string{"SUCCESS", "DATA_NOT_FOUND", "UNKNOWN_ERROR", "INVALID_REQUEST", "UNAUTHORIZED"}[r-1] //Acceso al index como constante 1 se resta
}

func (r ResponseStatus) GetResponseMessage() string {
	return [...]string{"Success", "Data Not Found", "Unknown Error", "Invalid Request", "Unauthorized"}[r-1]
}
