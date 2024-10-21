package sqlbuilder

import (
	"fmt"
	"testing"

	"github.com/huandu/go-sqlbuilder"
)

func TestName(t *testing.T) {
	sb := sqlbuilder.NewSelectBuilder()
	sb.Where(
		sb.Or(
			sb.E("name", "foo"),
			sb.E("code", "foo@%"),
		),
		sb.Or(
			sb.E("start_site_id", "foo"),
			sb.E("end_site_id", "foo@%"),
		),
	)
	sql, args := sb.Build()
	fmt.Println(sql)
	fmt.Println(sql)
	fmt.Println(args)
}
