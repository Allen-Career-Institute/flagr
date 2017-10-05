package entity

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

// ===== BEGIN of all query sets

// ===== BEGIN of query set SegmentQuerySet

// SegmentQuerySet is an queryset type for Segment
type SegmentQuerySet struct {
	db *gorm.DB
}

// NewSegmentQuerySet constructs new SegmentQuerySet
func NewSegmentQuerySet(db *gorm.DB) SegmentQuerySet {
	return SegmentQuerySet{
		db: db.Model(&Segment{}),
	}
}

func (qs SegmentQuerySet) w(db *gorm.DB) SegmentQuerySet {
	return NewSegmentQuerySet(db)
}

// All is an autogenerated method
// nolint: dupl
func (qs SegmentQuerySet) All(ret *[]Segment) error {
	return qs.db.Find(ret).Error
}

// Count is an autogenerated method
// nolint: dupl
func (qs SegmentQuerySet) Count() (int, error) {
	var count int
	err := qs.db.Count(&count).Error
	return count, err
}

// Create is an autogenerated method
// nolint: dupl
func (o *Segment) Create(db *gorm.DB) error {
	return db.Create(o).Error
}

// CreatedAtEq is an autogenerated method
// nolint: dupl
func (qs SegmentQuerySet) CreatedAtEq(createdAt time.Time) SegmentQuerySet {
	return qs.w(qs.db.Where("created_at = ?", createdAt))
}

// CreatedAtGt is an autogenerated method
// nolint: dupl
func (qs SegmentQuerySet) CreatedAtGt(createdAt time.Time) SegmentQuerySet {
	return qs.w(qs.db.Where("created_at > ?", createdAt))
}

// CreatedAtGte is an autogenerated method
// nolint: dupl
func (qs SegmentQuerySet) CreatedAtGte(createdAt time.Time) SegmentQuerySet {
	return qs.w(qs.db.Where("created_at >= ?", createdAt))
}

// CreatedAtLt is an autogenerated method
// nolint: dupl
func (qs SegmentQuerySet) CreatedAtLt(createdAt time.Time) SegmentQuerySet {
	return qs.w(qs.db.Where("created_at < ?", createdAt))
}

// CreatedAtLte is an autogenerated method
// nolint: dupl
func (qs SegmentQuerySet) CreatedAtLte(createdAt time.Time) SegmentQuerySet {
	return qs.w(qs.db.Where("created_at <= ?", createdAt))
}

// CreatedAtNe is an autogenerated method
// nolint: dupl
func (qs SegmentQuerySet) CreatedAtNe(createdAt time.Time) SegmentQuerySet {
	return qs.w(qs.db.Where("created_at != ?", createdAt))
}

// Delete is an autogenerated method
// nolint: dupl
func (o *Segment) Delete(db *gorm.DB) error {
	return db.Delete(o).Error
}

// Delete is an autogenerated method
// nolint: dupl
func (qs SegmentQuerySet) Delete() error {
	return qs.db.Delete(Segment{}).Error
}

// DeletedAtEq is an autogenerated method
// nolint: dupl
func (qs SegmentQuerySet) DeletedAtEq(deletedAt time.Time) SegmentQuerySet {
	return qs.w(qs.db.Where("deleted_at = ?", deletedAt))
}

// DeletedAtGt is an autogenerated method
// nolint: dupl
func (qs SegmentQuerySet) DeletedAtGt(deletedAt time.Time) SegmentQuerySet {
	return qs.w(qs.db.Where("deleted_at > ?", deletedAt))
}

// DeletedAtGte is an autogenerated method
// nolint: dupl
func (qs SegmentQuerySet) DeletedAtGte(deletedAt time.Time) SegmentQuerySet {
	return qs.w(qs.db.Where("deleted_at >= ?", deletedAt))
}

// DeletedAtIsNotNull is an autogenerated method
// nolint: dupl
func (qs SegmentQuerySet) DeletedAtIsNotNull() SegmentQuerySet {
	return qs.w(qs.db.Where("deleted_at IS NOT NULL"))
}

// DeletedAtIsNull is an autogenerated method
// nolint: dupl
func (qs SegmentQuerySet) DeletedAtIsNull() SegmentQuerySet {
	return qs.w(qs.db.Where("deleted_at IS NULL"))
}

