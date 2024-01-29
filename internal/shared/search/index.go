package search

type SearchIndexes struct {
	PropertyIndex string
}

func GetSearchIndexes() SearchIndexes {
	return SearchIndexes{
		PropertyIndex: "property_index",
	}
}
