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

func (d *St) tPrepareFieldsToCreate(fields map[string]any) (string, string) {
	var keys, values string
	for k := range fields {
		if keys != `` {
			keys += `,`
			values += `,`
		}
		keys += k
		values += `${` + k + `}`
	}
	return keys, values
}

func (d *St) tPrepareFieldsToUpdate(fields map[string]any) string {
	var result string
	for k := range fields {
		if result != `` {
			result += `,`
		}
		result += k + `=${` + k + `}`
	}
	return result
}
