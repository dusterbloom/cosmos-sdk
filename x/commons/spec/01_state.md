<!--
order: 1
-->

# State

The `x/commons` module keeps state of seven primary objects:

1. Commons metadata - e.g. NetworkZero
2. Denomination metadata - e.g. EUR
3. Commitments - intents to buy / sell quantified in Denom
4. Credit Limits - allowed negative balance  committer
5. Total committed value - The total value of commitments to the Commons and Denomination  
6. Maximum available credit - The max supply of credit available per Commons and Denomination 
7. Total supply of credit - Current credit issued and circulating per Commons and Denomination
8. Balances 

In addition, the `x/commons` module keeps the following indexes to manage the
aforementioned state:

* Commons Metadata Index: `0x0 | byte(commons.denom) -> ProtocolBuffer(Metadata)`
* Denom Metadata Index: `0x1 | byte(commons.denom) -> ProtocolBuffer(Metadata)`
* Commitments Index: `0x2 | byte(commons.denom) -> byte(amount)`
* Credit limits Index: `0x3 | byte(commons.denom) -> byte(amount)`
* Current balances Index: `0x4 | byte(address length) | []byte(address) | []byte(balance.commons.Denom) -> ProtocolBuffer(balance)`
* Reverse Denomination to Address Index: `0x05 | byte(commons.denom) | 0x00 | []byte(address) -> 0`



