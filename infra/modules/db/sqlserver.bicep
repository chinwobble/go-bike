metadata description = 'Creates an Azure SQL Server instance.'
param location string = resourceGroup().location
param environment string

param appUser string = 'appUser'
param sqlAdmin string = 'sqlAdmin'
@allowed([
  'Free'
  'Basic'
  'S0'
  'S1'
  'S2'
  'S3'
  'S4'
  'S6'
  'S7'
  'S9'
  'S12'
  'P1'
  'P2'
  'P4'
  'P6'
  'P11'
  'P15'
])
param skuName string

@secure()
param sqlAdminPassword string = newGuid()

param suffix string

resource sqlServer 'Microsoft.Sql/servers@2022-05-01-preview' = {
  name: '${environment}-dbserver-${suffix}'
  location: location
  tags: {
    env: environment
  }
  properties: {
    version: '12.0'
    minimalTlsVersion: '1.2'
    publicNetworkAccess: 'Enabled'
    administratorLogin: sqlAdmin
    administratorLoginPassword: sqlAdminPassword
  }

  resource database 'databases' = {
    name: '${environment}-db-${suffix}'
    location: location
    sku: {
      name: skuName
    }
    tags: {
      env: environment
    }
  }

  resource firewall 'firewallRules' = {
    name: 'Azure Services'
    properties: {
      // Allow all clients
      // Note: range [0.0.0.0-0.0.0.0] means "allow all Azure-hosted clients only".
      // This is not sufficient, because we also want to allow direct access from developer machine, for debugging purposes.
      startIpAddress: '0.0.0.1'
      endIpAddress: '255.255.255.254'
    }
  }
}

// resource sqlDeploymentScript 'Microsoft.Resources/deploymentScripts@2020-10-01' = {
//   name: '${name}-deployment-script'
//   location: location
//   kind: 'AzureCLI'
//   properties: {
//     azCliVersion: '2.37.0'
//     retentionInterval: 'PT1H' // Retain the script resource for 1 hour after it ends running
//     timeout: 'PT5M' // Five minutes
//     cleanupPreference: 'OnSuccess'
//     environmentVariables: [
//       {
//         name: 'APPUSERNAME'
//         value: appUser
//       }
//       {
//         name: 'APPUSERPASSWORD'
//         secureValue: appUserPassword
//       }
//       {
//         name: 'DBNAME'
//         value: databaseName
//       }
//       {
//         name: 'DBSERVER'
//         value: sqlServer.properties.fullyQualifiedDomainName
//       }
//       {
//         name: 'SQLCMDPASSWORD'
//         secureValue: sqlAdminPassword
//       }
//       {
//         name: 'SQLADMIN'
//         value: sqlAdmin
//       }
//     ]

//     scriptContent: '''
// wget https://github.com/microsoft/go-sqlcmd/releases/download/v0.8.1/sqlcmd-v0.8.1-linux-x64.tar.bz2
// tar x -f sqlcmd-v0.8.1-linux-x64.tar.bz2 -C .

// cat <<SCRIPT_END > ./initDb.sql
// drop user if exists ${APPUSERNAME}
// go
// create user ${APPUSERNAME} with password = '${APPUSERPASSWORD}'
// go
// alter role db_owner add member ${APPUSERNAME}
// go
// SCRIPT_END

// ./sqlcmd -S ${DBSERVER} -d ${DBNAME} -U ${SQLADMIN} -i ./initDb.sql
//     '''
//   }
// }

output connectionString string = 'Server=${sqlServer.properties.fullyQualifiedDomainName}; Database=${sqlServer::database.name}; User=${appUser}'
