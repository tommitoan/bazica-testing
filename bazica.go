package bazica

import (
	"encoding/json"
	"github.com/tommitoan/bazica/internal/datasvc"
	toerr "github.com/tommitoan/bazica/internal/toerr"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

type BazicaSvc struct {
	DataSvc *datasvc.DataSvc
}

func (bzc *BazicaSvc) GetSolarTermsByYear(prefix, year string) (*datasvc.SolarTerms, error) {
	// Handle year from 1900 -> 2100 only
	i, err := strconv.Atoi(year)
	if err != nil {
		panic(err)
		return nil, toerr.NewValidationError(http.StatusInternalServerError, "Something wrong")
	}
	if i < 1900 || i > 2100 {
		return nil, toerr.NewValidationError(http.StatusBadRequest, "Year not found")
	}

	// Open file
	fileToRead := prefix + year + ".json"
	file, err := os.Open(fileToRead)
	if err != nil {
		return nil, toerr.NewValidationError(http.StatusInternalServerError, "File not found")
		log.Fatal(err)
	}
	defer file.Close()

	// Read as byte array
	byteValue, err := io.ReadAll(file)
	if err != nil {
		return nil, toerr.NewValidationError(http.StatusInternalServerError, "Can't read file")
	}
	var solarTerm datasvc.SolarTerms
	json.Unmarshal(byteValue, &solarTerm)

	return &solarTerm, nil
}