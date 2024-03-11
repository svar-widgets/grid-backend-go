FROM debian:10-slim
WORKDIR /app
ADD ./grid-backend-go /app

CMD ["/app/grid-backend-go"]