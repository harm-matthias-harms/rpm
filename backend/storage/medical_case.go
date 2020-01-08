package storage

import (
	"context"
	"errors"
	"io"
	"mime/multipart"

	"github.com/harm-matthias-harms/rpm/backend/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// CountMedicalCases returns the number of total documents inserted
func CountMedicalCases(ctx context.Context, params map[string]interface{}) (int64, error) {
	if len(params) == 0 {
		return mcCollection().CountDocuments(ctx, bson.M{})
	}
	return mcCollection().CountDocuments(ctx, params)
}

// GetMedicalCases returns an array of the medicalcases in the short version
func GetMedicalCases(ctx context.Context, params map[string]interface{}, page int, pageSize int) (result []model.MedicalCaseShort, err error) {
	c := mcCollection()
	options := options.Find()
	options.SetSort(bson.D{{Key: "_id", Value: -1}})
	options.SetSkip(int64((page - 1) * pageSize))
	options.SetLimit(int64(pageSize))
	cursor, err := c.Find(ctx, params, options)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var preset model.MedicalCaseShort
		cursor.Decode(&preset)
		result = append(result, preset)
	}
	return result, nil
}

// FindMedicalCase finds a medical case
func FindMedicalCase(ctx context.Context, id primitive.ObjectID) (result *model.MedicalCase, err error) {
	c := mcCollection()
	result = new(model.MedicalCase)
	err = c.FindOne(ctx, bson.D{{Key: "_id", Value: id}}).Decode(result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateMedicalCase will store a medical case
func CreateMedicalCase(ctx context.Context, mc *model.MedicalCase) (err error) {
	c := mcCollection()
	res, err := c.InsertOne(ctx, mc)
	mc.ID = res.InsertedID.(primitive.ObjectID)
	return
}

// UpdateMedicalCase updates a medical case
func UpdateMedicalCase(ctx context.Context, mc *model.MedicalCase) (err error) {
	c := mcCollection()
	res, err := c.ReplaceOne(ctx, bson.M{"_id": mc.ID}, mc)
	if res.MatchedCount == 0 {
		return errors.New("no document was found")
	}
	return
}

// StoreMedicalCaseFile stores the provided documents to a medical case
func StoreMedicalCaseFile(mc *model.MedicalCase, fh *multipart.FileHeader) (err error) {
	bucket, err := fileBucket()
	if err != nil {
		return
	}
	src, err := fh.Open()
	if err != nil {
		return
	}
	defer src.Close()
	id, err := bucket.UploadFromStream(fh.Filename, src)
	if err != nil {
		return
	}
	mc.Files = append(mc.Files, model.MedicalCaseFile{ID: id, Name: fh.Filename, Size: fh.Size})
	return
}

// GetMedicalCaseFile returns a specific File for a medical case
func GetMedicalCaseFile(id primitive.ObjectID, stream io.Writer) (err error) {
	bucket, err := fileBucket()
	if err != nil {
		return
	}
	_, err = bucket.DownloadToStream(id, stream)
	return
}

func mcCollection() *mongo.Collection {
	return MongoSession.Collection("medicalCases")
}

func fileBucket() (bucket *gridfs.Bucket, err error) {
	bucket, err = gridfs.NewBucket(MongoSession, options.GridFSBucket().SetName("files"))
	return
}
