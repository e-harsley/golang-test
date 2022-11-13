/*

orm

@Author: Harsley Ekhorutomwen
@Date:  November 12, 2022

The base struct for application repository. All repository that extend from the parent struct will carry basic functionality.
	- create: Create a new object
	- update: Update an existing object
	- get: Retrieve an object by ID
	- delete: Delete an object by ID
*/

package ormwrapper

import (
	"context"
	"fmt"

	"github.com/go-play/utils"
	"gorm.io/gorm"
)

type OrmModel[E any] interface {
	ToModel() E
	FromModel(entity E) interface{}
}

func NewRepository[M OrmModel[E], E any](db *gorm.DB) *ORMRepository[M, E] {
	return &ORMRepository[M, E]{
		db: db,
	}
}

type ORMRepository[M OrmModel[E], E any] struct {
	db *gorm.DB
}

func (r *ORMRepository[M, E]) Create(ctx context.Context, data OrmModel[E]) error {

	model := data.ToModel()

	err := r.db.WithContext(ctx).Create(&model).Error

	if err != nil {
		return err
	}

	data = data.FromModel(model).(M)

	return nil

}

func (r *ORMRepository[M, E]) Update(ctx context.Context, id uint, data OrmModel[E]) error {

	model := data.ToModel()

	fmt.Println(model)

	single, err := r.FindByID(ctx, id)
	fmt.Println(single)

	if err != nil {
		utils.PanicException(err, "could not find item by id")
	}

	// value := structs.Map(&data)
	// fmt.Println(value)
	mo := single.ToModel()

	err = r.db.WithContext(ctx).Model(&mo).Where("id = ?", id).Updates(model).Error

	if err != nil {
		return err
	}

	data = data.FromModel(model).(M)

	return nil
}

func (r *ORMRepository[M, E]) FindByID(ctx context.Context, id uint, preload ...string) (M, error) {

	var preloadValue string
	var err error
	for i, item := range preload {
		if (i + 1) == len(preload) {
			preloadValue += item
		} else {
			preloadValue += item + ","
		}
	}

	var entity E
	var model M

	var length = len([]rune(preloadValue))

	if length > 0 {
		err = r.db.WithContext(ctx).Preload(preloadValue).First(&entity, id).Error
	} else {
		err = r.db.WithContext(ctx).First(&entity, id).Error
	}

	if err != nil {
		return *new(M), err
	}

	return model.FromModel(entity).(M), nil
}

func (r *ORMRepository[M, E]) Find(ctx context.Context, preload ...string) ([]M, error) {

	var preloadValue string
	var err error
	for i, item := range preload {
		if (i + 1) == len(preload) {
			preloadValue += item
		} else {
			preloadValue += item + ","
		}
	}

	var models M
	var entity []E

	var length = len([]rune(preloadValue))

	if length > 0 {
		err = r.db.WithContext(ctx).Preload(preloadValue).Find(&entity).Error
	} else {
		err = r.db.WithContext(ctx).Find(&entity).Error
	}

	if err != nil {
		return nil, err
	}

	result := make([]M, 0, len(entity))
	for _, row := range entity {
		result = append(result, models.FromModel(row).(M))
	}
	return result, nil
}
