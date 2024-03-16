targetScope = 'subscription'
param location string = 'australiaeast'

resource platformRg 'Microsoft.Resources/resourceGroups@2021-04-01' = {
  name: 'rg-platform-shared'
  location: location
  tags: {
    env: 'shared'
  }
}

module platformSharedModule 'platform-shared/main.bicep' = {
  name: 'platform-shared-module'
  scope: platformRg
  params: {
    containerRegistryName: 'acrab'
    sku: 'Basic'
    location: location
  }
}
