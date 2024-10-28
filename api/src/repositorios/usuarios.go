package repositorios

import (
	"api/src/modelos"
	"database/sql"
	"fmt"
)

// Usuarios repositório de usuários
type Usuarios struct {
	db *sql.DB
}

// NovoRepositorioUsuarios cria um repositório para usuários
func NovoRepositorioUsuarios(db *sql.DB) *Usuarios {
	return &Usuarios{db}
}

// Criar insere usuário no banco de dados
func (repositorio Usuarios) Criar(usuario modelos.Usuario) (uint64, error) {
	statement, erro := repositorio.db.Prepare("INSERT INTO usuarios (nome, nick, email, senha) VALUES (?, ?, ?, ?)")
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)
	if erro != nil {
		return 0, erro
	}

	ultimoID, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoID), nil
}

// Buscar busca usuários por filtro
func (repositorio Usuarios) Buscar(stringBusca string) ([]modelos.Usuario, error) {
	stringBusca = fmt.Sprintf("%%%s%%", stringBusca)

	linhas, erro := repositorio.db.Query(
		"SELECT id, nome, nick, email, criado_em FROM usuarios WHERE nome LIKE ? OR nick LIKE ?",
		stringBusca,
		stringBusca,
	)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var usuarios []modelos.Usuario

	for linhas.Next() {
		var usuario modelos.Usuario

		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil {
			return nil, erro
		}

		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil
}

// BuscarPorID retorna um usuário por ID
func (repositorio Usuarios) BuscarPorID(ID uint64) (modelos.Usuario, error) {
	linhas, erro := repositorio.db.Query(
		"SELECT id, nome, nick, email, criado_em FROM usuarios WHERE id = ?",
		ID,
	)
	if erro != nil {
		return modelos.Usuario{}, erro
	}
	defer linhas.Close()

	var usuario modelos.Usuario
	if linhas.Next() {
		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil {
			return modelos.Usuario{}, erro
		}
	}

	return usuario, nil
}
