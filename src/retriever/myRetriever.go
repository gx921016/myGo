package gaoxiang

type Retriever struct {
	Contents string
	D        string
}

func (r Retriever) Get(url string) string {
	return r.D
}

func main() {

}
