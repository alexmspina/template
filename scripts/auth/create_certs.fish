#!/usr/bin/fish

set -o nounset \
    -o errexit \
    -o verbose \
    -o xtrace

# Generate CA key
openssl req -new -x509 -keyout ../certs/salesadmin-ca-1.key -out ../certs/salesadmin-ca-1.crt -days 365 -subj '/CN=ca1' -passin pass:salesadmin -passout pass:salesadmin
# openssl req -new -x509 -keyout salesadmin-ca-2.key -out salesadmin-ca-2.crt -days 365 -subj '/CN=ca2.test/OU=TEST/O=SALESADMIN/L=Washington/S=DC/C=US' -passin pass:salesadmin -passout pass:salesadmin

for i in salesadmin.postgres.client postgres salesadmin.grpc.client localhost
    echo $i
    # Create keystores
    keytool -genkey -noprompt \
        -alias $i \
        -dname "CN=$i" \
        -keystore ../certs/$i.keystore.jks \
        -keyalg RSA \
        -storepass salesadmin \
        -keypass salesadmin \
        -storetype pkcs12

    # Create CSR, sign the key and import back into keystore
    keytool -keystore ../certs/$i.keystore.jks -alias $i -certreq -file ../certs/$i.csr -storepass salesadmin -keypass salesadmin

    openssl x509 -req -CA ../certs/salesadmin-ca-1.crt -CAkey ../certs/salesadmin-ca-1.key -in ../certs/$i.csr -out ../certs/$i-ca1-signed.crt -days 9999 -CAcreateserial -passin pass:salesadmin

    keytool -keystore ../certs/$i.keystore.jks -alias CARoot -import -file ../certs/salesadmin-ca-1.crt -storepass salesadmin -keypass salesadmin

    keytool -keystore ../certs/$i.keystore.jks -alias $i -import -file ../certs/$i-ca1-signed.crt -storepass salesadmin -keypass salesadmin

    # Create truststore and import the CA cert.
    keytool -keystore ../certs/$i.truststore.jks -alias CARoot -import -file ../certs/salesadmin-ca-1.crt -storepass salesadmin -keypass salesadmin

    # convert keystore to Golang readable format
    keytool -importkeystore -srckeystore ../certs/$i.keystore.jks -destkeystore ../certs/$i.p12 -deststoretype PKCS12 -srcstorepass salesadmin -deststorepass salesadmin
    openssl pkcs12 -in ../certs/$i.p12 -nokeys -out ../certs/$i.cer.pem -passin pass:salesadmin
    openssl pkcs12 -in ../certs/$i.p12 -nodes -nocerts -out ../certs/$i.key.pem -passin pass:salesadmin

    touch ../certs/{$i}_sslkey_creds
    echo salesadmin >../certs/{$i}_sslkey_creds
    touch ../certs/{$i}_keystore_creds
    echo salesadmin >../certs/{$i}_keystore_creds
    touch ../certs/{$i}_truststore_creds
    echo salesadmin >../certs/{$i}_truststore_creds
end
