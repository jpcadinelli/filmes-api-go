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