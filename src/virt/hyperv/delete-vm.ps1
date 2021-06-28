param (
  [Parameter(Mandatory)]
  [string]
  $name
)

#####

# Force-Stop the VM before deleting it
Stop-VM `
-Name "$name" `
-TurnOff

# Retrieve disk of vm before deleting the vm
$diskpath = Get-VM "$name" | Select-Object -ExpandProperty HardDrives | Select-Object -Expand Path

# Delete vm
Remove-VM `
-Name "$name" `
-Force

# Delete disk
Remove-Item -Path "$diskpath"

# Delete folder containing previously deleted disk
Remove-Item -Path "$(Get-VMHost | Select-Object -ExpandProperty VirtualHardDiskPath)/$name/" # halts&promtps when the folder is non-empty
