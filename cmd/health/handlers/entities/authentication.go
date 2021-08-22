package entities

type HeaderAuth struct {
	Auth string `header:"Auth" binding:"required"`
}
