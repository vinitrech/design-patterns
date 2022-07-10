package hexagonal_architecture

type RedirectService interface {
	Find(code string) (*Redirect, error)
	Store(redirect *Redirect) error
}
