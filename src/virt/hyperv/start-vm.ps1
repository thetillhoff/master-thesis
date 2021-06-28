param (
  [Parameter(Mandatory)]
  [string]
  $name
)

#####

Start-VM `
-Name "$name"
