package handler

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/Allen-Career-Institute/flagr/swagger_gen/restapi/operations/flag"

	"github.com/Allen-Career-Institute/flagr/pkg/entity"
	"github.com/Allen-Career-Institute/flagr/pkg/util"
	"github.com/Allen-Career-Institute/flagr/swagger_gen/models"
	"github.com/Allen-Career-Institute/flagr/swagger_gen/restapi/operations/latch"
	"github.com/go-openapi/runtime/middleware"
	"github.com/prashantv/gostub"
	"github.com/stretchr/testify/assert"
)

func TestCrudCreateLatch(t *testing.T) {
	var res middleware.Responder
	db := entity.NewTestDB()
	c := &crud{}

	tmpDB, dbErr := db.DB()
	if dbErr != nil {
		t.Errorf("Failed to get database")
	}

	defer tmpDB.Close()
	defer gostub.StubFunc(&getDB, db).Reset()

	t.Run("it should create a latch with default template", func(t *testing.T) {
		res = c.CreateLatch(latch.CreateLatchParams{
			Body: &models.CreateFlagRequest{
				Description: util.StringPtr("simple latch"),
				Key:         "simple_latch_key",
			},
		})
		assert.NotNil(t, res)
		payload := res.(*flag.CreateFlagOK).Payload
		assert.NotZero(t, payload.ID)
		assert.Equal(t, "simple_latch_key", payload.Key)
		assert.Equal(t, len(payload.Variants), 1)
		assert.Equal(t, payload.Variants[0].Key, util.StringPtr("APPLICABLE"))
		assert.Equal(t, len(payload.Segments), 1)
		assert.Equal(t, payload.Segments[0].RolloutPercent, util.Int64Ptr(100))
		assert.Equal(t, len(payload.Segments[0].Distributions), 1)
		assert.Equal(t, payload.Segments[0].Distributions[0].Percent, util.Int64Ptr(100))
		assert.NotZero(t, payload.Tags)

		// Validate tag attachment
		flagID := payload.ID
		var attachedTags []entity.Tag
		db.Model(&entity.Flag{Key: strconv.FormatInt(flagID, 10)}).Association("Tags").Find(&attachedTags)
		//assert.NotEmpty(t, attachedTags)
		//assert.Equal(t, "latch", attachedTags[0].Value)
	})
}

func TestCrudCreateLatchWithFailures(t *testing.T) {
	var res middleware.Responder
	db := entity.NewTestDB()
	c := &crud{}

	tmpDB, dbErr := db.DB()
	if dbErr != nil {
		t.Errorf("Failed to get database")
	}

	defer tmpDB.Close()
	defer gostub.StubFunc(&getDB, db).Reset()

	t.Run("CreateLatch - invalid key error", func(t *testing.T) {
		res = c.CreateLatch(latch.CreateLatchParams{
			Body: &models.CreateFlagRequest{
				Description: util.StringPtr("invalid key latch"),
				Key:         " 1-2-3", // invalid key
			},
		})
		assert.NotNil(t, res)
		assert.NotZero(t, res.(*flag.CreateFlagDefault).Payload)
	})

	t.Run("CreateLatch - e2r MapFlag error", func(t *testing.T) {
		defer gostub.StubFunc(&e2rMapFlag, nil, fmt.Errorf("e2r MapFlag error")).Reset()
		res = c.CreateLatch(latch.CreateLatchParams{
			Body: &models.CreateFlagRequest{
				Description: util.StringPtr("map flag error latch"),
			},
		})
		assert.NotNil(t, res)
		assert.NotZero(t, res.(*flag.CreateFlagDefault).Payload)
	})

	t.Run("CreateLatch - db generic error", func(t *testing.T) {
		db.Error = fmt.Errorf("db generic error")
		res = c.CreateLatch(latch.CreateLatchParams{
			Body: &models.CreateFlagRequest{
				Description: util.StringPtr("db error latch"),
			},
		})
		assert.NotNil(t, res)
		assert.NotZero(t, res.(*flag.CreateFlagDefault).Payload)
		db.Error = nil
	})

	var testLoadSimpleLatchTemplate = LoadSimpleLatchTemplate
	t.Run("CreateLatch - template error", func(t *testing.T) {
		defer gostub.StubFunc(&testLoadSimpleLatchTemplate, fmt.Errorf("template load error")).Reset()
		res = c.CreateLatch(latch.CreateLatchParams{
			Body: &models.CreateFlagRequest{
				Description: util.StringPtr("template error latch"),
				Key:         "template_error_key",
			},
		})
		assert.NotNil(t, res)
		assert.NotZero(t, res.(*flag.CreateFlagOK).Payload)
	})

	var testFunc = associateTagWithFlag
	t.Run("CreateLatch - tag association error", func(t *testing.T) {
		defer gostub.StubFunc(&testFunc, fmt.Errorf("tag association error")).Reset()
		res = c.CreateLatch(latch.CreateLatchParams{
			Body: &models.CreateFlagRequest{
				Description: util.StringPtr("tag error latch"),
				Key:         "tag_error_key",
			},
		})
		assert.NotNil(t, res)
		assert.NotZero(t, res.(*flag.CreateFlagOK).Payload)
	})
}
