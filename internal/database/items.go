package database

import (
	"context"
	"errors"
	"github.com/wizedkyle/artifactsmmo/v2/internal/models"
	"github.com/wizedkyle/artifactsmmo/v2/internal/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// CreateItems
// Bulk creates new items in the database. This is primarily used for sync items from the Artifacts API.
func (d *database) CreateItems(items []interface{}) error {
	opts := options.InsertMany().SetOrdered(false)
	_, err := d.ItemsCollection.InsertMany(context.TODO(), items, opts)
	if err != nil {
		return err
	}
	return nil
}

// DeleteItems
// Deletes all items in the item database with the specified type. This is only used by the sync task as it's quicker than updating each record.
func (d *database) DeleteItems(itemType string) error {
	_, err := d.ItemsCollection.DeleteMany(context.TODO(), bson.D{{"type", itemType}})
	if err != nil {
		return err
	}
	return nil
}

// GetItem
// Returns an item from the database based on the specified item name.
func (d *database) GetItem(code string) (*models.ItemDetails, error) {
	var result models.ItemDetails
	err := d.ItemsCollection.FindOne(context.TODO(), bson.D{{"code", code}}).Decode(&result)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, utils.ErrItemNotFound
	} else if err != nil {
		return nil, err
	}
	return &result, nil
}

// ListItems
// Returns items from the database based on the supplied filters.
func (d *database) ListItems(limit int64, params models.ListItemParameters) (*[]models.ItemDetails, error) {
	var items []models.ItemDetails
	opts := options.Find().SetLimit(limit)
	filter := bson.D{
		{"type", params.Type},
		{"subType", params.SubType},
		{"level", params.Level},
	}
	cursor, err := d.ItemsCollection.Find(context.TODO(), filter, opts)
	if err != nil {
		return nil, err
	}
	for cursor.Next(context.TODO()) {
		var item models.ItemDetails
		err := cursor.Decode(&item)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	cursor.Close(context.TODO())
	if len(items) == 0 {
		return nil, utils.ErrItemsNotFound
	}
	return &items, nil
}
