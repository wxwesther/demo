FROM oraclelinux:7.4
COPY ./demo /wxw/
EXPOSE 5000
CMD ["/wxw/demo"]
