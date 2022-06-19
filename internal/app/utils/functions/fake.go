package functions

import "github.com/jaswdr/faker"

type fakerNS struct {
}

func (_ fakerNS) Name() string {
	return faker.New().Person().Name()
}

func (_ fakerNS) FirstName() string { return faker.New().Person().FirstName() }
func (_ fakerNS) LastName() string  { return faker.New().Person().LastName() }
func (_ fakerNS) UUID() string      { return faker.New().UUID().V4() }
