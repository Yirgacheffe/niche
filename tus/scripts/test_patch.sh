curl --request PATCH --data-binary
"@/Users/aaron/Downloads/pet-mov.png" http://localhost:8093/api/files/{%id} --header "Upload-Offset: 0"
--header "Expect:" -i --limit-rate 1K