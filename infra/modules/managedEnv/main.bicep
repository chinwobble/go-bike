
@description('Specifies the location for all resources.')
param location string = resourceGroup().location
param suffix string

@allowed([
  'dev'
  'staging'
  'prod'
])
param environment string

param logAnalyticsCustomerId string
param logAnalyticsprimarySharedKey string

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
        customerId: logAnalyticsCustomerId
        sharedKey: logAnalyticsprimarySharedKey
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
