# Certification

Getting cert using lets encrypt:

```shell
certbot certonly --manual --preferred-challenges dns -d example.com -d www.example.com --config-dir ./my-certs --work-dir ./my-certs --logs-dir ./my-certs
```

- This certificate will not be renewed automatically. Autorenewal of --manual certificates requires the use of an authentication hook script (--manual-auth-hook) but one was not provided. To renew this certificate, repeat this same certbot command before the certificate's expiry date.

For `nginx` cert use `--nginx` option.