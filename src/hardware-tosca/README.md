# hardware-tosca

## notes

- vendor and model should be enough to fill all other informations. Create public database fot that. With EAN?
- firmware database (required if used)
- user can then use opa on hardware, f.e. 
  - no server in rack x must be connected to ups a
  - only ecc-ram in prod servers
- no naming for non-human stuff -> serial numbers are more efficient and hw should be managed as cattle anyway.
- formfactor.rackable == 19"rack-1U, -halfwidth, tower
- firmware always with version
- field "buy date"?
- zones could be recommended based on detected topology and some best-practices - opa to the rescue:
  - no zone larger than X servers
  - when a group of servers is only connected via <3 cables, recommend new zone for them
- when every power-socket, room etc is documented properly, it is easy/easier to setup building-automation (energy-supply on/off at workplace/table/room/...), building-security (door-control). Especially in times of home-office and corresponding rare, somewhat "random" days at the workplace.
  

## types

### power cable (or more generic: see "cable" below)
- voltage
- source plug (angled?)
- target plug (angled?)
- length

### cable
- type (power, network, audio, video, usb, ...) (cable.network.copper.CAT6, cable.network.fiber.SingleMode.OS2)
- optional type-specific property like voltage, hdmi-version, usb-version, ...)
  could be integrated into type like power-230v, usb-3.1 -> cable.power.230v, cable.usb.31
- source plug (angled?)
- target plug (angled=?)
- length
- vendor (sometimes important)

### pdu
- fuse ampere
- output jacks
  - max jacks type a ...
- input jacks
  - max jacks type a ...
- vendor
- model
- serial number

### psu
- formfactor
- input jacks
  - max jacks type a
- vendor
- model
- serial number
- formfactor.psu (ATX,...)

### ups
- formfactor
- battery size (kwh)
- input jacks
  - max jacks type a ...
- output jacks
  - max jacks type a ...
- vendor
- model
- serial number
- firmware

### router
- jacks
  - max jacks type a ...
    - MAC
- vendor
- model
- serial number
- firmware
- formfactor.rackable

### switch
- jacks
  - max jacks type a ...
- vendor
- model
- serial number
- firmware
- formfactor.rackable

### ram-module
- size
- speed
- type (ECC)
- formfactor.ram (DIMM, SODIMM)
- vendor
- model
- serial number

### cpu
- socket
- cores
- clockspeed
- vendor
- model
- serial number

### motherboard
- socket
- sockets (either socket or sockets must be filled)
  - max sockets type a ...
- formfactor.ram
- number of ram-slots
- connectors (==capability)
  - connector a .. (sata, sas, pcie, ...)
- vendor
- model
- serial number
- firmwares
  - firmware type a ...
  alternatively have a "devices" section and attach firmware to each device -> for onboard-nic etc

### nic
- (requires) connector of type ... (pcie) ( only when standalone, else integrated/on-board/embedded)
- jacks
  - max jacks of type a ...
    - MAC
    - reachable MACs (only when running)
      - ...
- firmware
- vendor
- model
- serial number

### storagedevice
- type (ssd,hdd)
- size
- connector/s (sata, sas, usb) (data+power)
- formfactor.storagedevice (2.5, 3.5, thumbdrive)
- vendor
- model
- serial number
- firmware

### cooler (cpu, but works for all the same)
- vendor
- model
- serial number
- formfactor.cooler
- input jack type (power)

### case
- formfactor.rackable (or formfactor.case, which extend rackable with small formfactors)
- vendor
- model
- serial number
- slots
  - max slots type a ...

### slot
- formfactor.slot
- configurations
  - max subslot type a ...
    max subslot type b ...
  - max subslot type a ...
    max subslot type b ...

### user
- name (login)
- prename (opt)
- surname (opt)
- login-methods
  - fingerprint
  - facial
  - authn
  - sms
  - mail
  - password
  - iris/coronal?
  - 2FA
  - 3rd party
- additional properties
  - f.e. title, mailaddress, address

### security-group
- name
- users
  - user a ...
- permissions
  - permission a ...
- additional properties
  - f.e. company unit
