/*
Package ctx implements the core functionality of Contextinformation Routing Networks (CRNs).

Contextinformation Paket (CIP)

Contextinformation Paket (CIP) is the datastructure to transfer contextinformation plus header and possibly application
data through Contextinformation Routing Networks (CRNs). All information which has to be transfered inside CRNs has to
be encapsulated within CIPs.

For an overview about CRNs please see Internet-Draft "Concepts of Contextinformation Routing Networks (CRNs)"

	https://github.com/stefanhans/golang-contexting/blob/master/RFC/CRN_Concepts.txt

A CIP is divided into three parts:

Header Data

The header data contains metadata and specifies, among other things, how the data of the following two parts has to be interpreted. The header data starts with a part of fixed size and static structure followed by a dynamic part. The static part defines also the dynamic part's type and its length.

	0                   1                   2                   3
	0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2
	+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
	|   purpose (1) |  profile (1)  |  version (1)  |  channel (1)  | |
	|                                                               | |
	|                            UUID (16)                          | |
	|                                                               | |
	|                                                               | fix
	|                          IP address (4)                       | |
	|            IP port (2)        |                               | |
	|                            time (8)                           | |
	|                               |   type (1)   |    size (1)    |---
	| ............................................................  | |
	| .............. additional data up to 255 bytes .............  | dyn
	| ............................................................  | |
	+---------------------------------------------------------------+


Contextinformation

The Contextinformation starts with a part of fixed size and static structure followed by a dynamic part. The static part defines also the dynamic part's type and its length. The dynamic part consists of CIC-Bricks only.

	0                   1                   2                   3
	0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2
	+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
	|   type (1)    |         root-CIC (2)          |   size (1)    | fix
	| ............................................................  | |
	| .............. additional data up to 510 bytes .............  | dyn
	| .............. i.e. up to 255 CIC-Bricks  ..................  | |
	| ............................................................  | |
	+---------------------------------------------------------------+


Application Data

The application data starts with a part of fixed size and static structure followed by a dynamic part. The static part defines only the dynamic part's type and its length.

	0                   1                   2                   3
	0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2
	+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
	|   type (1)    |   size (1)    | ............................  | fix
	| ............................................................  | |
	| .......... additional data up to 255 bytes (size) ..........  | dyn
	| ............................................................  | |
	+---------------------------------------------------------------+


All information which has to be transfered inside CRNs has to be encapsulated within CIPs.





External Packages

Code copied to and customized in uuid.go to avoid dependencies

	"UUID package for Go"

	Copyright (C) 2013-2015 by Maxim Bublis <b@codemonkey.ru>

	Package uuid provides implementation of Universally Unique Identifier (UUID).
	Supported versions are 1, 3, 4 and 5 (as specified in RFC 4122) and
	version 2 (as specified in DCE 1.1).

	https://github.com/satori/go.uuid

 */
package ctx

