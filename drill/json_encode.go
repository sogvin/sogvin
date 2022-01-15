// Encode struct to json
package drill

import (
	"encoding/json"
	"log"
	"os"
)

func init() {
	car := struct {
		Model string
		Year  int
	}{
		Model: "audi",
		Year:  2021,
	}

	enc := json.NewEncoder(os.Stdout)
	if err := enc.Encode(car); err != nil {
		log.Fatal(err)
	}
}
