# Events Driver - Projeto Go

## 🚀 Como Executar Programas Go

### 1. Comandos Básicos

```bash
# Executar diretamente (desenvolvimento)
go run main.go

# Ou executar da raiz do projeto
go run .

# Compilar para binário
go build

# Compilar com nome específico
go build -o events-driver

# Executar o binário compilado
./events-driver
```

### 2. Comandos de Desenvolvimento

```bash
# Verificar se o código compila sem executar
go build

# Formatar código automaticamente (importante!)
go fmt .

# Verificar problemas de qualidade
go vet .

# Executar testes
go test .

# Executar testes com mais detalhes
go test -v .
```

### 3. Gerenciamento de Dependências

```bash
# Adicionar uma dependência
go get github.com/exemplo/pacote

# Limpar dependências não usadas
go mod tidy

# Ver dependências do projeto
go list -m all

# Atualizar dependências
go get -u
```

### 4. Estrutura do Projeto

```
events-driver/
├── go.mod          # Arquivo de configuração do módulo
├── main.go         # Ponto de entrada (deve ser package main)
├── README.md       # Documentação
└── ...
```

## 🎯 Diferenças vs Java/Python

### Java
```bash
# Java precisa compilar primeiro
javac Main.java
java Main
```

### Python
```bash
# Python executa diretamente
python main.py
```

### Go
```bash
# Go pode executar diretamente OU compilar
go run main.go      # Como Python
go build && ./app   # Como Java (mas mais simples)
```

## 🔧 Próximos Passos

1. **Criar funcionalidades** no `main.go`
2. **Organizar código** em pacotes (pastas)
3. **Adicionar testes** (`*_test.go`)
4. **Usar dependências** externas com `go get`
5. **Configurar CI/CD** para deployment

## 📚 Comandos Úteis para Aprender

```bash
# Ver documentação de uma função
go doc fmt.Println

# Ver documentação de um pacote
go doc fmt

# Instalar ferramentas de desenvolvimento
go install golang.org/x/tools/cmd/goimports@latest

# Ver informações sobre Go instalado
go version
go env
```