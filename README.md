# fullcycle-deploy-cloud-run


## Rodando o projeto

Execute `go run main.go` e acesse o endpoint localhost:8080/weather?cep=<ALGUM_CEP>

Obs: Nome do projeto definido no GCP: weather-api-270420251943

## Construindo o projeto no GCP
1. Rebuild com --platform=linux/amd64:
```
docker build --platform=linux/amd64 -t gcr.io/weather-api-270420251943/weather-api .
```
2. Push da imagem:
```
docker push gcr.io/weather-api-270420251943/weather-api
```

3. Deploy no Cloud Run:
```
gcloud run deploy weather-api \
  --image gcr.io/weather-api-270420251943/weather-api \
  --platform managed \
  --region us-central1 \
  --allow-unauthenticated \
  --set-env-vars WEATHER_API_KEY=<API_KEY>
```

## Demonstração (link de referência)
https://weather-api-536946588194.us-central1.run.app/weather?cep=29216090