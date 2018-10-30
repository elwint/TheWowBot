#!/bin/sh
openssl req -newkey rsa:4096 -sha256 -nodes -keyout wow.key -x509 -days 365 -out wow.pem
