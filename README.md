# Introdução

Primeira aplicação completa utilizando Golang.

## Utilização

A criação do banco de dados e das tabelas não está automatizada. Precisa executar o arquivo */api/sql/sql.sql* ou o seu conteúdo para as criações.

## Observações

A chave secreta no *.env* (*SECRET_KEY*) foi gerada com seguinte código no arquivo *main.go*:

```
import(
	"crypto/rand"
	"encoding/base64"
)

func init() {
	chave := make([]byte, 64)

	if _, erro := rand.Read(chave); erro != nil {
		log.Fatal(erro)
	}

	stringBase64 := base64.StdEncoding.EncodeToString(chave)
	fmt.Println(stringBase64)
}
```

Esse código foi removido da versão final, mas poderia ter sido implementada uma lógica para gerá-lo apenas uma vez.

## Misc

Originalmente o *.env* estava no Git, para remover um arquivo já acompanhado pelo Git precisa adicioná-lo ao *.gitignore* e rodar os seguintes comandos:

```
git rm --cached folder/file.ext
git add .
# OR
git add [path/to/files/modified]
git commit -m "Commit message"
git pull # if it's necessary
git push origin master
```