// DeletedAtLt is an autogenerated method
// nolint: dupl
func (qs SegmentQuerySet) DeletedAtLt(deletedAt time.Time) SegmentQuerySet {
	return qs.w(qs.db.Where("deleted_at < ?", deletedAt))
}

// DeletedAtLte is an autogenerated method
// nolint: dupl
func (qs SegmentQuerySet) DeletedAtLte(deletedAt time.Time) SegmentQuerySet {
	return qs.w(qs.db.Where("deleted_at <= ?", deletedAt))
}

// DeletedAtNe is an autogenerated method
// nolint: dupl
func (qs SegmentQuerySet) DeletedAtNe(deletedAt time.Time) SegmentQuerySet {
	return qs.w(qs.db.Where("deleted_at != ?", deletedAt))
}

// DescriptionEq is an autogenerated method
// nolint: dupl
func (qs SegmentQuerySet) DescriptionEq(description string) SegmentQuerySet {
	return qs.w(qs.db.Where("description = ?", description))
}

// DescriptionIn is an autogenerated method
// nolint: dupl
func (qs SegmentQuerySet) DescriptionIn(description string, descriptionRest ...string) SegmentQuerySet {
	iArgs := []interface{}{description}
	for _, arg := range descriptionRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("description IN (?)", iArgs))
}

// DescriptionNe is an autogenerated method
// nolint: dupl
func (qs SegmentQuerySet) DescriptionNe(description string) SegmentQuerySet {
	return qs.w(qs.db.Where("description != ?", description))
}

// DescriptionNotIn is an autogenerated method
// nolint: dupl
func (qs SegmentQuerySet) DescriptionNotIn(description string, descriptionRest ...string) SegmentQuerySet {
	iArgs := []interface{}{description}
	for _, arg := range descriptionRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("description NOT IN (?)", iArgs))
}

// DistributionIsNotNull is an autogenerated method
// nolint: dupl
func (qs SegmentQuerySet) DistributionIsNotNull() SegmentQuerySet {
	return qs.w(qs.db.Where("distribution IS NOT NULL"))
}

// DistributionIsNull is an autogenerated method
// nolint: dupl
func (qs SegmentQuerySet) DistributionIsNull() SegmentQuerySet {
	return qs.w(qs.db.Where("distribution IS NULL"))
}

// FlagIDEq is an autogenerated method
// nolint: dupl
func (qs SegmentQuerySet) FlagIDEq(flagID uint) SegmentQuerySet {
	return qs.w(qs.db.Where("flag_id = ?", flagID))
}

// FlagIDGt is an autogenerated method
// nolint: dupl
func (qs SegmentQuerySet) FlagIDGt(flagID uint) SegmentQuerySet {
	return qs.w(qs.db.Where("flag_id > ?", flagID))
}

// FlagIDGte is an autogenerated method
// nolint: dupl
func (qs SegmentQuerySet) FlagIDGte(flagID uint) SegmentQuerySet {
	return qs.w(qs.db.Where("flag_id >= ?", flagID))
}

// FlagIDIn is an autogenerated method
// nolint: dupl
func (qs SegmentQuerySet) FlagIDIn(flagID uint, flagIDRest ...uint) SegmentQuerySet {
	iArgs := []interface{}{flagID}
	for _, arg := range flagIDRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("flag_id IN (?)", iArgs))
}

// FlagIDLt is an autogenerated method
// nolint: dupl
func (qs SegmentQuerySet) FlagIDLt(flagID uint) SegmentQuerySet {
	return qs.w(qs.db.Where("flag_id < ?", flagID))
}

// FlagIDLte is an autogenerated method
// nolint: dupl
func (qs SegmentQuerySet) FlagIDLte(flagID uint) SegmentQuerySet {
	return qs.w(qs.db.Where("flag_id <= ?", flagID))
}

// FlagIDNe is an autogenerated method
// nolint: dupl
func (qs SegmentQuerySet) FlagIDNe(flagID uint) SegmentQuerySet {
	return qs.w(qs.db.Where("flag_id != ?", flagID))
}

// FlagIDNotIn is an autogenerated method
// nolint: dupl
func (qs SegmentQuerySet) FlagIDNotIn(flagID uint, flagIDRest ...uint) SegmentQuerySet {
	iArgs := []interface{}{flagID}
	for _, arg := range flagIDRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("flag_id NOT IN (?)", iArgs))
}

