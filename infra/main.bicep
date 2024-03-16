targetScope = 'subscription'
param location string = 'australiaeast'

@allowed([
  'dev'
  'staging'
  'prod'
])
param environment string

resource platformRg 'Microsoft.Resources/resourceGroups@2021-04-01' = {
  name: '${environment}-rg-platform'
  location: location
  tags: {
    env: environment
  }
}

module platform 'platform/main.bicep' = {
  name: 'platform-module'
  scope: platformRg
  params: {
    environment: environment
    location: location
    suffix: 'ab'
  }
}


resource tradingAppRg 'Microsoft.Resources/resourceGroups@2021-04-01' = {
  name: '${environment}-rg-trading'
  location: location
  tags: {
    env: environment
  }
}


module tradingApp 'apps/trading/main.bicep' = {
  name: 'trading-app-module'
  scope: tradingAppRg
  params: {
    environment: environment
    location: location
    suffix: 'trading'
  }
}
