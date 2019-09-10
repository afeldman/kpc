package kpc

//Repository is a structure to represend the developmentfiles
//the basic repository protocol is git
//the tag is the hash in the git tree.
type Repository struct {
	URL string `json:"url"`
	Tag string `json:"version"`
}

func InitRepository() *Repository {
	return &Repository{URL: "", Tag: ""}
}

func (repo *Repository) SetTag(tag string) {
	repo.Tag = tag
}

func (repo *Repository) GetTag() *string {
	return &(repo.Tag)
}

func (repo *Repository) SetURL(address string) {
	repo.URL = address
}

func (repo *Repository) GetURL() *string {
	return &(repo.URL)
}
