#!/bin/bash

domains=("example.com www.example.com")
email="main@example.com" # Adding a valid address is strongly recommended
staging=1 # Set to 1 if you're testing your setup to avoid hitting request limits

data_path="./certbot"
rsa_key_size=4096
regex="([^www.].+)"

# root required
if [ "$EUID" -ne 0 ]; then echo "Please run $0 as root." && exit; fi

clear

# Menu for existing folder
for domain in ${domains[@]}; do
  domain_name=`echo $domain | grep -o -P $regex`
  if [ -d "$data_path/conf/live/$domain_name" ]; then
    echo "### Existing data found for some domains..."
    echo
    PS3='Your choice: '
    select opt in "Skip registered domains" "Remove registered domains and continue" "Remove registered domains and exit" "Exit"; do
      echo; echo;
      case $REPLY in
          1) echo " Installed certificates will be skipped" echo; echo; break;;
          2) echo " Old certificates removed"; echo; echo; rm -rf "$data_path"; break;;
          3) echo " Old certificates removed"; echo; rm -rf "$data_path"; echo " Exit..."; echo; echo; sleep 2; clear; exit;;
          4) echo " Exit..."; echo; echo; sleep 0.5; clear; exit;;
          *) echo "invalid option $REPLY";;
      esac
    done
    break
  fi
done

mkdir -p "$data_path"

if [ ! -e "$data_path/conf/options-ssl-nginx.conf" ] && [ ! -e "$data_path/conf/ssl-dhparams.pem" ]; then
  echo "### Downloading recommended TLS parameters ..."
  mkdir -p "$data_path/conf"
  curl -s https://raw.githubusercontent.com/certbot/certbot/master/certbot-nginx/certbot_nginx/_internal/tls_configs/options-ssl-nginx.conf > "$data_path/conf/options-ssl-nginx.conf"
  curl -s https://raw.githubusercontent.com/certbot/certbot/master/certbot/certbot/ssl-dhparams.pem > "$data_path/conf/ssl-dhparams.pem"
fi

# Dummy certificate
for domain in ${!domains[*]}; do
  domain_set=(${domains[$domain]})
  domain_name=`echo ${domain_set[0]} | grep -o -P $regex`

  mkdir -p "$data_path/conf/live/$domain_name"

  if [ ! -e "$data_path/conf/live/$domain_name/cert.pem" ]; then
    echo "### Creating dummy certificate for $domain_name domain..."
    path="/etc/letsencrypt/live/$domain_name"
    docker-compose run --rm --entrypoint "openssl req -x509 -nodes -newkey rsa:1024 \
    -days 1 -keyout '$path/privkey.pem' -out '$path/fullchain.pem' -subj '/CN=localhost'" certbot
  fi

done

echo "### Starting nginx ..."
# Restarting for case if nginx container is already started
docker-compose up -d nginx && docker-compose restart nginx

# Select appropriate email arg
case "$email" in
  "") email_arg="--register-unsafely-without-email" ;;
  *) email_arg="--email $email" ;;
esac

# Enable staging mode if needed
if [ $staging != "0" ]; then staging_arg="--staging"; fi

for domain in ${!domains[*]}; do
  domain_set=(${domains[$domain]})
  domain_name=`echo ${domain_set[0]} | grep -o -P $regex`

  if [ -e "$data_path/conf/live/$domain_name/cert.pem" ]; then
    echo "Skipping $domain_name domain"; else

    echo "### Deleting dummy certificate for $domain_name domain ..."
    rm -rf "$data_path/conf/live/$domain_name"


    echo "### Requesting Let's Encrypt certificate for $domain_name domain ..."

    #Join $domains to -d args
    domain_args=""
    for domain in "${domain_set[@]}"; do
      domain_args="$domain_args -d $domain"
    done

    mkdir -p "$data_path/www"
    docker-compose run --rm --entrypoint "certbot certonly --webroot -w /var/www/certbot --cert-name $domain_name $domain_args \
    $staging_arg $email_arg --rsa-key-size $rsa_key_size --agree-tos --force-renewal --non-interactive" certbot
  fi
done
