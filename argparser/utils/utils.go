package utils

import "fmt"

func Coalesce(vals []any) (any, error) {

	//Checks only for nil values, so values such as empty strings or 0's will return

	for _, v := range vals {

		if v != nil {
			return v, nil
		}

	}

	return nil, fmt.Errorf("no non-null values were found")

}
