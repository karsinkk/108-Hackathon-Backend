FROM busybox

COPY ./main /app/

CMD [“/app/main”]