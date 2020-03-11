# Teste Pismo

Projeto criado para teste prático com GO para Pismo.
O projeto consiste em realizar operação de pagamento de uma transação.

Todas as saídas dos resultados das operações estãos no stdout do docker que está executando o projeto.

# Configuração

As pendências do projeto:
* docker

# Iniciar o projeto

Baixe o projeto: git clone ..
Execute os comandos:
* cd pismo
* docker build -t execute_project:pismo . 
* docker run -p 8090:8090 --name execute_project --rm execute_project:pismo

# libs
    * go get -u github.com/gorilla/mux
    " go get -u github.com/kataras/tablewriter"
    * go get -u github.com/landoop/tableprinter

# Rotas

## Accounts
GET /v1/accounts/<id>/limits

POST /v1/accounts
    {
        "available_credit_limit": {
            "amount": 10.00
        },
        "available_withdrawal_limit": {
            "amount": 10.00
        }
    }
PATCH /v1/accounts/<id>
    {
        "available_credit_limit": {
            "amount": 123.45 // enviar valor negativo para subtrair
        },
        "available_withdrawal_limit": {
            "amount": 123.45 // enviar valor negativo para subtrair
        }
    }

## Transaction

POST /v1/transactions
    {
        "account_id": 1,
        "operation_type_id": 1,
        "amount": 123.45
    }

## Payments

POST /v1/payments (pode enviar multiplos pagamentos)
    [
        {
            "account_id": 1,
            "amount": 123.45
        } ,
        {
            "account_id": 1,
            "amount": 423.45
        } ,
    ]
