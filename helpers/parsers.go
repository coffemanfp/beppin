package helpers

import (
	"database/sql"
	"fmt"

	dbm "github.com/coffemanfp/beppin-server/database/models"
	"github.com/coffemanfp/beppin-server/errors"
	"github.com/coffemanfp/beppin-server/models"
)

// ParseDBModelToModel - Parse any valid database model to a normal model.
func ParseDBModelToModel(dbModel interface{}) (model interface{}, err error) {
	switch dbModel.(type) {
	// Users
	case dbm.User:
		model = parseDBUserToUser(dbModel.(dbm.User))
	case dbm.Users:
		model = parseDBUsersToUsers(dbModel.(dbm.Users))

	// Products
	case dbm.Product:
		model = parseDBProductToProduct(dbModel.(dbm.Product))

	case dbm.Products:
		model = parseDBProductsToProducts(dbModel.(dbm.Products))

	// Offers
	case dbm.Offer:
		model = parseDBOfferToOffer(dbModel.(dbm.Offer))
	case dbm.Offers:
		model = parseDBOffersToOffers(dbModel.(dbm.Offers))

	default:
		err = fmt.Errorf("failed to parse database model (%T) to normal model: %w", model, errors.ErrNotSupportedType)
	}

	return
}

// ParseModelToDBModel - Parse any valid normal model to a database model.
func ParseModelToDBModel(model interface{}) (dbModel interface{}, err error) {
	switch model.(type) {
	case models.User:
		dbModel = parseUserToDBUser(model.(models.User))
	case models.Product:
		dbModel = parseProductToDBProduct(model.(models.Product))
	case models.Offer:
		dbModel = parseOfferToDBOffer(model.(models.Offer))
	default:
		err = fmt.Errorf("failed to parse normal model (%T) to database model: %w", model, errors.ErrNotSupportedType)
	}

	return
}

// User parsers

func parseDBUserToUser(dbUser dbm.User) (user models.User) {
	user = models.User{
		ID:       dbUser.ID,
		Language: dbUser.Language.Code,
		Username: dbUser.Username,
		Email:    dbUser.Email,
		Password: dbUser.Password,
		Name:     dbUser.Name,
		LastName: dbUser.LastName,
		Theme:    dbUser.Theme,
	}

	if dbUser.CreatedAt != nil {
		if &dbUser.CreatedAt.Time != nil {
			user.CreatedAt = &dbUser.CreatedAt.Time
		}
	}

	if dbUser.UpdatedAt != nil {
		if &dbUser.UpdatedAt.Time != nil {
			user.UpdatedAt = &dbUser.UpdatedAt.Time
		}
	}

	if dbUser.Birthday != nil {
		if &dbUser.Birthday.Time != nil {
			user.Birthday = &dbUser.Birthday.Time
		}
	}

	return
}

func parseDBUsersToUsers(dbUsers dbm.Users) (users models.Users) {
	var user models.User

	for _, dbUser := range dbUsers {
		user = parseDBUserToUser(dbUser)
		users = append(users, user)
	}
	return
}

func parseUserToDBUser(user models.User) (dbUser dbm.User) {
	dbUser = dbm.User{
		ID: user.ID,
		Language: dbm.Language{
			Code: user.Language,
		},
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
		Name:     user.Name,
		LastName: user.LastName,
		Theme:    user.Theme,
	}

	if user.CreatedAt != nil {
		if &dbUser.CreatedAt.Time != nil {
			dbUser.CreatedAt = new(sql.NullTime)
			dbUser.CreatedAt.Time = *user.CreatedAt
		}
	}

	if user.UpdatedAt != nil {
		if &dbUser.UpdatedAt != nil {
			dbUser.UpdatedAt = new(sql.NullTime)
			dbUser.UpdatedAt.Time = *user.UpdatedAt
		}
	}

	if user.Birthday != nil {
		dbUser.Birthday = new(sql.NullTime)
		dbUser.Birthday.Time = *user.Birthday
	}

	return
}

func parseUsersToDBUsers(users models.Users) (dbUsers dbm.Users) {
	var dbUser dbm.User

	for _, user := range users {
		dbUser = parseUserToDBUser(user)
		dbUsers = append(dbUsers, dbUser)
	}
	return
}

