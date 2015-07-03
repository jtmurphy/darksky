/*
	Package darksky implements a simple function that retreives 
	weather data at a given location.

	Use of this package require that you have access to a dark sky API
	key. You can obtain an API key by creating an account at
	https://developer.forecast.io/


	An example usage:

		import (
			"fmt"
			"github.com/jtmurphy/darksky"
		)

		const KEY = "abcdef0123456789abcdef0123456789"

		func main() {
			latitude := 10.0
			longitude := 10.0
			fc, err := darksky.Get(KEY, latitude, longitude)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Printf("%#v\n", fc)
		}
*/
package darksky
