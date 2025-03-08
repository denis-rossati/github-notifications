package internal

import (
	"errors"
	"flag"
)

type Arguments struct {
	Token string
}

func GetArgs() (Arguments, error) {
	args := Arguments{}

	token := flag.String(
		"token",
		"",
		"The authentication token used to read GitHub events",
	)

	flag.Parse()

	if *token == "" {
		return args, errors.New("authentication token not provided")
	}

	args.Token = *token

	return args, nil
}
