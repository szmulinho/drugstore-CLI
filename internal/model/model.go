package model

import (
	"os"
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

type JwtToken struct {
	Token string `json:"token"`
}

type Exception struct {
	Message string `json:"message"`
}

type Response struct {
	Data string `json:"data"`
}

var JwtKey = []byte(os.Getenv("JWT_KEY"))

var Prescs []CreatePrescInput

var Expiration string

var Drugs []string

type CreatePrescInput struct {
	PreId      string    `json:"pre-id"`
	Drugs      []string  `json:"drugs"`
	Expiration time.Time `json:"expiration"`
}

var PrescriptionID string
