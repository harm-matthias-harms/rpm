package storage

import (
	"bufio"
	"bytes"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/harm-matthias-harms/rpm/backend/model"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestMedicalCase(t *testing.T) {
	// setup
	_ = SetMongoDatabase()

	// models
	mc := &model.MedicalCase{Title: "title"}
	mc2 := &model.MedicalCase{Title: "title 2"}
	noneExistingMC := &model.MedicalCase{Title: "title"}

	// tests
	t.Run("create medical case", func(t *testing.T) {
		err := CreateMedicalCase(nil, mc)
		assert.NoError(t, err)
		err = CreateMedicalCase(nil, mc2)
		assert.NoError(t, err)
	})

	t.Run("update medical case", func(t *testing.T) {
		mc.MedicalHistroy.Allergies = "allergy"
		err := UpdateMedicalCase(nil, mc)
		assert.NoError(t, err)

		// medical case doesn't exists
		err = UpdateMedicalCase(nil, noneExistingMC)
		if assert.Error(t, err) {
			assert.Equal(t, "no document was found", err.Error())
		}
	})

	t.Run("find medical case", func(t *testing.T) {
		mcFound, err := FindMedicalCase(nil, mc.ID)
		if assert.NoError(t, err) {
			assert.Equal(t, "title", mcFound.Title)
			assert.Equal(t, "allergy", mcFound.MedicalHistroy.Allergies)
		}

		// doesn't find none existing medical case
		_, err = FindMedicalCase(nil, noneExistingMC.ID)
		assert.Error(t, err)
	})

	t.Run("get medical case", func(t *testing.T) {
		result, err := GetMedicalCases(nil, nil, 1, 1)
		if assert.NoError(t, err) {
			assert.Equal(t, 1, len(result))
		}
		// test filter
		filter := map[string]interface{}{"title": "title"}
		result, err = GetMedicalCases(nil, filter, 1, 1)
		if assert.NoError(t, err) {
			assert.Equal(t, 1, len(result))
		}
		// test multiple filters
		filter["medicalHistory.allergies"] = "allergy"
		result, err = GetMedicalCases(nil, filter, 1, 1)
		if assert.NoError(t, err) {
			assert.Equal(t, 1, len(result))
			assert.Equal(t, mc.ID, result[0].ID)
		}
		// test load all
		result, err = GetMedicalCases(nil, nil, 1, 0)
		if assert.NoError(t, err) {
			assert.Greater(t, len(result), 1)
		}
	})

	t.Run("count medical case", func(t *testing.T) {
		count, err := CountMedicalCases(nil, nil)
		if assert.NoError(t, err) {
			assert.Greater(t, count, int64(0))
		}
		filter := map[string]interface{}{"title": "title"}
		count, err = CountMedicalCases(nil, filter)
		if assert.NoError(t, err) {
			assert.Equal(t, int64(1), count)
		}
	})

	t.Run("create file", func(t *testing.T) {
		// setup
		file, _ := os.Open("./medical_case_test.go")
		defer file.Close()
		body := new(bytes.Buffer)
		writer := multipart.NewWriter(body)
		part, _ := writer.CreateFormFile("files", filepath.Base("./medical_case_test.go"))
		io.Copy(part, file)
		writer.Close()

		// reading form
		r := multipart.NewReader(body, writer.Boundary())
		form, _ := r.ReadForm(10)

		// test file create
		for _, file := range form.File["files"] {
			err := StoreMedicalCaseFile(mc, file)
			assert.NoError(t, err)
		}
	})

	t.Run("get file", func(t *testing.T) {
		// setup
		bucket, _ := fileBucket()

		// upload a file
		ustream, _ := bucket.OpenUploadStream("test.txt", options.GridFSUpload())
		_, _ = ustream.Write([]byte("test file"))
		fileID := ustream.FileID
		ustream.Close()

		// test get
		var b bytes.Buffer
		w := bufio.NewWriter(&b)
		err := GetMedicalCaseFile(fileID.(primitive.ObjectID), w)
		assert.NoError(t, err)
		assert.Equal(t, "test file", b.String())

		// cleanup
		_ = bucket.Delete(fileID)
	})

	// cleanup
	resetMedicalCaseFile(mc)
	resetMedicalCaseDatabase(mc)
	resetMedicalCaseDatabase(mc2)
}

func resetMedicalCaseFile(mc *model.MedicalCase) {
	bucket, _ := fileBucket()
	for _, f := range mc.Files {
		_ = bucket.Delete(f.ID)
	}
}

func resetMedicalCaseDatabase(mc *model.MedicalCase) {
	_, _ = mcCollection().DeleteMany(nil, bson.M{"title": mc.Title})
}
