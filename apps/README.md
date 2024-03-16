# dotnet templates with dapr integration

This folder contains a boilerplate for dotnet and dapr hosted on azure container apps.

Features:
- health checks
- liveliness probe, readiness and startup probes
- metrics / logs / traces to grafana
- scheduled tasks with azure container jobs
- Exporters to grafana
- Azure User Assigned Managed Identity to pull private images without the need for secrets

## Run locally
You need to install dapr locally.

After installing dapr you run the below command to start the side cars.
```bash
dapr init
```

run the apps with sidecars
```
C:\dev\a-and-b-tech\infra\apps\src\Template.Api>dapr run --app-id template-api --app-port 5115 --dapr-http-port 3500 --resources-path ..\..\components -- dotnet run
```

this will register all the input and output bindings defined in the components folder.

## Deploying the apps to azure
`aca-components` folder is how the intput / output bindings are defined configured in the azure container environment.

### Deploy the containers
```bash
az acr build `
  --registry $AZURE_CONTAINER_REGISTRY_NAME `
  --image "template-api" `
  --file 'src/Template.Api/Dockerfile' .
```

```bash
./apps/scripts/deploy-jobs.sh staging {hash}
./apps/scripts/deploy-apps.sh staging {hash}
```

# Authenticate to ACR
```bash
az acr login --name acrab
docker tag local acrab.azurecr.io/local
docker push acrab.azurecr.io/local
```
