package main

import (
	"errors"
	"fmt"
)

type RowData struct {
	RankWebsite int
	Domain      string
	TLD         string
	IDN_TLD     string
	Valid       bool
	RefIPs      int
}

func GetTLD(domain string) (TLD string, IDN_TLD string) {
	var ListIDN_TLD = map[string]string{
		".com": ".co.id",
		".org": ".org.id",
		".gov": ".go.id",
	}

	for i := len(domain) - 1; i >= 0; i-- {
		if domain[i] == '.' {
			TLD = domain[i:]
			break
		}
	}

	if _, ok := ListIDN_TLD[TLD]; ok {
		return TLD, ListIDN_TLD[TLD]
	} else {
		return TLD, TLD
	}
}

func ProcessGetTLD(website RowData, ch chan RowData, chErr chan error) {
	if website.Domain == "" {
        chErr <- errors.New("domain name is empty")
        return
    } else if !website.Valid {
        chErr <- errors.New("domain not valid")
        return
    } else if website.RefIPs == -1 {
        chErr <- errors.New("domain RefIPs not valid")
        return
    }
	
	// TODO: replace this
	TLD, IDN_TLD := GetTLD(website.Domain)
	website.TLD = TLD
    website.IDN_TLD = IDN_TLD

    ch <- website
}

// Gunakan variable ini sebagai goroutine di fungsi FilterAndGetDomain
var FuncProcessGetTLD = ProcessGetTLD

func FilterAndFillData(TLD string, data []RowData) ([]RowData, error) {
	ch := make(chan RowData, len(data))
	errCh := make(chan error)

	for _, website := range data {
		go FuncProcessGetTLD(website, ch, errCh)
	}

	proccessedData := make([]RowData, 0)
	errors := make([]error, 0)

	for i := 0; i < len(data); i++ {
		select {
		case website := <-ch:
			proccessedData = append(proccessedData, website)
		case err := <-errCh:
			errors = append(errors, err)
		}
	}

	if len(errors) > 0 {
		return nil, errors[0]
	}

	filteredData := make([]RowData, 0)
	for _, website := range proccessedData {
		if website.TLD == TLD {
			filteredData = append(filteredData, website)
		}
	}
	return filteredData, nil // TODO: replace this
}

// gunakan untuk melakukan debugging
func main() {
	rows, err := FilterAndFillData(".com", []RowData{
		{1, "google.com", "", "", true, 100},
		{2, "facebook.com", "", "", true, 100},
		{3, "golang.org", "", "", true, 100},
	})

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(rows)
}
