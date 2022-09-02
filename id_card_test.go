package validatorx

import (
	"testing"
)

func TestIdCard(t *testing.T) {
	type Test struct {
		IdCard string `validate:"id_card"`
	}

	type Test15 struct {
		IdCard string `validate:"id_card_15"`
	}

	t.Run(`valid_18_len_id_cards`, func(t *testing.T) {
		idCards := []string{
			"110101200503078111",
			"32010220100307654X",
			"510105198205074024",
			"510105198205076361",
			"51010519820507928X",
			"32010220100307654X",
			"510105198205070680",
			"510105198205075246",
			"350102200104077999",
			"350102200104079433",
			"220102201003076431",
			"510105198205076169",
			"350781196403070104",
			"350781196403075044",
			"350781196403071780",
		}
		for _, idCard := range idCards {
			test := Test{}
			test.IdCard = idCard
			err := New().Struct(test)
			if err != nil {
				t.Error(err)
			}
		}
	},
	)

	t.Run(`valid_15len`, func(t *testing.T) {
		test := Test15{}
		test.IdCard = `510802950901131`
		err := New().Struct(test)
		if err != nil {
			t.Error(err)
		}
	},
	)

	t.Run(`invalid_18len_brith_day_part`, func(t *testing.T) {
		test1 := Test{}
		test1.IdCard = `510802177009014131`
		err := New().Struct(test1)
		if err == nil {
			t.Error(`510802177009014131 should be invalid`)
		}

		test2 := Test{}
		test2.IdCard = `510802199513014131`
		err = New().Struct(test2)
		if err == nil {
			t.Error(`510802199513014131 should be invalid`)
		}

		test3 := Test{}
		test3.IdCard = `510802199502314139`
		err = New().Struct(test3)
		if err == nil {
			t.Error(`510802199502314139 should be invalid`)
		}
	},
	)

	t.Run(`invalid_18len_last_index `, func(t *testing.T) {
		test := Test{}
		test.IdCard = `51080219940921413x`
		err := New().Struct(test)
		if err == nil {
			t.Error(`51080219940921413x should be invalid`)
		}
	},
	)
}
