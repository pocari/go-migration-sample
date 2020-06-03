package main

import (
	"context"
	"fmt"
	"migration-sample/infra/db"
	"migration-sample/models"
	"os"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func init() {
	c, err := db.NewDB()
	if err != nil {
		panic(err)
	}
	boil.SetDB(c)
	boil.DebugMode = true
	boil.DebugWriter = os.Stdout
}

func OrderByDesc(column string) string {
	return fmt.Sprintf("%s desc", column)
}

func main() {
	// select * from foos where field1 in ('field1-2', field1-3') order by field1 desc
	fooList, err := models.Foos(
		models.FooWhere.Field1.IN([]string{
			"field1-2",
			"field1-3",
		}),
		qm.OrderBy(OrderByDesc(models.FooColumns.Field1)),
	).All(context.Background(), boil.GetContextDB())

	if err != nil {
		panic(err)
	}

	for i, foo := range fooList {
		fmt.Printf("%02d: %+v\n", i, foo)
	}
}
