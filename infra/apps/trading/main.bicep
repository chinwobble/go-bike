
@description('Specifies the location for resources.')
param location string = 'australiaeast'

@allowed([
  'dev'
  'staging'
  'prod'
])
param environment string

param suffix string

targetScope = 'resourceGroup'

resource id 'Microsoft.ManagedIdentity/userAssignedIdentities@2023-07-31-preview' = {
  name: '${environment}-id-${suffix}'
  location: location
  tags: {
    env: environment
  }
}

module db '../../modules/db/sqlserver.bicep' = {
  name: 'db-module'
  params: {
    environment: environment
    suffix: suffix
    location: location
    skuName: 'S0'
  }
}
