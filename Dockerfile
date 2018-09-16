FROM iron/base

EXPOSE 8080

ADD kal-shopping-linux-amd64 /
ADD data.db /

ENTRYPOINT ["./kal-shopping-linux-amd64"]