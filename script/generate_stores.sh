#!/bin/bash

# Função para gerar um valor aleatório para cada parâmetro
generate_random_value() {
    local length=$1
    local random_value=""
    for ((i=0; i<length; i++))
    do
        random_digit=$((RANDOM % 10))
        random_value="${random_value}${random_digit}"
    done
    echo $random_value
}

# Loop para enviar as requisições
for ((i=1; i<=1000; i++))
do
    # Gerar valores aleatórios para os parâmetros
    fantasy_name=$(generate_random_value 10)
    corporate_name=$(generate_random_value 10)
    cnpj=$(generate_random_value 14)

    # Montar o JSON com os parâmetros
    json="{\"fantasy_name\":\"$fantasy_name\",\"corporate_name\":\"$corporate_name\",\"cnpj\":\"$cnpj\"}"

    echo $json

    # Enviar a requisição POST
    curl -X POST -H "Content-Type: application/json" -d "$json" http://localhost:5000/v1/stores

    # Gerar um intervalo aleatório entre 1 e 5 segundos
    sleep $((RANDOM % 5 + 1))
done