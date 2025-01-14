FROM jybl/timezone AS tz

FROM frolvlad/alpine-glibc

#修改容器时区
ENV TZ=Asia/Shanghai LANG=C.UTF-8
COPY --from=tz /usr/share/zoneinfo/$TZ /usr/share/zoneinfo/$TZ
RUN echo $TZ > /etc/timezone && ln -sf /usr/share/zoneinfo/$TZ /etc/localtime

WORKDIR /app

ADD ./ddns /app
ENTRYPOINT ["./ddns"]
CMD ["-token",""]
