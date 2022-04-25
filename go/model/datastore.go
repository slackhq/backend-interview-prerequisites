package model

// Datastore is the main entrypoint into the data storage models
type Datastore struct {
	TestTableStore *TestTableStore
}

// InitDatastore bootstraps the datastore
func InitDatastore(dbPath, schemaPath string) *Datastore {
	db := initDatabase(dbPath, schemaPath)

	datastore := Datastore{
		TestTableStore: initTestStore(db),
	}

	return &datastore
}
