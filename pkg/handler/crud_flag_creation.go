package handler

import (
	"fmt"
	"github.com/Allen-Career-Institute/flagr/pkg/entity"
	"github.com/Allen-Career-Institute/flagr/pkg/util"
	"github.com/Allen-Career-Institute/flagr/swagger_gen/restapi/operations/flag"
	"github.com/Allen-Career-Institute/flagr/swagger_gen/restapi/operations/latch"
	"github.com/go-openapi/runtime/middleware"
	"gorm.io/gorm"
)

func (c *crud) CreateFlag(params flag.CreateFlagParams) middleware.Responder {
	f := &entity.Flag{}
	if params.Body != nil {
		f.Description = util.SafeString(params.Body.Description)
		f.CreatedBy = getSubjectFromRequest(params.HTTPRequest)

		key, err := entity.CreateFlagKey(params.Body.Key)
		if err != nil {
			return flag.NewCreateFlagDefault(400).WithPayload(
				ErrorMessage("cannot create flag. %s", err))
		}
		f.Key = key
	}

	tx := getDB().Begin()

	if err := tx.Create(f).Error; err != nil {
		tx.Rollback()
		return flag.NewCreateFlagDefault(500).WithPayload(
			ErrorMessage("cannot create flag. %s", err))
	}

	if params.Body.Template == "simple_boolean_flag" {
		if err := LoadSimpleBooleanFlagTemplate(f, tx); err != nil {
			tx.Rollback()
			return flag.NewCreateFlagDefault(500).WithPayload(
				ErrorMessage("cannot create flag. %s", err))
		}
	} else if params.Body.Template != "" {
		return flag.NewCreateFlagDefault(400).WithPayload(
			ErrorMessage("unknown value for template: %s", params.Body.Template))
	}

	err := tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return flag.NewCreateFlagDefault(500).WithPayload(ErrorMessage("%s", err))
	}

	resp := flag.NewCreateFlagOK()
	payload, err := e2rMapFlag(f)
	if err != nil {
		return flag.NewCreateFlagDefault(500).WithPayload(
			ErrorMessage("cannot map flag. %s", err))
	}
	resp.SetPayload(payload)

	entity.SaveFlagSnapshot(getDB(), f.ID, getSubjectFromRequest(params.HTTPRequest))

	return resp
}

func (c *crud) CreateLatch(params latch.CreateLatchParams) middleware.Responder {
	f := &entity.Flag{}
	if params.Body != nil {
		f.Description = util.SafeString(params.Body.Description)
		f.CreatedBy = getSubjectFromRequest(params.HTTPRequest)

		key, err := entity.CreateFlagKey(params.Body.Key)
		if err != nil {
			return flag.NewCreateFlagDefault(400).WithPayload(
				ErrorMessage("cannot create flag. %s", err))
		}
		f.Key = key
	}

	tx := getDB().Begin()

	if err := tx.Create(f).Error; err != nil {
		tx.Rollback()
		return flag.NewCreateFlagDefault(500).WithPayload(
			ErrorMessage("cannot create flag. %s", err))
	}

	if err := LoadSimpleLatchTemplate(f, tx); err != nil {
		tx.Rollback()
		return flag.NewCreateFlagDefault(500).WithPayload(
			ErrorMessage("cannot create flag. %s", err))
	}

	err := tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return flag.NewCreateFlagDefault(500).WithPayload(ErrorMessage("%s", err))
	}

	resp := flag.NewCreateFlagOK()
	payload, err := e2rMapFlag(f)
	if err != nil {
		return flag.NewCreateFlagDefault(500).WithPayload(
			ErrorMessage("cannot map flag. %s", err))
	}
	resp.SetPayload(payload)

	entity.SaveFlagSnapshot(getDB(), f.ID, getSubjectFromRequest(params.HTTPRequest))

	return resp
}

// LoadSimpleBooleanFlagTemplate loads the simple boolean flag template into
// a new flag. It creates a single segment, variant ('on'), and distribution.
func LoadSimpleBooleanFlagTemplate(flag *entity.Flag, tx *gorm.DB) error {
	// Create our default segment
	s := &entity.Segment{}
	s.FlagID = flag.ID
	s.RolloutPercent = uint(100)
	s.Rank = entity.SegmentDefaultRank

	if err := tx.Create(s).Error; err != nil {
		return err
	}

	// .. and our default Variant
	v := &entity.Variant{}
	v.FlagID = flag.ID
	v.Key = "on"

	if err := tx.Create(v).Error; err != nil {
		return err
	}

	// .. and our default Distribution
	d := &entity.Distribution{}
	d.SegmentID = s.ID
	d.VariantID = v.ID
	d.VariantKey = v.Key
	d.Percent = uint(100)

	if err := tx.Create(d).Error; err != nil {
		return err
	}

	s.Distributions = append(s.Distributions, *d)
	flag.Variants = append(flag.Variants, *v)
	flag.Segments = append(flag.Segments, *s)

	return nil
}

// LoadSimpleLatchTemplate loads the simple boolean flag template into
// a new flag. It creates a single segment, variant ('on'), and distribution.
func LoadSimpleLatchTemplate(flag *entity.Flag, tx *gorm.DB) error {
	// create a tag for latch
	latchTagStr := "latch"
	t := &entity.Tag{}
	t.Value = util.SafeString(latchTagStr)
	if ok, reason := util.IsSafeValue(t.Value); !ok {
		return fmt.Errorf("error creating tag: %v", reason)
	}

	tx.Where("value = ?", util.SafeString(latchTagStr)).Find(t) // Find the existing tag to associate if it exists
	// associate tag to flag
	if err := tx.Model(flag).Association("Tags").Append(t); err != nil {
		return fmt.Errorf("error creating tag: %v, err while associating with flags", err)
	}

	// Create our default segment
	s := &entity.Segment{}
	s.FlagID = flag.ID
	s.RolloutPercent = uint(100)
	s.Rank = entity.SegmentDefaultRank

	if err := tx.Create(s).Error; err != nil {
		return err
	}

	// .. and our default Variant
	v := &entity.Variant{}
	v.FlagID = flag.ID
	v.Key = "APPLICABLE"

	if err := tx.Create(v).Error; err != nil {
		return err
	}

	// .. and our default Distribution
	d := &entity.Distribution{}
	d.SegmentID = s.ID
	d.VariantID = v.ID
	d.VariantKey = v.Key
	d.Percent = uint(100)

	if err := tx.Create(d).Error; err != nil {
		return err
	}

	s.Distributions = append(s.Distributions, *d)
	flag.Variants = append(flag.Variants, *v)
	flag.Segments = append(flag.Segments, *s)

	return nil
}
