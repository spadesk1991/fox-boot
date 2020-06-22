FROM xs23933/alpinecst
WORKDIR /app
ADD server ./
ADD config ./config

CMD ["./server"]
