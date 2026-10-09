echo 'Tien hanh Build Docker cho app BarCodes'

echo 'Login vao Docker su dung Credential mac dinh tu truoc'
docker login


echo 'Build spring-rest-project for using ConfigMap'
export MAIN_SERVICE=tdchien1982/go:barcodes-1.0
docker build . -t $MAIN_SERVICE
docker push $MAIN_SERVICE
cd ..
