FROM scratch AS prod

COPY ./build/draftsman /app/draftsman

ENTRYPOINT ["/app/draftsman"]
