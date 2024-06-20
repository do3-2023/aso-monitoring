package handlers

import (
	"apimonitoring/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	spinhttp "github.com/fermyon/spin/sdk/go/v2/http"
)

type Person struct {
	Id          int64
	LastName    string
	PhoneNumber string
}

type CreatePersonDto struct {
	LastName    string
	PhoneNumber string
}

func FindAllPerson(w http.ResponseWriter, _ *http.Request, ps spinhttp.Params) {
	db, err := utils.OpenPostgres()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM person;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var persons []*Person
	for rows.Next() {
		var person Person
		if err := rows.Scan(&person.Id, &person.LastName, &person.PhoneNumber); err != nil {
			fmt.Println(err)
		}

		persons = append(persons, &person)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(persons)
}

func FindOnePerson(w http.ResponseWriter, _ *http.Request, ps spinhttp.Params) {
	db, err := utils.OpenPostgres()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	defer db.Close()

	id, err := strconv.ParseInt(ps.ByName("id"), 10, 32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var person Person
	err = db.QueryRow("SELECT * FROM person WHERE id = $1;", int32(id)).Scan(&person.Id, &person.LastName, &person.PhoneNumber)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(person)
}

func CreatePerson(w http.ResponseWriter, r *http.Request, ps spinhttp.Params) {
	db, err := utils.OpenPostgres()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	defer db.Close()

	decoder := json.NewDecoder(r.Body)
	var dto CreatePersonDto
	err = decoder.Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	var person Person
	err = db.QueryRow("INSERT INTO person (last_name, phone_number) VALUES ($1, $2) RETURNING *;", dto.LastName, dto.PhoneNumber).Scan(&person.Id, &person.LastName, &person.PhoneNumber)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(person)
}

func DeletePerson(w http.ResponseWriter, _ *http.Request, ps spinhttp.Params) {
	db, err := utils.OpenPostgres()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	defer db.Close()

	id, err := strconv.ParseInt(ps.ByName("id"), 10, 32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = db.Exec("DELETE FROM person WHERE id = $1;", int32(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
