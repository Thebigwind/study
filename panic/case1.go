package panic

import "log"

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("recover: %v", err)
		}
	}()

	panic("EDDYCJY.")
}
