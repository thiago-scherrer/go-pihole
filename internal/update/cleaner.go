package update

import (
	"fmt"
	"log"
	"regexp"
)

// Clean function can be used for remove noise from blocklist
func Clean(domains []byte) ([]string, error) {
	//myRegex, _ := regexp.Compile(`(#)`)

	com, err := rmComent(domains)

	if err != nil {
		log.Print(err)
		return nil, err
	}

	spc, err := rmNonDomains(com)

	if err != nil {
		log.Print(err)
		return nil, err
	}
	fmt.Println(spc)
	return nil, nil

}

func rmComent(domains []byte) (string, error) {
	var altered string

	myRegex := regexp.MustCompile(`(?m)^#(.*)`)
	a1 := myRegex.ReplaceAllString(string(domains), "")

	myRegex = regexp.MustCompile(`(?m)#(.*)`)
	altered = myRegex.ReplaceAllString(string(a1), "")

	return altered, nil
}

func rmNonDomains(domains string) (string, error) {
	var altered string

	myRegex := regexp.MustCompile(`(?m)^repo(.*)`)
	altered = myRegex.ReplaceAllString(domains, "")

	myRegex, _ = regexp.Compile(`(0.0.0.0 )`)
	altered = myRegex.ReplaceAllString(altered, "")

	myRegex, _ = regexp.Compile(`(127.0.0.1 )`)
	altered = myRegex.ReplaceAllString(altered, "")

	fmt.Println(altered)
	return altered, nil

}
