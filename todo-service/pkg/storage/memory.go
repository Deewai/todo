package storage

// DB represents an in-memory database
type DB struct {
	store map[string]interface{}
}

// NewDB returns a new instance of the memory db
func NewDB() *DB {
	return &DB{store: make(map[string]interface{})}
}

// AddItem adds a new item to the memory database using the primary key
func (db *DB) AddItem(pk string, item interface{}) error {
	db.store[pk] = item
	return nil
}

// GetItem get an item from the memory database using the primary key
func (db *DB) GetItem(pk string) (interface{}, error) {
	if item, ok := db.store[pk]; ok {
		return item, nil
	}
	return nil, ErrNotFound
}

// DeleteItem deletes an item from the memory database using the primary key
func (db *DB) DeleteItem(pk string) error {
	if _, ok := db.store[pk]; ok {
		delete(db.store, pk)
		return nil
	}
	return ErrNotFound
}

// GetItems gets all the items in the memory database
func (db *DB) GetItems() ([]interface{}, error) {
	items := []interface{}{}
	for _, item := range db.store {
		items = append(items, item)
	}
	return items, nil
}
