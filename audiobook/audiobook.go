package audiobook

type AudioBook struct {
	BookName  string
	Author    string
	Reader    string
	AudioUrl  string
	StorePath string //set the value after download -> BookName-Author/Reader/
}
