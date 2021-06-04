# Discord-Clone-API <img src="https://img.shields.io/badge/1.16.4-00ADD8.svg?&style=for-the-badge&logo=go&logoColor=white" height="25"/> <img src="https://img.shields.io/badge/Fiber-v2-FECC00.svg?&style=for-the-badge&logo=go&logoColor=white" height="25"/> <img src="https://img.shields.io/badge/GORM-v1.21.10-F01428.svg?&style=for-the-badge&logo=go&logoColor=white" height="25"/>


Esse projeto serve como um complemento do projeto realizado com Angular e visa simular o funcionamento básico do discord, isto é:

- Sistema de servidores e canais
- Envio de mensagens pelos usuários

Para isso, as tecnologias principais utilizadas foram [Fiber](https://docs.gofiber.io/) (um framework express-like para golang), [GORM](https://gorm.io/) (ORM que serve para uma diversidade de bancos relacionais) e [PostgreSQL](https://www.postgresql.org/) para banco de dados (mas é uma opção que pode ser modificada de maneira bem simples)

## Uso
A exigência principal é ter o Go instalado na sua máquina. Caso ainda não tenha, ele pode ser obtido por meio do [site oficial](https://golang.org/dl/) da linguagem. Com esse setup principal, você já consegue rodar e buildar o sistema, mas se quiser aproveitar de uma mini gambiarra que te dar __"live-reload"__ instale o [entr](http://eradman.com/entrproject/) <s>(usuários de windows, perdão, mas não sei como funciona no seu caso)</s>. Vale ressaltar que os comandos do Go não são usados diretamente nesse processo, pois optei pela criação de um Makefile, caso tenha alguma curiosidade, dá uma olhadinha no arquivo :)

Não esqueça de configurar um arquivo `.env` seguindo o template definido no `.env-example`.

#### Executar
```console
api@discord-clone:$ make run
```
Com esse comando, o projeto irá iniciar de maneira estática e você terá que reiniciar o comando sempre que fizer uma alteração e quiser ver o resultado.
##### Live-reload
```console
api@discord-clone:$ make dev
```
Dessa forma, o projeto irá executar usando o artifício do entr e sempre que ocorrer um novo save nos arquivos `.go` ele irá executar tudo novamente.
#### Buildar
```console
api@discord-clone:$ make build
```
Seu projeto irá compilar e poderá ser encontrado como um executável na pasta `build/`. Vale destacavar que o projeto só irá funcionar se houver um arquivo `.env` no mesmo diretório que ele, mas isso é algo que ainda vou avaliar.
