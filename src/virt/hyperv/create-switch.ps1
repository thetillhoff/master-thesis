param (
  [Parameter(Mandatory)]
  [string]
  $name,
  [Parameter(Mandatory)]
  [ValidateSet("Private","Internal")] # External is left out for now (so noone breaks anything)
  [string]
  $type
)

#####

# check if switch "PrivateSwitch" exists, if not, create it

# Get existing switches
$vmswitches = Get-VMSwitch | Select-Object -Expand "Name"

# If names of existing switches don't contain the name of the to-be-created switch
if (-Not $vmswitches.contains($name)) {
  # Create Switch
  New-VMSwitch -Name "$name" -SwitchType "$type" > $null
}
