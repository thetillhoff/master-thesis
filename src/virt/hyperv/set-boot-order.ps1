param (
  [Parameter(Mandatory)]
  [string]
  $name,
  [Parameter(Mandatory)]
  [ValidateSet("disk","iso","net")]
  [string]
  $boot
)

#####

switch ($boot) { #TODO move to set-boot-order.ps1
  "disk" {
    $disk = Get-VMHardDiskDrive -VMName "$name"

    Set-VMFirmware `
    -VMName "$name" `
    -FirstBootDevice $disk
  }
  "iso" {
    Invoke-Expression -Command "`"$PSScriptRoot\attachISO.ps1`" -name `"$name`" -path `"$PSScriptRoot\boot.iso`""

    Set-VMFirmware `
    -VMName "$name" `
    -FirstBootDevice $(Get-VMDvdDrive -VMName "$name")
  }
  "net" {
    $net = Get-VMNetworkAdapter -VMName "$name"

    Set-VMFirmware `
    -VMName "$name" `
    -FirstBootDevice $net
  }
}

