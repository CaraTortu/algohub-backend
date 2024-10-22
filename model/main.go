package model

// Returns models that will be reset in the database when resetting the DB
func GetModels() []interface{} {
	return []interface{}{
		&Course{}, // This will delete Course, Chapters, Sections, and Posts because of the cascading
	}
}
