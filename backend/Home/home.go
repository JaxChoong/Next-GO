package home

type Book struct {
	Id          string
	Title       string
	Author      string
	Year        int
	Genre       string
	Description string
	CoverImage  string
	Rating      float64
	Pages       int
	ISBN        string
	Publisher   string
	Tags        []string
}

var books = []Book{
	{
		Id:          "bk_1",
		Title:       "The Last Light of Avalon",
		Author:      "Marina K. Rhodes",
		Year:        2020,
		Genre:       "Speculative Fantasy",
		Description: "A lyrical quest novel about a lost isle, a reluctant guardian, and the small acts that keep hope alive.",
		CoverImage:  "https://images.unsplash.com/photo-1524995997946-a1c2e315a42f?q=80&w=600&auto=format&fit=crop",
		Rating:      4.4,
		Pages:       384,
		ISBN:        "978-1-904823-12-3",
		Publisher:   "Northwind Press",
		Tags:        []string{"quest", "myth", "female protagonist"},
	},
	{
		Id:          "bk_2",
		Title:       "Algorithms of Us",
		Author:      "Ravi S. Menon",
		Year:        2022,
		Genre:       "Technology / Essays",
		Description: "Short, sharp essays exploring how algorithms reshape everyday decisions, privacy, and public policy.",
		CoverImage:  "https://images.unsplash.com/photo-1515879218367-8466d910aaa4?q=80&w=600&auto=format&fit=crop",
		Rating:      4.0,
		Pages:       240,
		ISBN:        "978-0-7636-9874-5",
		Publisher:   "Cityline Books",
		Tags:        []string{"technology", "essays", "data"},
	},
	{
		Id:          "bk_3",
		Title:       "A Small Map to the Sea",
		Author:      "Lena Torvald",
		Year:        2018,
		Genre:       "Literary Fiction",
		Description: "An intimate family story spanning three summers on a northern shore, balancing memory, regret, and reconciliation.",
		CoverImage:  "https://images.unsplash.com/photo-1507525428034-b723cf961d3e?q=80&w=600&auto=format&fit=crop",
		Rating:      4.7,
		Pages:       312,
		ISBN:        "978-1-4028-9463-1",
		Publisher:   "Harbor & Co.",
		Tags:        []string{"family", "memory", "coastal"},
	},
	{
		Id:          "bk_4",
		Title:       "The Quantum Gardener",
		Author:      "Diego Alvarez",
		Year:        2024,
		Genre:       "Sci-Fi",
		Description: "Near-future sci-fi about a botanist who engineers plants that can compute — and the ethical choices that follow.",
		CoverImage:  "https://images.unsplash.com/photo-1517836357463-d25dfeac3438?q=80&w=600&auto=format&fit=crop",
		Rating:      4.1,
		Pages:       288,
		ISBN:        "978-1-9787-3320-7",
		Publisher:   "Pioneer Press",
		Tags:        []string{"sci-fi", "bioengineering", "ethics"},
	},
}

func Home() []Book {
	message := books
	return message
}
