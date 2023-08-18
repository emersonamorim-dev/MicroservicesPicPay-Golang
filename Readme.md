# Microserviços PicPay - Golang

Codificação em Golang com Framework GIN para uma implementação de Microserviços para gerenciamento financeiro e de seguros estilo da aplicação PicPay. Ele é composto por vários serviços, cada um responsável por uma funcionalidade específica. O projeto foi implementado com Docker para subir conteiners, Kubernetes para gerenciar os serviços e pods da aplicação, essa aplicação utiliza MongoDB um banco de dado não relacional, Apache Kafka para fazer a parte gerenciamento de Tópicos e Mensagens de Fila, para gerenciamento de Logs e Metrificação usamos o Promoteus e Elastic Search, também está implementado uma Api-Gateway para gerenciamentos de Endpoints de forma funcional e mais segura e implementamos o padrão Circuit Breaker é uma técnica de design de software que melhora a resiliência de um sistema ao prevenir que este faça chamadas a operações que provavelmente falharão. Em vez de esperar que uma operação falhe, o Circuit Breaker pode interromper a operação por um período de tempo.
A aplicação também usa Terraform para Infra que configura um provedor Google Cloud Platform (GCP) e cria um Google Kubernetes Engine (GKE) cluster. Vou descrever cada parte do código e, em seguida, fornecer um exemplo de como você pode estender isso para implantar um microserviço

## Serviços

- **Carteira** (Porta: 8081): Gerencia a carteira do usuário.
- **Cobrar** (Porta: 8082): Serviço de cobrança.
- **Crypto** (Porta: 8083): Serviço relacionado a criptomoedas.
- **Empréstimos** (Porta: 8084): Gerencia empréstimos.
- **Finance** (Porta: 8085): Serviço financeiro geral.
- **Investir** (Porta: 8086): Serviço de investimentos.
- **Notificações** (Porta: 8087): Gerencia notificações do usuário.
- **Pagar** (Porta: 8088): Serviço de pagamentos.
- **Sacar** (Porta: 8089): Serviço de saques.
- **Seguros** (Porta: 8090): Gerencia seguros.

## Tecnologias

- **Linguagem**: Go
- **Banco de Dados**: MongoDB
- **Mensageria**: Kafka
- **Monitoramento**: Prometheus e ElasticSearch
- **Orquestração**: Docker e Kubernetes
- **Resiliência**: Circuit Breaker com gobreaker

## Como Rodar

1. Clone este repositório:
```bash
git clone [URL_DO_REPOSITÓRIO]
Navegue até a pasta do projeto:
cd [NOME_DA_PASTA_DO_PROJETO]

## Use Docker Compose para iniciar todos os serviços:

docker-compose up -d

## Testando os Serviços
Cada serviço pode ser acessado em sua respectiva porta. Por exemplo, para acessar o serviço de Carteira:

curl http://localhost:8081/[ENDPOINT_ESPECÍFICO]

## Monitoramento
O Prometheus e o ElasticSearch estão configurados para monitorar a saúde e o desempenho dos serviços. Acesse suas interfaces web para visualizar métricas e logs.

## Autor:
Emerson Amorim
