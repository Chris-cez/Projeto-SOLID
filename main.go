package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type Person struct {
	ID   int    `csv:"id"`
	Name string `csv:"name"`
	Age  int    `csv:"age"`
}

func readPeople(filename string) ([]Person, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	people := []Person{}

	for {
		record, err := reader.Read()
		if err != nil {
			break
		}

		id, _ := strconv.Atoi(record[0])
		name := record[1]
		age, _ := strconv.Atoi(record[2])

		person := Person{
			ID:   id,
			Name: name,
			Age:  age,
		}

		people = append(people, person)
	}

	return people, nil
}

func createPerson(filename string, person Person) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	record := []string{strconv.Itoa(person.ID), person.Name, strconv.Itoa(person.Age)}
	err = writer.Write(record)
	if err != nil {
		return err
	}

	return nil
}

func readPerson(filename string, id int) (Person, error) {
	people, err := readPeople(filename)
	if err != nil {
		return Person{}, err
	}

	for _, person := range people {
		if person.ID == id {
			return person, nil
		}
	}

	return Person{}, fmt.Errorf("person with ID %d not found", id)
}

func updatePerson(filename string, person Person) error {
	people, err := readPeople(filename)
	if err != nil {
		return err
	}

	found := false
	for i, p := range people {
		if p.ID == person.ID {
			people[i] = person
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("person with ID %d not found", person.ID)
	}

	return writePeople(filename, people)
}

func deletePerson(filename string, id int) error {
	people, err := readPeople(filename)
	if err != nil {
		return err
	}

	newPeople := []Person{}
	for _, person := range people {
		if person.ID != id {
			newPeople = append(newPeople, person)
		}
	}

	return writePeople(filename, newPeople)
}

func writePeople(filename string, people []Person) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, person := range people {
		record := []string{strconv.Itoa(person.ID), person.Name, strconv.Itoa(person.Age)}
		err = writer.Write(record)
		if err != nil {
			return err
		}
	}

	return nil
}

func main() {
	people, err := readPeople("people.csv")
	if err != nil {
		panic(err)
	}

	fmt.Println(people)

	newPerson := Person{
		ID:   4,
		Name: "John Doe",
		Age:  30,
	}
	err = createPerson("people.csv", newPerson)
	if err != nil {
		panic(err)
	}

	createdPerson, err := readPerson("people.csv", 4)
	if err != nil {
		panic(err)
	}
	fmt.Println(createdPerson)

	updatedPerson := Person{
		ID:   4,
		Name: "Jane Doe",
		Age:  35,
	}
	err = updatePerson("people.csv", updatedPerson)
	if err != nil {
		panic(err)
	}
	updatedPerson, err = readPerson("people.csv", 4)
	if err != nil {
		panic(err)
	}
	fmt.Println(updatedPerson)
	err = deletePerson("people.csv", 4)
	if err != nil {
		panic(err)
	}
	people, err = readPeople("people.csv")
	if err != nil {
		panic(err)
	}
	fmt.Println(people)
}
