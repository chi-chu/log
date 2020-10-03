package log


type Hook interface {
	Set(*Entry)
}