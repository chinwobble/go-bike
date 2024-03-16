# Azure bicep infrastructure project


## Deploying the Azure Infrastructure
`infra` folder contains a template for deploying two resource groups:

### Platform-shared
Creates resources shared across multiple environments:
- Azure Container Registry
```bash
# choose the correct subscription
# az account set -s "xxxxx"
az deployment sub create -n first --location australiaeast --template-file infra/main-platform-shared.bicep
az role assignment create \
  --role "Key Vault Reader" \
  --assignee-object-id $principalId \
  --scope "subscriptions/$subscriptionId/resourcegroups/$appRg" \
  --assignee-principal-type ServicePrincipal
```

### Environment Specific Resources

creates:
- Container managed env
- vnet (subnet)
- app insights
- log analytics workspace
- key vault
- app configuration

```bash
az deployment sub create -n first --location australiaeast --template-file infra/main.bicep --parameters environment=staging
```


## Assign Permissions to managed Identity
```bash
subscriptionId="614fa76d-a87d-4650-bfc9-51272298fb73"
platformRg="rg-platform"
appRg="staging-rg"
principalId="cfea5bfe-a9c8-48e0-bee3-9989487cb516"
# these commands may fail if the permission is already assigned
az role assignment create \
--role "AcrPull" \
--assignee-object-id "$principalId" \
--scope "subscriptions/$subscriptionId/resourcegroups/$platformRg" \
--assignee-principal-type ServicePrincipal
az role assignment create \
  --role "Key Vault Reader" \
  --assignee-object-id $principalId \
  --scope "subscriptions/$subscriptionId/resourcegroups/$appRg" \
  --assignee-principal-type ServicePrincipal
az role assignment create --role "Key Vault Secrets User" \
--assignee-object-id /$principalId \
--scope "subscriptions/$subscriptionId/resourcegroups/$appRg" \
--assignee-principal-type ServicePrincipal
```


### Grant Access to sql database
You access for the managed identity to the database.
Login to the database and run the following commands.
```sql
create user [staging-id-template] from external provider;
alter role db_owner add member [staging-id-template];
```

## Deploy the container app
Azure Container App configuration is contained in yaml files contained in
`apps/container-app-config/`.

Each subfolder is intended to be a separate resource group.
You can place config for:
- `app-config/` permanently running containers
- `job-config/` containers that are run to completion typically with an event or schedule.

Each container app must have a three yaml files
```
apps/
  container-app-config/
    {resourceGroup}/
      app-config/
        {app}.base.yaml
        {app}.staging.yaml
        {app}.prod.yaml
```

The below script will merge the base and env variants yaml files together.

Deploy the apps using the bash script.
```bash
# apt install yq
./apps/scripts/deploy-apps.sh staging $GITHUB_SHA rg-trading
./apps/scripts/deploy-apps.sh prod $GITHUB_SHA rg-trading
```


## References
https://github.com/Azure/azure-quickstart-templates/blob/master/quickstarts/microsoft.network/vnet-to-vnet-peering/main.bicep

https://azure.github.io/aca-dotnet-workshop/aca/02-aca-comm/
