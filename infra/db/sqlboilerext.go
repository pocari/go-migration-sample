package db

import "fmt"

func OrderByDesc(column string) string {
	return fmt.Sprintf("%s desc", column)
}
