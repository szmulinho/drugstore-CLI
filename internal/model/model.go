package model

import (
	"time"
)

type Drug struct {
	DrugID string `json:"drug-id"`
	Name   string `json:"name"`
	Type   string `json:"type"`
	Price  string `json:"price"`
}

var Drugs []Drug

var DrugId string

var Name string

var Type string

var Price string

var D Drug

var P CreatePrescInput

type User struct {
	UserID   string `json:"user-id"`
	Password string `json:"password"`
}

var Prescs []CreatePrescInput

type CreatePrescInput struct {
	PreId      string    `json:"pre-id"`
	Drugs      []string  `json:"drugs"`
	Expiration time.Time `json:"expiration"`
}
