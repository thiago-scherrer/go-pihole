package update

import (
	"errors"
	"log"
	"os"
	"regexp"
)

var (
	errNeedOutput = errors.New("[error]: need OUTPUT env to save blocklist")
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
	writeDom(spc)
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

	myRegex, _ = regexp.Compile(`( localhost)`)
	altered = myRegex.ReplaceAllString(altered, "")

	return altered, nil
}

func writeDom(domains string) ([]string, error) {
	output := os.Getenv("OUTPUT")

	if len(output) <= 1 {
		log.Println(errNeedOutput)
		return nil, errNeedOutput
	}

	f, err := os.OpenFile(output, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0755)
	if err != nil {
		log.Println("[error] failed on open file", err)
		return nil, err
	}

	defer f.Close()
	_, err = f.WriteString(domains)
	if err != nil {
		log.Println("[error] failed on write file", err)
		return nil, err
	}
	return nil, nil
}
