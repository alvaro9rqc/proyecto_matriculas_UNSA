package main

import (
	"log"

	"github.com/enrollment/src/db/ports"
	"github.com/jaswdr/faker"
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

func seedEnrollmentRepository(ctx context.Context, oauthRepo ports.OauthRepositoryInterface) {
	faker := faker.New()

	log.Println("Seeding OAuth providers...")

	err := oauthRepo.CreateOauthProvider(ctx, GOOGLE_PROVIDER)
	if err != nil {
		log.Fatal(err)
	}
	err = oauthRepo.CreateOauthProvider(ctx, MICROSOFT_PROVIDER)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Seeding OAuth accounts...")
	for range 100 {
		account := createRandomAccount(faker)
		err = oauthRepo.CreateAccountWithProviderName(ctx, account)
		if err != nil {
			log.Fatal(err)
		}
	}

	log.Println("Seeding completed successfully.")
}
