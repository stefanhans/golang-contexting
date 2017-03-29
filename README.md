# Reference Implementation for Contextinformation Routing Network (CRN)

Contextinformation Routing Network (CRN) is a communication framework enabling
an universal service to join matching contextinformation respectively its
communication partners.

The concept and the specifications are described in RFC documents:

 * [CRN Concepts](https://raw.githubusercontent.com/stefanhans/golang-contexting/master/RFC/CRN_Concepts.txt)
 (in large parts completed)

 * [CIP Specification](https://raw.githubusercontent.com/stefanhans/golang-contexting/master/RFC/CIP_Specification.txt)
 (in large parts completed)

 * [CIR Specification](https://raw.githubusercontent.com/stefanhans/golang-contexting/master/RFC/CIR_Specification.txt)
 (not yet mature)

The documents are still drafts and will be finalized, if the reference implementation will be.

The reference implementation in Go implements an open, unlimited and reactive
peer-to-peer overlay network design for all binary encodable kind of context
to connect its participants accordingly.

---
The main components of CRNs are the following:

### Contextinformation (CI)
Contextinformation (CI) refers mainly to the known terms information and context.
Due to the lack of an useful clear distinction between the two, CI is defined
here as information within its described context, i.e. context becomes part of
CI by describing it.
Another aspect of CI is the accuracy with regard to possible matches with other
CI. All of these has to be converted into a general format, in which the actual
meaning is not relevant to find matching CI.

### Contextinformation Coding (CIC)
Contextinformation Coding (CIC)is the conversion of CI into a binary format,
and vice versa.  CIC means both, the conversion rules (CIC-Ruleset) for a
particular type of CI and a concrete piece of encoded CI.
Every CIC has an identifier, called CIC-Number.
A CIC-Ruleset and all its encoded CI are linked by this CIC-Number. Encoded CI
has the form of two parallel bit strings of equal length.  This pair
consists of CIC-Content, an instance of the CIC-Ruleset, and CI-Mask, which is
used to define the accuracy of the searched match.  Then it is sufficient for
a bitwise match of two pieces of CI, if both CIC-Contents are equal or both
CI-Masks mark them as non-relevant. Let me illustrate this with an example.
Offer and Request are two communication roles, and location is a type of CI.
Offer says "I'm available for any Request with matching CI to contact me" and
Request says "I'm searching for any Offer with matching CI".
Both have an exact information about their location and can define a
surrounding area where Offer is available respectively Request is searching.
CIC-Content, as the encoded location, together with CI-Mask defines the
surrounding area by marking bits of CIC-Content as true in any case concerning
the match.  Here the CI of Offer and Request are matching, if the location of
one is in the surrounding area of the other and vice versa.

The calculation of a match between two CIs uses the following function resp. bitwise expression

    match(Offer, Request) = (NOT (Offer-Content XOR Request-Content)) OR (Offer-Mask AND Request-Mask)


### Contextinformation Packet (CIP)
Encoded CI is encapsulated in a datastructure named Contextinformation Packet (CIP).
A CIP is divided into three parts:

 * Header Data
 (static and dynamic)

 * Contextinformation
 (mainly dynamic)

 * Application Data
 (mainly dynamic)

All information which has to be transferred inside CRNs has to be encapsulated
within CIPs.

### Contextinformation Routing (CIR)
Contextinformation Routing (CIR) takes place in an overlay network built
normally on top of the TCP/IP layer. It is organized basically by using
CIC-Content as index. It is oriented towards known concepts of network routing,
peer-to-peer and others network principles and B-tree like datastructures.

Additionally it is committing to the four properties of reactive systems as described
in the [Reactive Manifesto](http://www.reactivemanifesto.org/):

![Reactive Manifesto](http://www.reactivemanifesto.org/images/reactive-traits.svg)

