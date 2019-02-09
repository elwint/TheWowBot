#!/bin/sh
read -p "IP address: " ip
read -p "Port: " port
read -p "Token: " token
curl -F "url=https://$ip:$port/$token" -F "allowed_updates=[\"message\", \"inline_query\"]" -F "certificate=@wow.pem" https://api.telegram.org/bot$token/setWebhook
