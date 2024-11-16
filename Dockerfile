# Build image
FROM golang:1.23.3 AS build

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o /app/build/draftsman .

FROM scratch AS prod

COPY --from=build /app/build/draftsman /app/draftsman

ENTRYPOINT ["/app/draftsman"]