// Product parsers

func parseDBProductToProduct(dbProduct dbm.Product) (product models.Product) {
	product = models.Product{
		ID:          dbProduct.ID,
		UserID:      dbProduct.UserID,
		Name:        dbProduct.Name,
		Description: dbProduct.Description,
		Categories:  dbProduct.Categories,
	}

	if dbProduct.CreatedAt != nil {
		if &dbProduct.CreatedAt.Time != nil {
			product.CreatedAt = &dbProduct.CreatedAt.Time
		}
	}

	if dbProduct.UpdatedAt != nil {
		if &dbProduct.UpdatedAt.Time != nil {
			product.UpdatedAt = &dbProduct.UpdatedAt.Time
		}
	}

	var offer models.Offer

	if dbProduct.Offer != nil {
		offer = parseDBOfferToOffer(*dbProduct.Offer)
	}

	product.Offer = &offer
	return
}

func parseDBProductsToProducts(dbProducts dbm.Products) (products models.Products) {
	var product models.Product

	for _, dbProduct := range dbProducts {
		product = parseDBProductToProduct(dbProduct)
		products = append(products, product)
	}
	return
}

func parseProductToDBProduct(product models.Product) (dbProduct dbm.Product) {
	dbProduct = dbm.Product{
		ID:          product.ID,
		UserID:      product.UserID,
		Name:        product.Name,
		Description: product.Description,
		Categories:  product.Categories,
	}

	var dbOffer dbm.Offer

	if product.Offer != nil {
		dbOffer = parseOfferToDBOffer(*product.Offer)
	}

	dbProduct.Offer = &dbOffer

	if product.CreatedAt != nil {
		if &product.CreatedAt != nil {
			dbProduct.CreatedAt.Time = *product.CreatedAt
		}
	}

	if product.UpdatedAt != nil {
		if &product.UpdatedAt != nil {
			dbProduct.UpdatedAt.Time = *product.UpdatedAt
		}
	}

	return
}

func parseProductsToDBProducts(products models.Products) (dbProducts dbm.Products) {
	var dbProduct dbm.Product

	for _, product := range products {
		dbProduct = parseProductToDBProduct(product)
		dbProducts = append(dbProducts, dbProduct)
	}
	return
}

// Offer parsers

func parseDBOfferToOffer(dbOffer dbm.Offer) (offer models.Offer) {
	offer = models.Offer{
		ID:          dbOffer.ID,
		ProductID:   dbOffer.ProductID,
		Type:        dbOffer.Type,
		Value:       dbOffer.Value,
		ExpiratedAt: &dbOffer.ExpiratedAt.Time,
		CreatedAt:   &dbOffer.CreatedAt.Time,
		UpdatedAt:   &dbOffer.UpdatedAt.Time,
	}
	return
}

func parseDBOffersToOffers(dbOffers dbm.Offers) (offers models.Offers) {
	var offer models.Offer

	for _, dbOffer := range dbOffers {
		offer = parseDBOfferToOffer(dbOffer)
		offers = append(offers, offer)
	}
	return
}

func parseOfferToDBOffer(offer models.Offer) (dbOffer dbm.Offer) {
	dbOffer = dbm.Offer{
		ID:        offer.ID,
		ProductID: offer.ProductID,
		Type:      offer.Type, // FIXME
		Value:     offer.Value,
	}

	dbOffer.ExpiratedAt.Time = *offer.ExpiratedAt

	if offer.CreatedAt != nil {
		if &offer.CreatedAt != nil {
			dbOffer.CreatedAt.Time = *offer.CreatedAt
		}
	}

	if offer.UpdatedAt != nil {
		if &offer.UpdatedAt != nil {
			dbOffer.UpdatedAt.Time = *offer.UpdatedAt
		}
	}

	if offer.ExpiratedAt != nil {
		if &offer.ExpiratedAt != nil {
			dbOffer.ExpiratedAt.Time = *offer.ExpiratedAt
		}
	}

	return
}

func parseOffersToDBOffers(offers models.Offers) (dbOffers dbm.Offers) {
	var dbOffer dbm.Offer

	for _, offer := range offers {
		dbOffer = parseOfferToDBOffer(offer)
		dbOffers = append(dbOffers, dbOffer)
	}
	return
}