// GetUpdater is an autogenerated method
// nolint: dupl
func (qs SegmentQuerySet) GetUpdater() SegmentUpdater {
	return NewSegmentUpdater(qs.db)
}

// IDEq is an autogenerated method
// nolint: dupl
func (qs SegmentQuerySet) IDEq(ID uint) SegmentQuerySet {
	return qs.w(qs.db.Where("id = ?", ID))
}

// IDGt is an autogenerated method
// nolint: dupl
func (qs SegmentQuerySet) IDGt(ID uint) SegmentQuerySet {
	return qs.w(qs.db.Where("id > ?", ID))
}

// IDGte is an autogenerated method
// nolint: dupl
func (qs SegmentQuerySet) IDGte(ID uint) SegmentQuerySet {
	return qs.w(qs.db.Where("id >= ?", ID))
}

// IDIn is an autogenerated method
// nolint: dupl
func (qs SegmentQuerySet) IDIn(ID uint, IDRest ...uint) SegmentQuerySet {
	iArgs := []interface{}{ID}
	for _, arg := range IDRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("id IN (?)", iArgs))
}

// IDLt is an autogenerated method
// nolint: dupl
func (qs SegmentQuerySet) IDLt(ID uint) SegmentQuerySet {
	return qs.w(qs.db.Where("id < ?", ID))
}

// IDLte is an autogenerated method
// nolint: dupl
func (qs SegmentQuerySet) IDLte(ID uint) SegmentQuerySet {
	return qs.w(qs.db.Where("id <= ?", ID))
}

// IDNe is an autogenerated method
// nolint: dupl
func (qs SegmentQuerySet) IDNe(ID uint) SegmentQuerySet {
	return qs.w(qs.db.Where("id != ?", ID))
}

// IDNotIn is an autogenerated method
// nolint: dupl
func (qs SegmentQuerySet) IDNotIn(ID uint, IDRest ...uint) SegmentQuerySet {
	iArgs := []interface{}{ID}
	for _, arg := range IDRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("id NOT IN (?)", iArgs))
}

// Limit is an autogenerated method
// nolint: dupl
func (qs SegmentQuerySet) Limit(limit int) SegmentQuerySet {
	return qs.w(qs.db.Limit(limit))
}

// One is used to retrieve one result. It returns gorm.ErrRecordNotFound
// if nothing was fetched
func (qs SegmentQuerySet) One(ret *Segment) error {
	return qs.db.First(ret).Error
}

// OrderAscByCreatedAt is an autogenerated method
// nolint: dupl
func (qs SegmentQuerySet) OrderAscByCreatedAt() SegmentQuerySet {
	return qs.w(qs.db.Order("created_at ASC"))
}

// OrderAscByDeletedAt is an autogenerated method
// nolint: dupl
func (qs SegmentQuerySet) OrderAscByDeletedAt() SegmentQuerySet {
	return qs.w(qs.db.Order("deleted_at ASC"))
}

// OrderAscByFlagID is an autogenerated method
// nolint: dupl
func (qs SegmentQuerySet) OrderAscByFlagID() SegmentQuerySet {
	return qs.w(qs.db.Order("flag_id ASC"))
}

// OrderAscByID is an autogenerated method
// nolint: dupl
func (qs SegmentQuerySet) OrderAscByID() SegmentQuerySet {
	return qs.w(qs.db.Order("id ASC"))
}

// OrderAscByRank is an autogenerated method
// nolint: dupl
func (qs SegmentQuerySet) OrderAscByRank() SegmentQuerySet {
	return qs.w(qs.db.Order("rank ASC"))
}

// OrderAscByUpdatedAt is an autogenerated method
// nolint: dupl
func (qs SegmentQuerySet) OrderAscByUpdatedAt() SegmentQuerySet {
	return qs.w(qs.db.Order("updated_at ASC"))
}

// OrderDescByCreatedAt is an autogenerated method
// nolint: dupl
func (qs SegmentQuerySet) OrderDescByCreatedAt() SegmentQuerySet {
	return qs.w(qs.db.Order("created_at DESC"))
}

// OrderDescByDeletedAt is an autogenerated method
// nolint: dupl
func (qs SegmentQuerySet) OrderDescByDeletedAt() SegmentQuerySet {
	return qs.w(qs.db.Order("deleted_at DESC"))
}

