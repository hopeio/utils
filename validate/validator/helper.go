/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package validator

// Validator is a general interface that allows a message to be validated.
type IValidator interface {
	Validate() error
}

func CallValidatorIfExists(candidate interface{}) error {
	if validator, ok := candidate.(IValidator); ok {
		return validator.Validate()
	}
	return nil
}
