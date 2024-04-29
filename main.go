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

// Interface para abstrair a operação de leitura de arquivos
type Reader interface {
	Read() ([]Person, error)
}

// Interface para abstrair a operação de escrita de arquivos
type Writer interface {
	Write(people []Person) error
}

// Classe para ler dados de um arquivo CSV
type CSVReader struct {
	filename string
}

// Implementa a interface Reader
func (r *CSVReader) Read() ([]Person, error) {
	file, err := os.Open(r.filename)
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

// Classe para escrever dados em um arquivo CSV
type CSVWriter struct {
	filename string
}

// Implementa a interface Writer
func (w *CSVWriter) Write(people []Person) error {
	file, err := os.Create(w.filename)
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

// Função para criar um novo registro
func createPerson(w Writer, person Person) error {
	people, err := w.Read()
	if err != nil {
		return err
	}

	people = append(people, person)

	return w.Write(people)
}

// Função para ler um registro específico
func readPerson(r Reader, id int) (Person, error) {
	people, err := r.Read()
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

// Função para atualizar um registro
func updatePerson(w Writer, person Person) error {
	people, err := w.Read()
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

	return w.Write(people)
}

// Função para deletar um registro
func deletePerson(w Writer, id int) error {
	people, err := w.Read()
	if err != nil {
		return err
	}

	newPeople := []Person{}
	for _, person := range people {
		if person.ID != id {
			newPeople = append(newPeople, person)
		}
	}

	return w.Write(newPeople)
}

func main() {
	// Criar um leitor e um escritor
	reader := &CSVReader{filename: "people.csv"}
	writer := &CSVWriter{filename: "people.csv"}

	// Ler os dados do arquivo CSV
	people, err := reader.Read()
	if err != nil {
		panic(err)
	}

	// Imprimir os dados
	fmt.Println(people)

	// Criar um novo registro
	newPerson := Person{
		ID:   4,
		Name: "John Doe",
		Age:  30,
	}
	err = createPerson(writer, newPerson)
	if err != nil {
		panic(err)
	}

	// Ler o registro criado
	createdPerson, err := readPerson(reader, 4)
	if err != nil {
		panic(err)
	}
	fmt.Println(createdPerson)

	// Atualizar o registro
	updatedPerson := Person{
		ID:   4,
		Name: "Jane Doe",
		Age:  35,
	}
	err = updatePerson(writer, updatedPerson)
	if err != nil {
		panic(err)
	}

	// Ler o registro atualizado
	updatedPerson, err = readPerson(reader, 4)
	if err != nil {
		panic(err)
	}
	fmt.Println(updatedPerson)

	// Deletar o registro
	err = deletePerson(writer, 4)
	if err != nil {
		panic(err)
	}

	// Ler os dados novamente
	people, err = reader.Read()
	if err != nil {
		panic(err)
	}

	// Imprimir os dados
	fmt.Println(people)
}
