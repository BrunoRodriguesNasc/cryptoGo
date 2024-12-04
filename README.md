# CriptoGo

CriptoGo é um projeto em Go que busca as taxas de câmbio de criptomoedas usando a API do CEX.IO. Ele suporta múltiplas moedas e exibe a taxa de câmbio atual para cada uma.

## Estrutura do Projeto

- `api/responses.go`: Define a estrutura para a resposta da API do CEX.IO.
- `api/cex.go`: Contém a função `GetRate` que faz a solicitação HTTP para a API e processa a resposta.
- `datatypes/data.go`: Define o tipo `Rate` usado para armazenar a taxa de câmbio.
- `main.go`: O ponto de entrada do programa, que busca as taxas de câmbio para uma lista de moedas.
- `go.mod`: Arquivo de configuração do módulo Go.

## Pré-requisitos

- Go 1.23.3 ou superior

## Instalação

1. Clone o repositório:

   ```bash
   git clone https://github.com/seu-usuario/criptogo.git
   cd criptogo
   ```

2. Execute o programa:

   ```bash
   go run main.go
   ```

## Uso

O programa busca as taxas de câmbio para as moedas especificadas no arquivo `main.go`. Atualmente, ele está configurado para buscar as taxas de BTC, ETH e SOL. Você pode modificar a lista de moedas no arquivo `main.go` conforme necessário.

## Contribuição

Sinta-se à vontade para abrir issues ou enviar pull requests. Todas as contribuições são bem-vindas!

## Licença

Este projeto está licenciado sob a Licença MIT. Veja o arquivo [LICENSE](LICENSE) para mais detalhes.