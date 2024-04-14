# generate certs
```bash
openssl req \
-newkey rsa:4096 \
-days 365 \
-nodes \
-x509 \
-subj "/C=KR/ST=Seoul/L=Seoul/O=Company/OU=Unit/CN=Unit" \
-keyout localhost.dev.key -out localhost.dev.crt

chmod 600 localhost.dev.key localhost.dev.crt
```

- crt가 public
- key가 private