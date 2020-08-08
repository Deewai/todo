package storage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDB(t *testing.T){
	db := DB{store:make(map[string]interface{})}
	assert.IsType(t, &DB{}, NewDB())
	assert.True(t, assert.ObjectsAreEqual(&db, NewDB()))
}

func TestAddItem(t *testing.T){
	db := DB{store:make(map[string]interface{})}
	item := map[string]interface{}{"key1": "value 1","key2": "value 2"}
	pk := "1"
	db.AddItem(pk, item)
	assert.EqualValues(t, item, db.store[pk])
}

func TestGetItem(t *testing.T){
	db := DB{store:make(map[string]interface{})}
	item := map[string]interface{}{"key1": "value 1","key2": "value 2"}
	pk := "1"
	db.store[pk] = item
	got, err := db.GetItem(pk)
	assert.NoError(t, err)
	assert.EqualValues(t, item, got)
}

func TestGetItemNotFound(t *testing.T){
	db := DB{store:make(map[string]interface{})}
	item, err := db.GetItem("1")
	assert.Nil(t, item)
	assert.Error(t, err)
	assert.Equal(t, ErrNotFound, err)
}

func TestDeleteItem(t *testing.T){
	db := DB{store:make(map[string]interface{})}
	item := map[string]interface{}{"key1": "value 1","key2": "value 2"}
	pk := "1"
	db.store[pk] = item
	err := db.DeleteItem(pk)
	assert.NoError(t, err)
	assert.Zero(t, len(db.store))
}

func TestDeleteItemNotFound(t *testing.T){
	db := DB{store:make(map[string]interface{})}
	err := db.DeleteItem("1")
	assert.Error(t, err)
	assert.Equal(t, ErrNotFound, err)
}

func TestGetItems(t *testing.T){
	db := DB{store:make(map[string]interface{})}
	item := map[string]interface{}{"key1": "value 1","key2": "value 2"}
	pk := "1"
	db.store[pk] = item
	got, err := db.GetItems()
	assert.NoError(t, err)
	assert.GreaterOrEqual(t,len(got),1)
	assert.Equal(t, item, got[0])
}


