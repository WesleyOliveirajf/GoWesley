package models

import "GoWesley/db"

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func BuscaTodosOsProdutos() []Produto {
	db := db.ConectaBancodeDados()
	defer db.Close()

	selectDeTodosOsProdutos, err := db.Query("select * from produtos")
	if err != nil {
		panic(err.Error())
	}
	defer selectDeTodosOsProdutos.Close()

	p := Produto{}
	produtos := []Produto{}

	for selectDeTodosOsProdutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectDeTodosOsProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Quantidade = quantidade
		p.Preco = preco

		produtos = append(produtos, p)
	}

	defer db.Close()
	return produtos

}

func CriaNovoProduto(nome, descricao string, preco float64, quantidade int) {
	db := db.ConectaBancodeDados()
	defer db.Close()

	insereDadosnoBanco, err := db.Prepare("insert into produtos(nome, descricao, preco, quantidade) values ($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}
	defer insereDadosnoBanco.Close()

	_, err = insereDadosnoBanco.Exec(nome, descricao, preco, quantidade)
	if err != nil {
		panic(err.Error())
	}
}

func DeletaProduto(id string) {
	db := db.ConectaBancodeDados()
	defer db.Close()

	deletarProduto, err := db.Prepare("delete from produtos where id=$1")
	if err != nil {
		panic(err.Error())
	}
	defer deletarProduto.Close()

	_, err = deletarProduto.Exec(id)
	if err != nil {
		panic(err.Error())
	}
}
