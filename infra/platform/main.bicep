@description('Specifies the location for all resources.')
param location string = resourceGroup().location
param suffix string

@allowed([
  'dev'
  'staging'
  'prod'
])
param environment string

resource logAnalytics 'Microsoft.OperationalInsights/workspaces@2021-06-01' = {
  name: '${environment}-log'
  location: location
  tags: {
    env: environment
  }
  properties: {
    sku: {
      name: 'PerGB2018'
    }
  }
}

// module platform '../modules/managedEnv/main.bicep' = {
//   name: 'managedEnv-module'
//   params: {
//     environment: environment
//     location: location
//     logAnalyticsCustomerId: logAnalytics.properties.customerId
//     logAnalyticsprimarySharedKey: logAnalytics.listKeys().primarySharedKey
//     suffix: 'ab'
//   }
// }

module platform '../modules/aks/main.bicep' = {
  name: 'aks-module'
  params: {
    clusterName: '${environment}-aks'
    nodeResourceGroup: '${environment}-rg-aks-managed'
    environment: environment
    dnsPrefix: 'staging-aks'
    location: location
    kubernetesVersion: '1.28'
    agentMinCount: 1
    agentCount: 1
    agentMaxCount: 1
    // logAnalyticsCustomerId: logAnalytics.properties.customerId
    // logAnalyticsprimarySharedKey: logAnalytics.listKeys().primarySharedKey
    // suffix: 'ab'
  }
}



resource appInsights 'Microsoft.Insights/components@2020-02-02' = {
  name: '${environment}-ai'
  location: location
  tags: {
    env: environment
  }
  kind: 'web'

  properties: {
    Application_Type: 'web'
    IngestionMode: 'LogAnalytics'
    WorkspaceResourceId: logAnalytics.id
  }
}

resource keyVault 'Microsoft.KeyVault/vaults@2023-07-01' = {
  name: '${environment}-kv-${suffix}'
  location: location

  tags: {
    env: environment
  }

  properties: {
    sku: {
      family: 'A'
      name: 'standard'
    }
    tenantId: tenant().tenantId
    accessPolicies: []
  }
}

resource appConfiguration 'Microsoft.AppConfiguration/configurationStores@2023-03-01' = {
  name: '${environment}-config-${suffix}'
  location: location

  tags: {
    env: environment
  }
  sku: {
    // Standard or Free
    name: 'Free'
  }
}

