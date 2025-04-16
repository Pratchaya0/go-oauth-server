package middlewarerepository

type (
	IMiddlewareRepository interface{}

	middlewareRepository struct{}
)

func NewMiddlewareRepository() IMiddlewareRepository {
	return &middlewareRepository{}
}
