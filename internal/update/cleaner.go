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
	myRegex := regexp.MustCompile(`(?m)^#(.*)`)
	altered := myRegex.ReplaceAllString(string(domains), "")

	myRegex = regexp.MustCompile(`(?m)#(.*)`)
	altered = myRegex.ReplaceAllString(string(altered), "")

	return altered, nil
}

func rmNonDomains(domains string) (string, error) {
	myRegex :=
		regexp.MustCompile(`(?m)^repo(.*)|(0.0.0.0)|(127.0.0.1)|(localhost)|(
  )|(\n\n)|(\t)|((?m)@)|( )|(ff02::)|(::1)`)
	altered := myRegex.ReplaceAllString(domains, "")

	myRegex =
		regexp.MustCompile(`(<br>)|(<BR>)`)
	altered = myRegex.ReplaceAllString(altered, "\n")

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
