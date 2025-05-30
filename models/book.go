package models

type Book struct {
	Title  string  `json:"title" firestore:"title"`
	Author string  `json:"author" firestore:"author"`
	Year   int     `json:"year" firestore:"year"`
	Genre  string  `json:"genre" firestore:"genre"`
	Pages  int     `json:"pages" firestore:"pages"`
	Rating float64 `json:"rating" firestore:"rating"`
	Review string  `json:"review" firestore:"review"`
}
