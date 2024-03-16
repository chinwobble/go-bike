@allowed([
  'Basic'
  'Classic'
  'Standard'
  'Premium'
])
param sku string
param location string = resourceGroup().location
param containerRegistryName string

resource containerRegistry 'Microsoft.ContainerRegistry/registries@2022-02-01-preview' = {
  name: containerRegistryName
  location: location
  sku: {
    name: sku
  }
  properties: {
    //You will need to enable an admin user account in your Azure Container Registry even when you use an Azure managed identity https://docs.microsoft.com/azure/container-apps/containers
    adminUserEnabled: true
  }
  tags: {
    env: 'shared'
  }
}


resource idAcrPull 'Microsoft.ManagedIdentity/userAssignedIdentities@2023-07-31-preview' = {
  name: 'id-acrpull'
  location: location
  tags: {
    env: 'shared'
  }
}


output id string = containerRegistry.id
output name string = containerRegistry.name
output loginServer string = containerRegistry.properties.loginServer
