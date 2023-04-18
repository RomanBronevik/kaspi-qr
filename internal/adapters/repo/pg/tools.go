package pg

import (
	"strings"
)

func (d *St) tOptionalWhere(conds []string) string {
	if len(conds) > 0 {
		return ` where ` + strings.Join(conds, " and ") + ` `
	}
	return ``
}
