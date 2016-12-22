FROM busybox

COPY ./main /home/

CMD ["/home/main"]