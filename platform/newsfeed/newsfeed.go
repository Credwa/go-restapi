package newsfeed

// Getter interface
type Getter interface {
	GetAll() []Item
}

// Adder interface
type Adder interface {
	Add(item Item)
}

type Messy struct {
	Title string
}

// Item struct
type Item struct {
	Title string `json:"title"`
	Post  string `json:"post"`
}

// Repo struct
type Repo struct {
	Items []Item
}

// New function
func New() *Repo {
	return &Repo{
		Items: []Item{},
	}
}

// Add function
func (r *Repo) Add(item Item) {
	r.Items = append(r.Items, item)
}

// GetAll function
func (r *Repo) GetAll() []Item {
	return r.Items
}
