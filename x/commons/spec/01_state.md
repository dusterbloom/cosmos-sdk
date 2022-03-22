<!--
order: 1
-->

# State

The `x/commons` module keeps state of eight primary objects:

1. Commons metadata - e.g. NetworkZero
2. Denomination metadata - e.g. EUR
3. Commitments - intents to buy / sell quantified in Denom
4. Credit Limits - allowed negative balances
5. Total committed value - The total value of commitments to the Commons and Denomination  
6. Maximum available credit - The max supply of credit available per Commons and Denomination 
7. Total supply of credit - Current credit issued and circulating per Commons and Denomination
8. Balances - Positive or negative number

In addition, the `x/commons` module keeps the following indexes to manage the
aforementioned state:

* Commons Metadata Index: `0x0 | byte(commons.denom) -> ProtocolBuffer(Metadata)`
* Denom Metadata Index: `0x1 | byte(commons.denom) -> ProtocolBuffer(Metadata)`
* Commitments Index: `0x2 |  byte(address length) | []byte(address) | []byte(commit.commons.Denom) -> ProtocolBuffer(commit)`
* Credit limits Index: `0x3 | byte(address length) | []byte(address) | []byte(limit.commons.Denom) -> ProtocolBuffer(limit)`
* Total committed value: `0x4 | byte(commits.commons.denom) -> byte(amount)`
* Max available credit: `0x5 | byte(maxCredit.commons.denom) -> byte(amount)`
* Total supply of credit value: `0x4 | byte(credit.commons.denom) -> byte(amount)`
* Balances Index: `0x6 | byte(address length) | []byte(address) | []byte(balance.commons.Denom) -> ProtocolBuffer(balance)`
* Reverse Denomination to Address Index: `0x05 | byte(commons.denom) | 0x00 | []byte(address) -> 0`



