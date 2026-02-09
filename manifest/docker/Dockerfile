FROM loads/alpine:3.8

###############################################################################
#                                INSTALLATION
###############################################################################

ENV WORKDIR  /app

# ADD ./bin/linux_amd64/main $WORKDIR/main
ADD ./gfast $WORKDIR/main
RUN chmod +x $WORKDIR/main

# 添加资源配置文件
ADD resource $WORKDIR/resource/

# 添加配置文件
ADD ./manifest/config/config.yaml $WORKDIR/
RUN chmod 644 $WORKDIR/config.yaml

###############################################################################
#                                   START
###############################################################################
WORKDIR $WORKDIR
CMD ./main
