[CmdletBinding()]
param (
  [Parameter(Mandatory)]
  [string]
  $name,
  [Parameter(Mandatory)]
  [ValidateRange(1,2)]
  [int]
  $gen,
  [Parameter(Mandatory)]
  [ValidateSet("disk","iso","net")]
  [string]
  $boot,
  [switch]
  $disk
)
# # Get the command name
# $CommandName = $PSCmdlet.MyInvocation.InvocationName;
# # Get the list of parameters for the command
# $ParameterList = (Get-Command -Name $CommandName).Parameters;
#  # Grab each parameter value, using Get-Variable
# foreach ($Parameter in $ParameterList) {
#     Get-Variable -Name $Parameter.Values.Name -ErrorAction SilentlyContinue;
#     #Get-Variable -Name $ParameterList;
# }

#####

$switchname = "PrivateSwitch"

# Create switch if it does not exist yet
Invoke-Expression -Command "& `"$PSScriptRoot\create-switch.ps1`" -name `"$switchname`" -type `"Private`"" # check if switch "PrivateSwitch" exists, if not, create it

if ($disk.IsPresent) {
  # Create VM with new VHDX
  New-VM `
  -Name "$name" `
  -MemoryStartupBytes 2GB `
  -NewVHDPath "$(Get-VMHost | Select-Object -ExpandProperty VirtualHardDiskPath)/$name/$name.vhdx" `
  -NewVHDSizeBytes 30GB `
  -Generation $gen `
  -Switch "$switchname" `
  > $null
} else {
  # Create VM without VHDX
  New-VM `
  -Name "$name" `
  -MemoryStartupBytes 2GB `
  -NoVHD `
  -Generation $gen `
  -Switch "$switchname" `
  > $null
}

# Set SecureBoot in such a way, that linux-os is possible
Set-VMFirmware -VMName "$name" -SecureBootTemplate "MicrosoftUEFICertificateAuthority" > $null
# Set-VMFirmware -VMName "$name" -PreferredNetworkBootProtocol "Ipv4|IPv6"
# Disable automatic checkpoints
Set-VM -VMName "$name" -AutomaticCheckpointsEnabled $False

# Set boot-device of vm
Invoke-Expression -Command "& `"$PSScriptRoot\set-boot-order.ps1`" -name `"$name`" -boot `"$boot`""

# Start VM
Invoke-Expression -Command "& `"$PSScriptRoot\start-vm.ps1`" -name `"$name`""

# Output VM
Get-VM -Name "$name"
