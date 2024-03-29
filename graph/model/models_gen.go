// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type BingoCard struct {
	Numbers []int `json:"numbers"`
}

type LoginInput struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

type User struct {
	ID   string `json:"id" bson:"_id"`
	Name string `json:"name"`
	Role string `json:"role"`
}

type ValidateResult struct {
	Row      []*int `json:"row"`
	Col      []*int `json:"col"`
	Diagonal []*int `json:"diagonal"`
	Numbers  []int  `json:"numbers"`
	IsValid  bool   `json:"isValid"`
}
