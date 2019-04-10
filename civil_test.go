package pqt

import (
	"testing"

	"cloud.google.com/go/civil"
)

func TestAddMonths(t *testing.T) {
	type Case struct {
		Date           Date
		Add            int
		ExpectedString string
	}
	cases := []Case{
		{
			Date:           Date{Date: civil.Date{Day: 1, Month: 1, Year: 2000}},
			Add:            5,
			ExpectedString: "2000-06-01",
		},
		{
			Date:           Date{Date: civil.Date{Day: 1, Month: 1, Year: 2000}},
			Add:            36,
			ExpectedString: "2003-01-01",
		},
		{
			Date:           Date{Date: civil.Date{Day: 1, Month: 1, Year: 2000}},
			Add:            14,
			ExpectedString: "2001-03-01",
		},
		{
			Date:           Date{Date: civil.Date{Day: 1, Month: 1, Year: 2000}},
			Add:            -1,
			ExpectedString: "1999-12-01",
		},
		{
			Date:           Date{Date: civil.Date{Day: 1, Month: 1, Year: 2000}},
			Add:            -12,
			ExpectedString: "1999-01-01",
		},
		{
			Date:           Date{Date: civil.Date{Day: 1, Month: 1, Year: 2000}},
			Add:            -14,
			ExpectedString: "1998-11-01",
		},
	}

	for _, c := range cases {
		c.Date.AddMonths(c.Add)
		if c.Date.String() != c.ExpectedString {
			t.Errorf("Expected string %s, but had %s", c.ExpectedString, c.Date.String())
		}
	}
}
