package movies

// TODO: Move close to handler
type UpdateMovieData struct {
	Title   *string       `json:"title"`
	Year    *int          `json:"year"`
	Runtime *MovieRuntime `json:"runtime"`
	Genres  []string      `json:"genres"`
}
