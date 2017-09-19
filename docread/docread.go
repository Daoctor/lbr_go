package docread

import (
	"github.com/sajari/docconv"
	"log"
)

func Read(path string) string {
	res, err := docconv.ConvertPath(path)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(res)
	return res.Body
}

