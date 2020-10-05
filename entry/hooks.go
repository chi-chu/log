package entry

type Hook interface {
	Set(*Entry)
}