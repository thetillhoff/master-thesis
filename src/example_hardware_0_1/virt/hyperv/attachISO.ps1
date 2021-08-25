[CmdletBinding()]
param (
  [Parameter(Mandatory)]
  [string]
  $name,
  [Parameter(Mandatory)]
  [string]
  $path
)

#####

# Add DVDDrive to VM
Add-VMScsiController `
-VMName "$name"

# Mount ISO
Add-VMDvdDrive `
-VMName "$name" `
-ControllerNumber 1 `
-ControllerLocation 0 `
-Path "$path"

#Set-VMDvdDrive -VMName DC -ControllerNumber 1 -Path "<path to ISO>