FROM registry.qtt6.cn/library/centos:7

COPY Dockerfiles/files         /

RUN  yum -y update          && \
     yum clean all          && \
     mkdir -p /data/logs    && \ 
     chmod 755 /run.sh /quconfd-dl.sh && \
     echo 'Defaults env_keep += "APP_NAME APP_SERVICE APP_NAMESPACE APP_ENV APP_LABEL APP_ENV_ID"' >> /etc/sudoers