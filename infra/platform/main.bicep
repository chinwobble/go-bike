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


resource managedEnv 'Microsoft.App/managedEnvironments@2023-05-01' = {
  name: '${environment}-me'
  location: location
  tags: {
    env: environment
  }
  // kind: 'consumption'
  properties: {
    appLogsConfiguration: {
      destination: 'log-analytics'
      logAnalyticsConfiguration: {
        customerId: logAnalytics.properties.customerId
        sharedKey: logAnalytics.listKeys().primarySharedKey
      }
    }
    peerAuthentication: {
      mtls: {
        enabled: true
      }
    }
    // vnetConfiguration: {
    //   dockerBridgeCidr: 'string'
    //   infrastructureSubnetId: 'string'
    //   internal: bool
    //   platformReservedCidr: 'string'
    //   platformReservedDnsIP: 'string'
    // }
    workloadProfiles: [
      {
        name: 'Consumption'
        workloadProfileType: 'Consumption'
      }
    ]
    zoneRedundant: false
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

