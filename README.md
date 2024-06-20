# About

This is the command to convert AWS Cloudfront access log to CSV or JSON format.

# Sample
```
#Version: 1.0
#Fields: date time x-edge-location sc-bytes c-ip cs-method cs(Host) cs-uri-stem sc-status cs(Referer) cs(User-Agent) cs-uri-query cs(Cookie) x-edge-result-type x-edge-request-id x-host-header cs-protocol cs-bytes time-taken x-forwarded-for ssl-protocol ssl-cipher x-edge-response-result-type cs-protocol-version fle-status fle-encrypted-fields c-port time-to-first-byte x-edge-detailed-result-type sc-content-type sc-content-len sc-range-start sc-range-end
2024-06-17	07:35:31	NRT00-A0	1195	xxx.xxx.xxx.xxx	POST	xxxxxx.cloudfront.net	/xxxxx	200	https://sample.com/xxxx	Mozilla/5.0xxxxxxxxxxxx	xxxxx	-	Ok	xxxx.sample.com	https	159444	0.316	-	TLSv1.3	TLS_AES_128_GCM_SHA256	Ok	HTTP/2.0	-	-	52358	0.316	-	text/html	951	-	-
2024-06-17	07:35:31	NRT00-A0	1195	xxx.xxx.xxx.xxx	POST	xxxxxx.cloudfront.net	/xxxxx	200	https://sample.com/xxxx	Mozilla/5.0xxxxxxxxxxxx	xxxxx	-	Ok	xxxx.sample.com	https	159444	0.316	-	TLSv1.3	TLS_AES_128_GCM_SHA256	Ok	HTTP/2.0	-	-	52358	0.316	-	text/html	951	-	-
2024-06-17	07:35:31	NRT00-A0	1195	xxx.xxx.xxx.xxx	POST	xxxxxx.cloudfront.net	/xxxxx	200	https://sample.com/xxxx	Mozilla/5.0xxxxxxxxxxxx	xxxxx	-	Ok	xxxx.sample.com	https	159444	0.316	-	TLSv1.3	TLS_AES_128_GCM_SHA256	Ok	HTTP/2.0	-	-	52358	0.316	-	text/html	951	-	-
2024-06-17	07:35:31	NRT00-A0	1195	xxx.xxx.xxx.xxx	POST	xxxxxx.cloudfront.net	/xxxxx	200	https://sample.com/xxxx	Mozilla/5.0xxxxxxxxxxxx	xxxxx	-	Ok	xxxx.sample.com	https	159444	0.316	-	TLSv1.3	TLS_AES_128_GCM_SHA256	Ok	HTTP/2.0	-	-	52358	0.316	-	text/html	951	-	-
2024-06-17	07:35:31	NRT00-A0	1195	xxx.xxx.xxx.xxx	POST	xxxxxx.cloudfront.net	/xxxxx	200	https://sample.com/xxxx	Mozilla/5.0xxxxxxxxxxxx	xxxxx	-	Ok	xxxx.sample.com	https	159444	0.316	-	TLSv1.3	TLS_AES_128_GCM_SHA256	Ok	HTTP/2.0	-	-	52358	0.316	-	text/html	951	-	-
```

# How to use
```
$ cd <repository root>
$ go build -o <command name you like>
$ ./<command name you like> -f <cloudfront access log file path> -o <output file path .csv>
```

