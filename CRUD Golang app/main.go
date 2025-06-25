package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"strings"

	"bufio"
	"os"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	ID string
	Name string
	Email string
	Age int
}




// Create a Table
func createTable(db *sql.DB) error {
	codeSQL := `
	CREATE TABLE IF NOT EXISTS users(
		id TEXT PRIMARY KEY,
		name TEXT,
		email TEXT,
		age INTEGER	
	);
	`

	_, err := db.Exec(codeSQL)

	return err
}


// Create User
func createUser(name string, email string, age int) *User {
	return &User {
		ID: uuid.New().String(),
		Name: name,
		Email: email,
		Age: age,
	}
}


// Insert User
func insertUser(db *sql.DB, user *User) error {
	stmt, err := db.Prepare("INSERT INTO users(id, name, email, age) VALUES($1, $2, $3, $4)")
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(user.ID, user.Name, user.Email, user.Age)
	if err != nil {
		return err
	}
	return nil
}


// Read User
func readUser(db *sql.DB) ([]User, error) {
	rows, err := db.Query("SELECT id, name, email, age FROM users")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []User

	for rows.Next() {
		var u User
		err = rows.Scan(&u.ID, &u.Name, &u.Email, &u.Age)
		if err != nil {
			return nil, err
		}
		users =append(users, u)
	}
	return users, nil
}


// Update User
func updateUser(db *sql.DB, user *User) error {
	stmt, err := db.Prepare("UPDATE users SET name = ?, email = ?, age = ? WHERE id = ?")
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(user.Name, user.Email, user.Age, user.ID)
	if err != nil {
		return err
	}
	return nil
}


// Delete User
func deleteUser(db *sql.DB, name string) error {
	stmt, err := db.Prepare("DELETE FROM users WHERE name = ?")
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(name)
	if err != nil {
		return err
	}
	return nil
}




func main() {

	// Create DataBase
	db, err := sql.Open("sqlite3", "DB/users.db")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	err = createTable(db)
	if err != nil {
		log.Fatal(err)
	}

	for {
		fmt.Println("------- Cadastro de Usuários -------")
		fmt.Println("(1) Criar Usuário")
		fmt.Println("(2) Ver Usuários Cadastrados")
		fmt.Println("(3) Atualizar Usuário")
		fmt.Println("(4) Remover Usuário")
		fmt.Println("(5) Sair")


		var input int

		fmt.Print("Escolha a opção: ")
		fmt.Scanln(&input)

		switch input {
		case 1:
			var (
				nome string
				email string
				idadeStr string
			)
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Digite o Nome Completo: ")
			nome, _ = reader.ReadString('\n')
			nome = strings.TrimSpace(nome)

			fmt.Print("Digite o email: ")
			email, _ = reader.ReadString('\n')
			email = strings.TrimSpace(email)

			fmt.Print("Digite a Idade: ")
			idadeStr, _ = reader.ReadString('\n')
			idadeStr = strings.TrimSpace(idadeStr)
			idade, err := strconv.Atoi(idadeStr) 
			if err != nil {
				fmt.Println("Houve um erro! Idade inválida!")
			}


			NewUser := createUser(nome, email, idade)
			insertUser(db, NewUser)
			fmt.Println("Usuário criado e cadastrado com sucesso!")

		case 2:
			users, err := readUser(db)
			if err != nil {
				log.Fatal(err)
			}
			var contUsers int = 0
			for _, user := range users {
				contUsers++
				fmt.Println("Usuário:", user)
			}
			fmt.Println("O total de:", contUsers,"Cadastrados")

		case 3:
			var userInputName string
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Digite o nome completo do usuário para busca: ")
			userInputName, _ = reader.ReadString('\n')
			userInputName = strings.TrimSpace(userInputName)

			rows, err := db.Query("SELECT id, name FROM users")
			if err != nil {
				log.Fatal(err)
			}
			defer rows.Close()

			var foundID, foundName string
			userFound := false
			for rows.Next() {
				var id, name string
				err := rows.Scan(&id, &name)
				if err != nil {
					log.Fatal(err)
				}
				if userInputName == name {
					foundID = id
					foundName = name
					userFound = true
					break
				}
			}

			if !userFound {
				fmt.Println("Nenhum usuário encontrado com esse nome!")
				break
			}

			fmt.Printf("\n👤 Usuário encontrado!\nNome atual: %s (ID: %s)\n", foundName, foundID)

			// Solicita as novas informações
			fmt.Print("Novo nome (deixe em branco para não alterar): ")
			newName, _ := reader.ReadString('\n')
			newName = strings.TrimSpace(newName)

			fmt.Print("Novo e-mail (deixe em branco para não alterar): ")
			newEmail, _ := reader.ReadString('\n')
			newEmail = strings.TrimSpace(newEmail)

			fmt.Print("Nova idade (deixe em branco para não alterar): ")
			newAgeStr, _ := reader.ReadString('\n')
			newAgeStr = strings.TrimSpace(newAgeStr)

			// Lê dados atualizados para não perder nada
			currentUser := User{ID: foundID, Name: foundName}
			currentDataRow := db.QueryRow("SELECT email, age FROM users WHERE id = ?", foundID)
			err = currentDataRow.Scan(&currentUser.Email, &currentUser.Age)
			if err != nil {
				log.Fatal(err)
			}

			// Aplica as mudanças
			if newName != "" {
				currentUser.Name = newName
			}
			if newEmail != "" {
				currentUser.Email = newEmail
			}
			if newAgeStr != "" {
				ageInt, err := strconv.Atoi(newAgeStr)
				if err != nil {
					fmt.Println("Idade inválida! Nenhuma alteração feita para a idade.")
				} else {
					currentUser.Age = ageInt
				}
			}

			// Salva atualizações no BD
			err = updateUser(db, &currentUser)
			if err != nil {
				fmt.Println("Erro ao atualizar o usuário:", err)
			} else {
				fmt.Println("Usuário atualizado com sucesso!")
			}

		case 4:
			var nameToDelete string

			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Digite o nome para deletar: ")
			nameToDelete, _ = reader.ReadString('\n')
			nameToDelete = strings.TrimSpace(nameToDelete)


			rows, err := db.Query("SELECT name FROM users")
			if err != nil {
				log.Fatal(err)
			}
			defer rows.Close()

			userFound := false
			for rows.Next() {
				var name string
				err := rows.Scan(&name)
				if err != nil {
					log.Fatal(err)
				}
				if nameToDelete == name {
					userFound = true
					break
				}
			}
			if !userFound {
				fmt.Printf("Usuário %s Não encontrado! \n", nameToDelete)
			}

			fmt.Println("Usuário encontrado!")
			var deleteQuestion string

			fmt.Printf("Deletar o Usuário: %s (s) (n)? \n", nameToDelete)
			deleteQuestion, _ = reader.ReadString('\n')
			deleteQuestion = strings.TrimSpace(deleteQuestion)

			if deleteQuestion == "s" {
				deleteUser(db, nameToDelete)
			} else if deleteQuestion == "n" {
				fmt.Println("O usuário não foi deletado!")
			} else {
				fmt.Println("Opção inválida!")
			}

		case 5: 
		fmt.Println("Você saiu do programa!")
		return

		default:
			fmt.Println("Opção inválida!")
		}

	}

}




