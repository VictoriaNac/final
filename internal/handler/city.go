package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"unicode"

	cities "github.com/Kit0b0y/SkillboxHomeWork/NewSkillbox/Interim_certification"
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
)


func (h *Handler) getFull(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		newMessageResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	log.Info().Msg(fmt.Sprintf("GET: Full info %v", id))

	city, err := h.services.City.GetFull(id)
	if err != nil {
		newMessageResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	message, err := json.Marshal(city)
	if err != nil {
		newMessageResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	log.Info().Msg(fmt.Sprintf("Sending: %v", string(message)))
	w.WriteHeader(http.StatusOK)
	w.Write(message)
}


func (h *Handler) create(w http.ResponseWriter, r *http.Request) {
	var city cities.CityRequest
	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		newMessageResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	log.Info().Msg(fmt.Sprintf("POST: New city %v", string(content)))

	err = json.Unmarshal(content, &city)
	if err != nil {
		newMessageResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	city.Name = toUpperFirst(city.Name)
	city.Region = toUpperFirst(city.Region)
	city.District = toUpperFirst(city.District)

	id, err := h.services.City.Create(city)
	if err != nil {
		newMessageResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	message := fmt.Sprintf("The city of %v was created with the id %v", city.Name, id)
	newMessageResponse(w, http.StatusCreated, message)
}


func (h *Handler) delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		newMessageResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	log.Info().Msg(fmt.Sprintf("DELETE: Delete city %v", id))

	err = h.services.City.Delete(id)
	if err != nil {
		newMessageResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	message := fmt.Sprintf("City with identifier %v deleted", id)
	newMessageResponse(w, http.StatusOK, message)
}


func (h *Handler) setPopulation(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		newMessageResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	log.Info().Msg(fmt.Sprintf("PUT: Update population city %v", id))

	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		newMessageResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	var population cities.SetPopulationRequest
	err = json.Unmarshal(content, &population)
	if err != nil {
		newMessageResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.SetPopulation(id, population.Population)
	if err != nil {
		newMessageResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	newMessageResponse(w, http.StatusOK, "Change was successful")
}


func (h *Handler) getFromRegion(w http.ResponseWriter, r *http.Request) {
	region := chi.URLParam(r, "region")
	region = toUpperFirst(region)

	log.Info().Msg(fmt.Sprintf("GET: Сities by region %v", region))

	cityNames, err := h.services.GetFromRegion(region)
	if err != nil {
		newMessageResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	newListResponse(w, http.StatusOK, cityNames)
}


func (h *Handler) getFromDistrict(w http.ResponseWriter, r *http.Request) {
	district := chi.URLParam(r, "district")
	district = toUpperFirst(district)

	log.Info().Msg(fmt.Sprintf("GET: Сities by district %v", district))

	cityNames, err := h.services.GetFromDistrict(district)
	if err != nil {
		newMessageResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	newListResponse(w, http.StatusOK, cityNames)
}


func (h *Handler) getFromPopulation(w http.ResponseWriter, r *http.Request) {
	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		newMessageResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	log.Info().Msg(fmt.Sprintf("GET: Сities by population range %v", string(content)))

	var populationRange cities.RangeRequest
	err = json.Unmarshal(content, &populationRange)
	if err != nil {
		newMessageResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	cityNames, err := h.services.GetFromPopulation(populationRange.Start, populationRange.End)
	if err != nil {
		newMessageResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	newListResponse(w, http.StatusOK, cityNames)
}


func (h *Handler) getFromFoundation(w http.ResponseWriter, r *http.Request) {
	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		newMessageResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	log.Info().Msg(fmt.Sprintf("GET: Сities by foundation range %v", string(content)))

	var foundationRange cities.RangeRequest
	err = json.Unmarshal(content, &foundationRange)
	if err != nil {
		newMessageResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	cityNames, err := h.services.GetFromFoundation(foundationRange.Start, foundationRange.End)
	if err != nil {
		newMessageResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	newListResponse(w, http.StatusOK, cityNames)
}


func toUpperFirst(text string) string {
	textRune := []rune(text)
	for i := range textRune {
		textRune[i] = unicode.ToLower(textRune[i])
	}
	return string(unicode.ToUpper(textRune[0])) + string(textRune[1:])
}