// OrderDescByFlagID is an autogenerated method
// nolint: dupl
func (qs SegmentQuerySet) OrderDescByFlagID() SegmentQuerySet {
	return qs.w(qs.db.Order("flag_id DESC"))
}

// OrderDescByID is an autogenerated method
// nolint: dupl
func (qs SegmentQuerySet) OrderDescByID() SegmentQuerySet {
	return qs.w(qs.db.Order("id DESC"))
}

// OrderDescByRank is an autogenerated method
// nolint: dupl
func (qs SegmentQuerySet) OrderDescByRank() SegmentQuerySet {
	return qs.w(qs.db.Order("rank DESC"))
}

// OrderDescByUpdatedAt is an autogenerated method
// nolint: dupl
func (qs SegmentQuerySet) OrderDescByUpdatedAt() SegmentQuerySet {
	return qs.w(qs.db.Order("updated_at DESC"))
}

// PreloadDistribution is an autogenerated method
// nolint: dupl
func (qs SegmentQuerySet) PreloadDistribution() SegmentQuerySet {
	return qs.w(qs.db.Preload("Distribution"))
}

// RankEq is an autogenerated method
// nolint: dupl
func (qs SegmentQuerySet) RankEq(rank uint) SegmentQuerySet {
	return qs.w(qs.db.Where("rank = ?", rank))
}

// RankGt is an autogenerated method
// nolint: dupl
func (qs SegmentQuerySet) RankGt(rank uint) SegmentQuerySet {
	return qs.w(qs.db.Where("rank > ?", rank))
}

// RankGte is an autogenerated method
// nolint: dupl
func (qs SegmentQuerySet) RankGte(rank uint) SegmentQuerySet {
	return qs.w(qs.db.Where("rank >= ?", rank))
}

// RankIn is an autogenerated method
// nolint: dupl
func (qs SegmentQuerySet) RankIn(rank uint, rankRest ...uint) SegmentQuerySet {
	iArgs := []interface{}{rank}
	for _, arg := range rankRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("rank IN (?)", iArgs))
}

// RankLt is an autogenerated method
// nolint: dupl
func (qs SegmentQuerySet) RankLt(rank uint) SegmentQuerySet {
	return qs.w(qs.db.Where("rank < ?", rank))
}

// RankLte is an autogenerated method
// nolint: dupl
func (qs SegmentQuerySet) RankLte(rank uint) SegmentQuerySet {
	return qs.w(qs.db.Where("rank <= ?", rank))
}

// RankNe is an autogenerated method
// nolint: dupl
func (qs SegmentQuerySet) RankNe(rank uint) SegmentQuerySet {
	return qs.w(qs.db.Where("rank != ?", rank))
}

// RankNotIn is an autogenerated method
// nolint: dupl
func (qs SegmentQuerySet) RankNotIn(rank uint, rankRest ...uint) SegmentQuerySet {
	iArgs := []interface{}{rank}
	for _, arg := range rankRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("rank NOT IN (?)", iArgs))
}

// SetCreatedAt is an autogenerated method
// nolint: dupl
func (u SegmentUpdater) SetCreatedAt(createdAt time.Time) SegmentUpdater {
	u.fields[string(SegmentDBSchema.CreatedAt)] = createdAt
	return u
}

// SetDescription is an autogenerated method
// nolint: dupl
func (u SegmentUpdater) SetDescription(description string) SegmentUpdater {
	u.fields[string(SegmentDBSchema.Description)] = description
	return u
}

// SetFlagID is an autogenerated method
// nolint: dupl
func (u SegmentUpdater) SetFlagID(flagID uint) SegmentUpdater {
	u.fields[string(SegmentDBSchema.FlagID)] = flagID
	return u
}

// SetID is an autogenerated method
// nolint: dupl
func (u SegmentUpdater) SetID(ID uint) SegmentUpdater {
	u.fields[string(SegmentDBSchema.ID)] = ID
	return u
}

// SetRank is an autogenerated method
// nolint: dupl
func (u SegmentUpdater) SetRank(rank uint) SegmentUpdater {
	u.fields[string(SegmentDBSchema.Rank)] = rank
	return u
}

// SetUpdatedAt is an autogenerated method
// nolint: dupl
func (u SegmentUpdater) SetUpdatedAt(updatedAt time.Time) SegmentUpdater {
	u.fields[string(SegmentDBSchema.UpdatedAt)] = updatedAt
	return u
}

