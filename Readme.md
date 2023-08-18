# Microserviços PicPay - Golang

Codificação em Golang com Framework GIN para uma implementação de Microserviços para gerenciamento financeiro e de seguros estilo da aplicação PicPay. Ele é composto por vários serviços, cada um responsável por uma funcionalidade específica. O projeto foi implementado com Docker para subir conteiners, Kubernetes para gerenciar os serviços e pods da aplicação, essa aplicação utiliza MongoDB um banco de dado não relacional, Apache Kafka para fazer a parte gerenciamento de Tópicos e Mensagens de Fila, para gerenciamento de Logs e Metrificação usamos o Promoteus e Elastic Search, também está implementado uma Api-Gateway para gerenciamentos de Endpoints de forma funcional e mais segura e implementamos o padrão Circuit Breaker é uma técnica de design de software que melhora a resiliência de um sistema ao prevenir que este faça chamadas a operações que provavelmente falharão. Em vez de esperar que uma operação falhe, o Circuit Breaker pode interromper a operação por um período de tempo.
A aplicação também usa Terraform para Infra que configura um provedor Google Cloud Platform (GCP) e cria um Google Kubernetes Engine (GKE) cluster. Vou descrever cada parte do código e, em seguida, fornecer um exemplo de como você pode estender isso para implantar um microserviço.


## Pré-requisitos

- Golang 1.20
- Framework Gin
- Docker e Docker Compose
- MongoDB
- Kafka
- -Prometeus
- Elastic Search
- WSL 2


## Serviços

- **Carteira** (Porta: 8081): Gerencia a carteira do usuário.
- Monte a imagem de cada módulo: docker build -t carteira:latest .
  
- **Cobrar** (Porta: 8082): Serviço de cobrança.
-  Monte a imagem de cada módulo: docker build -t cobrar:latest .
    
- **Crypto** (Porta: 8083): Serviço relacionado a criptomoedas.
- Monte a imagem de cada módulo: docker build -t crypto:latest .
    
- **Empréstimos** (Porta: 8084): Gerencia empréstimos.
-  Monte a imagem de cada módulo: docker build -t cemprestimos:latest .
    
- **Finance** (Porta: 8085): Serviço financeiro geral.
-  Monte a imagem de cada módulo: docker build -t finance:latest .
   
- **Investir** (Porta: 8086): Serviço de investimentos.
-  Monte a imagem de cada módulo: docker build -t investir:latest .
    
- **Notificações** (Porta: 8087): Gerencia notificações do usuário.
-  Monte a imagem de cada módulo: docker build -t notificacoes:latest .
    
- **Pagar** (Porta: 8088): Serviço de pagamentos.
-  Monte a imagem de cada módulo: docker build -t pagar:latest .
    
- **Sacar** (Porta: 8089): Serviço de saques.
- Monte a imagem de cada módulo: docker build -t sacar:latest .
    
- **Seguros** (Porta: 8090): Gerencia seguros.
- Monte a imagem de cada módulo: docker build -t seguros:latest .

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

## Rode a aplicação de Micoservices no diretório raiz:

go run Main.go

## Use Docker Compose para iniciar todos os serviços:

docker-compose up -d

## Testando os Serviços
Cada serviço pode ser acessado em sua respectiva porta. Por exemplo, para acessar o serviço de Carteira:

curl http://localhost:8081/[ENDPOINT_ESPECÍFICO]

## Monitoramento
O Prometheus e o ElasticSearch estão configurados para monitorar a saúde e o desempenho dos serviços. Acesse suas interfaces web para visualizar métricas e logs.

## Autor:
Emerson Amorim
