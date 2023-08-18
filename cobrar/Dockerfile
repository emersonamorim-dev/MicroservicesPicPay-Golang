FROM golang:1.20

# Define o diretório de trabalho dentro da imagem
WORKDIR /app

# Copia o código fonte para dentro da imagem
COPY . .

# Copia o go.mod e go.sum do diretório modules para o diretório de trabalho atual
COPY modules/go.mod .
COPY modules/go.sum .

# Baixa as dependências do projeto
RUN go mod download

# Compila o código fonte
RUN go build -o main .

# Adiciona permissões de execução ao arquivo main
RUN chmod +x main

COPY start-cobrar.sh /cobrar/start-cobrar.sh
RUN chmod +x /cobrar/start-cobrar.sh

# Expõe a porta em que o servidor irá rodar
EXPOSE 8082

# Define o comando que irá rodar a aplicação
CMD [ "./main" ]
