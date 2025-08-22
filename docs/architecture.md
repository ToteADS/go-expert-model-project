# Arquitetura do Projeto projeto-modelo

Este documento descreve a estrutura e organização do projeto seguindo os padrões da comunidade Go.

## Estrutura de Diretórios

### /cmd
**Aplicações principais do projeto**
- `cmd/projeto-modelo/main.go` - Ponto de entrada da aplicação

### /internal
**Código privado da aplicação**
- `internal/app/projeto-modelo/` - Lógica da aplicação
- `internal/pkg/` - Bibliotecas internas compartilhadas

### /pkg
**Bibliotecas públicas que podem ser usadas por outros projetos**
- Código que pode ser importado por aplicações externas

### /api
**Especificações de API**
- OpenAPI/Swagger specs
- Arquivos de esquema JSON
- Definições de protocolo

### /web
**Componentes web**
- Assets estáticos
- Templates do servidor
- SPAs

### /configs
**Arquivos de configuração**
- Templates de configuração
- Configurações padrão

### /build
**Scripts de build e CI/CD**
- `build/package/` - Scripts de empacotamento
- `build/ci/` - Configurações de CI/CD

### /deployments
**Configurações de deploy**
- Docker, Kubernetes, Terraform
- Configurações de orquestração

### /test
**Testes externos e dados de teste**
- Testes de integração
- Dados de teste

### /docs
**Documentação do projeto**
- Documentação técnica
- Guias de usuário

### /tools
**Ferramentas de suporte**
- Scripts utilitários
- Ferramentas de desenvolvimento

### /examples
**Exemplos de uso**
- Exemplos de como usar a aplicação
- Exemplos de integração

## Convenções

1. **Nomenclatura**: Use nomes descritivos e em inglês
2. **Estrutura**: Mantenha a hierarquia de diretórios
3. **Imports**: Use imports relativos ao módulo
4. **Testes**: Coloque testes junto com o código
5. **Documentação**: Mantenha documentação atualizada

## Comandos Úteis

- `make build` - Compilar a aplicação
- `make test` - Executar testes
- `make run` - Executar a aplicação
- `make deps` - Instalar dependências
- `make fmt` - Formatar código
- `make lint` - Executar linter

## Próximos Passos

1. Implementar a lógica da aplicação em `internal/app/projeto-modelo/`
2. Adicionar testes em `internal/app/projeto-modelo/`
3. Configurar CI/CD em `build/ci/`
4. Documentar APIs em `api/`
5. Adicionar exemplos em `examples/`
