# Utilisation de golang comme image de base
# Le GOPATH par défaut de cette image est /go.
FROM golang

# Copie des sources de notre projet dans le container,
# dans notre cas le main.go
COPY . /go/src/pao

# Lancement de la compilation avec go install
RUN go install pao

# Définissions du point d'entré de notre programme compilé
ENTRYPOINT /go/bin/pao

# Le port sur lequel notre serveur écoute
EXPOSE 8080