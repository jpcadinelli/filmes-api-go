# API de Filmes

## Configuração do Ambiente
Ao fazer o clone do projeto é necessário que pare o postgresql local na sua máquina.
<br>
<br>
No linux, o comando é:

```
$ sudo systemctl stop postgresql
```

<br>
Após parar o postgresql, entre no terminal da pasta raiz do projeto e utilise o comando:

```
$ docker compose up
```

Deixe o docker rodando e abra outro terminal na raiz do projeto e utilize o comando:

```
$ docker-compose exec postgres sh
```

Vai abrir uma linha de comando iniciada com # e nela utilize o comando:

```
$ hostname -i
```

O resultado é o ipadress que deve ser colocado no pgAdmin, ou seja guarde este número.
<br>
<br>
Após isto abra este link [localhost:54321](http://localhost:54321)
<br>
<br>
Nele irá pedir email e senha, que estão configurados nas linhas 17 e 18 do arquivo docker-compose.yml, respectivamente.
<br>
<br>
Ao fazer o login, adicione um novo server e preencha da seguinte forma:

Em General

1. Name: Escolha o nome que quiser;

Em Connection

2. Host name/address: deve ser o valor que foi pedido para guardar anteriormente neste mesmo arquivo;
3. Port: Pode utiliza ra que vem por padrão;
4. Maintanance database: a mesma informada na linha 8 do arquivo docker-compose.yml;
5. Username: o mesmo informado na linha 6 do arquivo docker-compose.yml;
6. Password: a mesma informada na linha 7 do arquivo docker-compose.yml;
<br>
<br>
Após isso, mantenha o docker rodando e em um terminal raiz do projeto ustilize o comando:

```
$ go run main.go
```

E acesse o link [localhost:8080/filmes](http://localhost:8080/filmes).
<br>
<br>

## Planejamento de Melhorias

* [ ] Incluir validação dos dados na criação do filme.
* [ ] Criar script inicial pro banco com mais de 100 filmes.
* [ ] Mudar id de int para uuid.
* [ ] Estudar uma forma diferente de retorno das funções.
* [ ] Adição do Gin Swagger.
* [ ] Criar padrão para documentar mudanças e controlar as versões.
<br>
<br>

## Para Contribuir com o Projeto

Estou utilizando o Padrão de Commits Semânticos para gerênciar as alterações do projeto e facilitar a revisão do código.

Vou estipular algumas regras

- As pull requests devem ser feitas para a hmg.
- Buscar seguir a tabela de commits semânticos que vou deixar discriminado a baixo.
- Caso tenham alguma ideia para melhorar o projeto ou caso tenha dúvida [abra uma issue](https://github.com/jpcadinelli/filmes-api-go/issues/new).
<br>
<br>

### Passos

1. Crie um fork deste repositório.
2. Envie seus commits em português.
3. Insira um checkbox no Planejamento de Melhorias Marcado e com data e nome.
4. Crie um pull request.
<br>
<br>


### Padrão Commits Semânticos

Atualmente seguiremos esta documentação do [iuricode/padroes-de-commits](https://github.com/iuricode/padroes-de-commits).

E para a criação de branches seguiremos estes padrões:

1. **bugfix/**[nome-da-brach]: Para branch de resolução de BUG.
2. **feature/**[nome-da-branch]: Para branch de uma feature que será adicionada ao projeto.
3. **hotfix/**[nome-da-branch]: Para branch de correção de cor, textos, alterações não tão urgentes.
4. **improvement/**[nome-da-branch]: Para branch de melhoria de algo que já exista, seja performance, escrita, layout, etc.