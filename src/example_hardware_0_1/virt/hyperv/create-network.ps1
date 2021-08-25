# select the networkadapter on the host that you want to use for the virtual switch
$hostnetadapter = Get-NetAdapter -Physical

# Create the new switch
# Private means, only the host is connected
New-VMSwitch -name PrivateSwitch -SwitchType Private
