package file

type RotateType string

const (
	ROTATE_MINITE		RotateType		= "*/1 * * * *"
	ROTATE_HOUR			RotateType		= "0 * * * *"
	ROTATE_DAY			RotateType		= "0 0 * * *"
	ROTATE_WEEK			RotateType		= "0 0 * * 0"
	ROTATE_MONTH		RotateType		= "0 0 1 * *"
)
