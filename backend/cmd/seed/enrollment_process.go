package main

import (
	"github.com/enrollment/src/db/ports"
	"golang.org/x/net/context"
)

// func createRandomAccount(faker faker.Faker) db.CreateAccountWithProviderNameParams {
// 	return db.CreateAccountWithProviderNameParams{
// 		Name:         faker.Person().Name(),
// 		Email:        faker.Internet().Email(),
// 		Surname:      faker.Person().LastName(),
// 		AvatarUrl:    AVATAR_URL,
// 		ProviderName: GOOGLE_PROVIDER,
// 		AccessToken:  faker.RandomStringWithLength(20),
// 		RefreshToken: faker.RandomStringWithLength(20),
// 	}
// }
//

// func createEnrollmentProcess(major db.Major) db.Create

func seedEnrollmentProcessTables(
	ctx context.Context,
	oauthRepo ports.OauthRepositoryInterface,
) {
	// faker := faker.New()

}
