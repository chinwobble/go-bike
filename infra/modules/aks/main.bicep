@description('The name of the Managed Cluster resource.')
param clusterName string

@allowed([
  'dev'
  'staging'
  'prod'
])
param environment string

@description('The location of the Managed Cluster resource.')
param location string = resourceGroup().location

@description('Optional DNS prefix to use with hosted Kubernetes API server FQDN.')
param dnsPrefix string

@description('Disk size (in GB) to provision for each of the agent pool nodes. This value ranges from 0 to 1023. Specifying 0 will apply the default disk size for that agentVMSize.')
@minValue(0)
@maxValue(1023)
param osDiskSizeGB int = 0

@description('The number of nodes for the cluster.')
@minValue(1)
@maxValue(50)
param agentCount int

@description('The number of nodes for the cluster.')
@minValue(1)
@maxValue(50)
param agentMinCount int

@description('The number of nodes for the cluster.')
@minValue(1)
@maxValue(50)
param agentMaxCount int


@description('The size of the Virtual Machine.')
@allowed([
  'Standard_E2as_v5'
  'Standard_E4as_v5'
  'Standard_E8as_v5'
  'Standard_E16as_v5'
  'Standard_E32as_v5'

  'Standard_D2as_v5'
  'Standard_D4as_v5'
  'Standard_D8as_v5'
  'Standard_D16as_v5'
  'Standard_D32as_v5'
])
param agentVMSize string = 'Standard_D2as_v5'

// @description('User name for the Linux Virtual Machines.')
// param linuxAdminUsername string

// @description('Configure all linux machines with the SSH RSA public key string. Your key should include three parts, for example \'ssh-rsa AAAAB...snip...UcyupgH azureuser@linuxvm\'')
// param sshRSAPublicKey string

@allowed([
  '1.28'
])
param kubernetesVersion string

param nodeResourceGroup string

resource aks 'Microsoft.ContainerService/managedClusters@2023-10-01' = {
  name: clusterName
  location: location
  identity: {
    type: 'SystemAssigned'
  }
  tags: {
    env: environment
  }
  sku: {
    // Free, Premium, Standard
    tier: 'Free'
    name: 'Base'
  }
  properties: {
    kubernetesVersion: kubernetesVersion
    enableRBAC: false
    nodeResourceGroup: nodeResourceGroup
    disableLocalAccounts: false
    dnsPrefix: dnsPrefix
    autoUpgradeProfile: {
      upgradeChannel: 'patch'
      nodeOSUpgradeChannel: 'NodeImage'
    }
    supportPlan: 'KubernetesOfficial'
    agentPoolProfiles: [
      {
        name: 'agentpool'
        osDiskSizeGB: osDiskSizeGB
        count: agentCount
        vmSize: agentVMSize
        enableAutoScaling: true
        minCount: agentMinCount
        maxCount: agentMaxCount
        osSKU: 'Ubuntu'
        osType: 'Linux'
        mode: 'System'
      }
    ]
    // linuxProfile: {
    //   adminUsername: linuxAdminUsername
    //   ssh: {
    //     publicKeys: [
    //       {
    //         keyData: sshRSAPublicKey
    //       }
    //     ]
    //   }
    // }
  }
}

output controlPlaneFQDN string = aks.properties.fqdn
