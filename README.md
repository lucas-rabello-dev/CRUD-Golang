# CRUD em Golang usando SQLite (CLI)

Um exemplo simples de aplicação CLI para realizar operações de **CRUD** (Create, Read, Update, Delete) em **Go** (Golang) usando SQLite.
<p align="center">
  <img src="https://upload.wikimedia.org/wikipedia/commons/thumb/0/05/Go_Logo_Blue.svg/1200px-Go_Logo_Blue.svg.png" alt="golang_imagem" width="200">
</p>

---

## 🧱 Arquitetura de Camadas

O projeto segue a **Arquitetura em Camadas (Layered Architecture)**, separando responsabilidades em diferentes pacotes para garantir organização, manutenibilidade e escalabilidade:

- **`main.go`**: ponto de entrada da aplicação, responsável por inicializar o banco e orquestrar o menu principal.
- **`db/`**: camada de infraestrutura. Responsável por abrir conexões e criar as tabelas no banco de dados SQLite.
- **`models/`**: contém as structs (como `User`) que representam os dados da aplicação.
- **`repository/`**: camada de persistência. Contém as funções que executam consultas SQL no banco de dados.
- **`service/`**: lógica de negócio. Define regras e construção dos dados que serão enviados ao repositório.
- **`utils/`**: utilitários para entrada de dados e interações com o usuário via terminal (CLI).

---

## Funcionalidades
1. Criar usuários  
2. Listar usuários cadastrados no `.db`  
3. Atualizar dados de usuários  
4.  Excluir usuários  

---

## 🗄️ Banco de Dados
- O arquivo `.db` é gerado automaticamente ao executar o programa.  
- Os dados poderão ser inseridos e manipulados conforme as opções do menu.

---

## 📦 Pacotes Utilizados

### Nativos do Go 1.24.2
- `database/sql`
- `fmt`
- `log`
- `strconv`
- `strings`
- `bufio`
- `os`

### Externos
- [`github.com/mattn/go-sqlite3`](https://github.com/mattn/go-sqlite3)
  
  Instalação:
  ```bash
  go get github.com/mattn/go-sqlite3
  ```

- [`github.com/google/uuid`](https://github.com/google/uuid)
  
  Instalação:
  ```bash
  go get github.com/google/uuid
  ```
## 💡​Dica
Ao inserir o nome do usuário para consulta, atualização ou remoção, utilize exatamente o mesmo nome salvo no banco de dados. Diferenças de maiúsculas, acentuação ou espaços podem impedir que o usuário seja localizado.
## Redes sociais:
<hr>

[![My Skills](https://skillicons.dev/icons?i=linkedin)](https://www.linkedin.com/in/lucas-rabello-42b23a339/) 
[![My Skills](https://skillicons.dev/icons?i=instagram)](https://www.instagram.com/lcs.carvalho_/?next=%2F) 
## LeetCode:

[![LeetCode](https://img.shields.io/badge/LeetCode-Lucas--Rabello--Dev-orange?style=for-the-badge&logo=leetcode&logoColor=white)](https://leetcode.com/u/lucas-rabello-dev/)
