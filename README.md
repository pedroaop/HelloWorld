# Hello World - Go + Jelastic

Este é um projeto simples de exemplo demonstrando uma aplicação web em Go (Golang) servindo uma interface HTML e uma API JSON, configurada para deploy na plataforma Jelastic PaaS.

## 🚀 Funcionalidades

- Servidor Web em Go (net/http).
- API REST (`/api/hello`).
- Frontend HTML/JS estático.
- Configuração pronta para nuvem (leitura da porta via variável de ambiente).

## 🛠️ Como rodar localmente

Certifique-se de ter o [Go instalado](https://go.dev/dl/).

1. Clone o repositório ou baixe os arquivos.
2. Na raiz do projeto, inicialize as dependências (se necessário):
   ```bash
   go mod tidy
   ```
3. Execute a aplicação:
   ```bash
   go run main.go
   ```
4. Acesse no navegador: `http://localhost:8080`

## ☁️ Como fazer Deploy no Jelastic

Este projeto foi otimizado para rodar em containers **Go** no Jelastic. Siga os passos abaixo:

### 1. Criar o Ambiente
1. No painel do Jelastic, crie um novo ambiente.
2. Selecione a linguagem **Go**.
3. Escolha a versão do Go desejada (ex: 1.22 ou superior).

### 2. Configurar o Repositório (Git)
1. Adicione o seu repositório Git ao projeto no Jelastic.
2. O Jelastic fará o clone, baixará as dependências (`go get`) e compilará o projeto (`go build`) automaticamente.

### 3. Configurar Variáveis de Ambiente
Para que a aplicação inicie corretamente, configure as seguintes variáveis no painel do ambiente:

| Variável | Valor Recomendado | Descrição |
|----------|-------------------|-----------|
| `GO_RUN` | `HelloWorld` | **Importante:** Nome do binário gerado. Deve corresponder ao nome do módulo no `go.mod` ou nome do arquivo principal. Evite usar `__AUTO__`. |
| `GOPATH` | `/home/jelastic/webapp` | Caminho padrão onde a aplicação é construída e executada. |
| `PORT` | `8080` | (Opcional) A aplicação lê essa variável automaticamente. O Jelastic redireciona a porta 80 externa para esta porta interna. |

### 4. Verificar o Deploy
1. Após o deploy, verifique o log (`run.log`) para confirmar a inicialização:
   ```text
   Servidor iniciando na porta 8080...
   ```
2. Clique em **"Open in Browser"** para acessar a aplicação.

## 📂 Estrutura de Arquivos Importante
O servidor Go serve arquivos estáticos a partir do diretório de execução.
- `main.go`: Código do servidor.
- `index.html`: Interface do usuário (deve estar na mesma pasta do binário executável ou na raiz de execução `webapp`).
- `go.mod`: Definição do módulo `helloworld`.
