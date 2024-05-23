npm run build
cp -r ../dist ./
docker build -t deploy-demo:v0.0.1
docker push deploy-demo:v0.0.1