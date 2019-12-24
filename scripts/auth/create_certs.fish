#!/usr/bin/fish

set -o nounset \
    -o errexit \
    -o verbose \
    -o xtrace

# Generate CA key
openssl req -new -x509 -keyout ../certs/portapp-ca-1.key -out ../certs/portapp-ca-1.crt -days 365 -subj '/CN=ca1' -passin pass:portapp -passout pass:portapp
# openssl req -new -x509 -keyout portapp-ca-2.key -out portapp-ca-2.crt -days 365 -subj '/CN=ca2.test/OU=TEST/O=PORTAPP/L=Washington/S=DC/C=US' -passin pass:portapp -passout pass:portapp

# Kafkacat
openssl genrsa -des3 -passout "pass:portapp" -out ../certs/kafkacat.client.key 2048
openssl req -passin "pass:portapp" -passout "pass:portapp" -key ../certs/kafkacat.client.key -new -out ../certs/kafkacat.client.req -subj '/CN=kafkacat'
openssl x509 -req -CA ../certs/portapp-ca-1.crt -CAkey ../certs/portapp-ca-1.key -in ../certs/kafkacat.client.req -out ../certs/kafkacat-ca1-signed.pem -days 9999 -CAcreateserial -passin "pass:portapp"

for i in broker1 pad.kafka.producer pad.kafka.consumer pad.postgres.client postgres pad.grpc.client localhost
    echo $i
    # Create keystores
    keytool -genkey -noprompt \
        -alias $i \
        -dname "CN=$i" \
        -keystore ../certs/$i.keystore.jks \
        -keyalg RSA \
        -storepass portapp \
        -keypass portapp \
        -storetype pkcs12

    # Create CSR, sign the key and import back into keystore
    keytool -keystore ../certs/$i.keystore.jks -alias $i -certreq -file ../certs/$i.csr -storepass portapp -keypass portapp

    openssl x509 -req -CA ../certs/portapp-ca-1.crt -CAkey ../certs/portapp-ca-1.key -in ../certs/$i.csr -out ../certs/$i-ca1-signed.crt -days 9999 -CAcreateserial -passin pass:portapp

    keytool -keystore ../certs/$i.keystore.jks -alias CARoot -import -file ../certs/portapp-ca-1.crt -storepass portapp -keypass portapp

    keytool -keystore ../certs/$i.keystore.jks -alias $i -import -file ../certs/$i-ca1-signed.crt -storepass portapp -keypass portapp

    # Create truststore and import the CA cert.
    keytool -keystore ../certs/$i.truststore.jks -alias CARoot -import -file ../certs/portapp-ca-1.crt -storepass portapp -keypass portapp

    # convert keystore to Golang readable format
    keytool -importkeystore -srckeystore ../certs/$i.keystore.jks -destkeystore ../certs/$i.p12 -deststoretype PKCS12 -srcstorepass portapp -deststorepass portapp
    openssl pkcs12 -in ../certs/$i.p12 -nokeys -out ../certs/$i.cer.pem -passin pass:portapp
    openssl pkcs12 -in ../certs/$i.p12 -nodes -nocerts -out ../certs/$i.key.pem -passin pass:portapp

    touch ../certs/{$i}_sslkey_creds
    echo portapp >../certs/{$i}_sslkey_creds
    touch ../certs/{$i}_keystore_creds
    echo portapp >../certs/{$i}_keystore_creds
    touch ../certs/{$i}_truststore_creds
    echo portapp >../certs/{$i}_truststore_creds
end
