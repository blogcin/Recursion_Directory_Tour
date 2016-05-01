package Tour

import (
	"io/ioutil"
	"log"
	"fmt"
)

type Tour struct {
	Pathspliter string
}

func (tour *Tour) Explore(dirName string) {
	go func() {
		files, err := ioutil.ReadDir(dirName)

		if err != nil {
			log.Print(err)
		}

		for _, f := range files {
			if(f.IsDir()) {
				if(dirName != tour.Pathspliter) {
					tour.Explore(dirName + tour.Pathspliter + f.Name())
				} else {
					tour.Explore(dirName + f.Name())
				}
			} else {
				if(dirName != tour.Pathspliter) {
					fmt.Println(dirName + tour.Pathspliter + f.Name())
				} else {
					fmt.Println(dirName + f.Name())
				}
			}
		}
	}()
}

