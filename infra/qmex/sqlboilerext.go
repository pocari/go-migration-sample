package qmex

import "fmt"

func Desc(column string) string {
	return fmt.Sprintf("%s desc", column)
}
func Asc(column string) string {
	return fmt.Sprintf("%s asc", column)
}