// Update is an autogenerated method
// nolint: dupl
func (u SegmentUpdater) Update() error {
	return u.db.Updates(u.fields).Error
}

// UpdateNum is an autogenerated method
// nolint: dupl
func (u SegmentUpdater) UpdateNum() (int64, error) {
	db := u.db.Updates(u.fields)
	return db.RowsAffected, db.Error
}

// UpdatedAtEq is an autogenerated method
// nolint: dupl
func (qs SegmentQuerySet) UpdatedAtEq(updatedAt time.Time) SegmentQuerySet {
	return qs.w(qs.db.Where("updated_at = ?", updatedAt))
}

// UpdatedAtGt is an autogenerated method
// nolint: dupl
func (qs SegmentQuerySet) UpdatedAtGt(updatedAt time.Time) SegmentQuerySet {
	return qs.w(qs.db.Where("updated_at > ?", updatedAt))
}

// UpdatedAtGte is an autogenerated method
// nolint: dupl
func (qs SegmentQuerySet) UpdatedAtGte(updatedAt time.Time) SegmentQuerySet {
	return qs.w(qs.db.Where("updated_at >= ?", updatedAt))
}

// UpdatedAtLt is an autogenerated method
// nolint: dupl
func (qs SegmentQuerySet) UpdatedAtLt(updatedAt time.Time) SegmentQuerySet {
	return qs.w(qs.db.Where("updated_at < ?", updatedAt))
}

// UpdatedAtLte is an autogenerated method
// nolint: dupl
func (qs SegmentQuerySet) UpdatedAtLte(updatedAt time.Time) SegmentQuerySet {
	return qs.w(qs.db.Where("updated_at <= ?", updatedAt))
}

// UpdatedAtNe is an autogenerated method
// nolint: dupl
func (qs SegmentQuerySet) UpdatedAtNe(updatedAt time.Time) SegmentQuerySet {
	return qs.w(qs.db.Where("updated_at != ?", updatedAt))
}

// ===== END of query set SegmentQuerySet

// ===== BEGIN of Segment modifiers

type segmentDBSchemaField string

// SegmentDBSchema stores db field names of Segment
var SegmentDBSchema = struct {
	ID           segmentDBSchemaField
	CreatedAt    segmentDBSchemaField
	UpdatedAt    segmentDBSchemaField
	DeletedAt    segmentDBSchemaField
	FlagID       segmentDBSchemaField
	Description  segmentDBSchemaField
	Rank         segmentDBSchemaField
	Constraints  segmentDBSchemaField
	Distribution segmentDBSchemaField
}{

	ID:           segmentDBSchemaField("id"),
	CreatedAt:    segmentDBSchemaField("created_at"),
	UpdatedAt:    segmentDBSchemaField("updated_at"),
	DeletedAt:    segmentDBSchemaField("deleted_at"),
	FlagID:       segmentDBSchemaField("flag_id"),
	Description:  segmentDBSchemaField("description"),
	Rank:         segmentDBSchemaField("rank"),
	Constraints:  segmentDBSchemaField("constraints"),
	Distribution: segmentDBSchemaField("distribution"),
}

// Update updates Segment fields by primary key
func (o *Segment) Update(db *gorm.DB, fields ...segmentDBSchemaField) error {
	dbNameToFieldName := map[string]interface{}{
		"id":           o.ID,
		"created_at":   o.CreatedAt,
		"updated_at":   o.UpdatedAt,
		"deleted_at":   o.DeletedAt,
		"flag_id":      o.FlagID,
		"description":  o.Description,
		"rank":         o.Rank,
		"constraints":  o.Constraints,
		"distribution": o.Distribution,
	}
	u := map[string]interface{}{}
	for _, f := range fields {
		fs := string(f)
		u[fs] = dbNameToFieldName[fs]
	}
	if err := db.Model(o).Updates(u).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return err
		}

		return fmt.Errorf("can't update Segment %v fields %v: %s",
			o, fields, err)
	}

	return nil
}

// SegmentUpdater is an Segment updates manager
type SegmentUpdater struct {
	fields map[string]interface{}
	db     *gorm.DB
}

// NewSegmentUpdater creates new Segment updater
func NewSegmentUpdater(db *gorm.DB) SegmentUpdater {
	return SegmentUpdater{
		fields: map[string]interface{}{},
		db:     db.Model(&Segment{}),
	}
}

// ===== END of Segment modifiers

// ===== END of all query sets
