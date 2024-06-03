package constants

type SEARCH_BY string

const (
	DEV_ENV string = "development"

	SEARCH_BY_TITLE  SEARCH_BY = "title"
	SEARCH_BY_ISBN   SEARCH_BY = "isbn"
	SEARCH_BY_AUTHOR SEARCH_BY = "author"
)
