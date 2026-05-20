# Ports and Adapters Architecture

## O que é?

É um estilo de design de código que tem como objetivo isolar ou blindar o core do negócio das mudanças ou implementações do mundo externo. Através dos conceitos de Portas e Adaptadores.

Temos Portas de Entrada como por exemplo:
* HTTP
* gRPC
* CLI
* Broker de leitura de mensagens

Temos Portas de Saídas como por exemplo:
* Banco de dados
    * Postgres
    * MYSQL
    * Mongo
* API de um parceiro
* SDK de um fornecedor
* Broker para escrita de mensagens

Essas portas irão definir como interagir com o core, são uma espécie de contrato. 

A implementação desse contrato para interagir com o core são os Adaptadores. Que irão interagir com o core através da camada de aplicação ou usecases, que dependem de uma porta.

## Pontos Positivos
- Core de negócio com baixo acoplamento
- Código mais testável
- Mudanças de fornecedor ou tecnologia não afetam o core

## Pontos Negativos
- Curva de aprendizado alta
- Verbosidade
- Em ambientes menos complexos pode trazer complexidade desnecessária

# Objetivo do Projeto
Aplicar na risca os conceitos da arquitetura de portas e adaptadores em um sistema de Pedido.

## Estrutura de Pastas