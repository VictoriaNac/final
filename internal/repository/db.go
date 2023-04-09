package repository

import (
	"encoding/csv"
	"os"
	"strconv"
)


const (
	nameIdx = iota
	regionIdx
	districtIdx
	populationIdx
	foundationIdx
)


const (
	errNotFoundId = "ERROR: CITY WITH THIS ID WAS NOT FOUND"
)


type DataBase struct {
	records map[int][]string
	lastID  int
}


func NewDataBase(filePath string) (*DataBase, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	cities, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}

	cityList := new(DataBase)
	cityList.records = make(map[int][]string)

	if len(cities) == 0 {
		return cityList, nil
	}

	for _, city := range cities {
		id, _ := strconv.Atoi(city[0])
		if err != nil {
			return nil, err
		}
		cityList.records[id] = make([]string, 5)
		copy(cityList.records[id], city[1:])
	}

	cityList.lastID = 0
	for cityID := range cityList.records {
		if cityList.lastID < cityID {
			cityList.lastID = cityID
		}
	}
	return cityList, nil
}


func (db *DataBase) SaveCSV(filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	var data [][]string
	for id, description := range db.records {
		var cityLine []string
		cityLine = append(cityLine, strconv.Itoa(id))
		cityLine = append(cityLine, description...)
		data = append(data, cityLine)
	}

	writer := csv.NewWriter(file)
	err = writer.WriteAll(data)
	if err != nil {
		return err
	}
	return nil
}