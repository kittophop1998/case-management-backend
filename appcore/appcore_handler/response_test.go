package appcore_handler

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

type info struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Age       int    `json:"age"`
	Phone     string `json:"phone"`
}

type pagniation struct {
	Limit      int    `json:"limit"`
	Page       int    `json:"page"`
	Sort       string `json:"sort"`
	TotalRows  int64  `json:"total_rows"`
	TotalPages int    `json:"total_pages"`
	Rows       []info `json:"rows"`
}

func TestNewResponseObject(t *testing.T) {
	i := info{
		FirstName: "Test1",
		LastName:  "test1",
		Age:       10,
		Phone:     "0222222222",
	}

	wanted := "{\"data\":{\"firstName\":\"Test1\",\"lastName\":\"test1\",\"age\":10,\"phone\":\"0222222222\"}}"
	a, err := json.Marshal(NewResponseObject(i))
	assert.Nil(t, err)
	assert.Equal(t, string(a), wanted)
}

func TestNewResponseCreated(t *testing.T) {
	wanted := "{\"data\":{\"id\":10}}"
	a, err := json.Marshal(NewResponseCreated(10))
	assert.Nil(t, err)
	assert.Equal(t, string(a), wanted)

	wanted2 := "{\"data\":{\"id\":\"123-ASR3-1234\"}}"
	a2, err2 := json.Marshal(NewResponseCreated("123-ASR3-1234"))
	assert.Nil(t, err2)
	assert.Equal(t, string(a2), wanted2)
}

func TestNewResponseError(t *testing.T) {
	errorText := "cannot parse json to r325"
	messageText := "your input is invalid"
	wanted := "{\"error\":\"cannot parse json to r325\",\"message\":\"your input is invalid\"}"

	a, err := json.Marshal(NewResponseError(errorText, messageText))
	assert.Nil(t, err)
	assert.Equal(t, string(a), wanted)
}

func TestNewResponseObjectWithSensitiveData(t *testing.T) {
	i := info{
		FirstName: "Test1",
		LastName:  "test1",
		Age:       10,
		Phone:     "0222222222",
	}

	unSelectedField := []string{"age"}
	hideValueField := []string{"phone"}
	wanted := "{\"data\":{\"firstName\":\"Test1\",\"lastName\":\"test1\",\"phone\":\"022xxxxxxx\"}}"

	a, err := json.Marshal(NewResponseObjectWithSensitiveData[info](i, unSelectedField, hideValueField))
	assert.Nil(t, err)
	assert.Equal(t, string(a), wanted)

	ii := []info{
		{
			FirstName: "Test1",
			LastName:  "test1",
			Age:       10,
			Phone:     "0222222222",
		},
		{
			FirstName: "Test2",
			LastName:  "test2",
			Age:       5,
			Phone:     "0822222222",
		},
	}

	unSelectedField = []string{"firstName", "lastName", "age"}
	hideValueField = []string{"phone"}
	wanted = "{\"data\":[{\"phone\":\"022xxxxxxx\"},{\"phone\":\"082xxxxxxx\"}]}"

	b, err := json.Marshal(NewResponseObjectWithSensitiveData[[]info](ii, unSelectedField, hideValueField))
	assert.Nil(t, err)
	assert.Equal(t, string(b), wanted)

	t.Run("for pagniation", func(t *testing.T) {
		iii := pagniation{
			Limit:      5,
			Page:       1,
			Sort:       "id desc",
			TotalPages: 1,
			TotalRows:  2,
			Rows: []info{
				{
					FirstName: "Test1",
					LastName:  "test1",
					Age:       10,
					Phone:     "0222222222",
				},
				{
					FirstName: "Test2",
					LastName:  "test2",
					Age:       5,
					Phone:     "0822222222",
				},
			},
		}
		unSelectedField = []string{"firstName", "lastName", "age"}
		hideValueField = []string{"phone"}
		wanted = "{\"data\":{\"limit\":5,\"page\":1,\"rows\":[{\"phone\":\"022xxxxxxx\"},{\"phone\":\"082xxxxxxx\"}],\"sort\":\"id desc\",\"total_pages\":1,\"total_rows\":2}}"
		b, err := json.Marshal(NewResponseObjectWithSensitiveData[pagniation](iii, unSelectedField, hideValueField))
		assert.Nil(t, err)
		assert.Equal(t, wanted, string(b))
	})
}
