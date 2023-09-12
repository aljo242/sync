cd proto
buf generate
cd ..

cp -r sync/* ./
rm -rf sync

swagger-combine ./docs/config.json -o ./docs/swagger.yml
rm -rf tmp-swagger-gen