# ipmi

The only source on how to get actual device information instead of sensor data is https://docs.oracle.com/cd/E19464-01/820-6850-11/IPMItool.html#50602039_53406.
Sadly, this seems to refer to special ROMs.

It seems that ipmi does not support retrieving information about installed components like size of ram etc.
This could be part of the outlook of the MA.

There are badly documented ipmi-calls for lenovo machines, where they can at least partly request device information. Yet, this also doesn't work for all machines (only lenovo, and only some of that).
The outlook of the MA could include, that a proper standard for that would be great.
