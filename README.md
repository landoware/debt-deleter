# Debt Deleter

This is a work in progress!

This project was primarily built to create a tool to optimize paying off your loans given a fixed budgeted amount you can pay towards your debt using the a modification of the Avalanche Strategy.

## How it works

Suppose I budget a total of $250 each month for my various debts, which have these balances and interest rates
- $1,000 at 10%, minimum payment of $25
    - $0.27 accrued daily
- $7,500 at 5%, minimum payment of $100
    - $1.02 accrued daily
- $5,000 at 7.125%, minimum payment of $75
    - $0.97 accrued daily

After allocating $200 to the minimum payments of each debt, I've got $50 left. The strategy allocates this remainder to the debt with is accruing the most interest.
In this case it is the $7,500 loan at 5%. So, the extra $50 goes to that loan. Doing so every month will reduce the number of payments I need to make on this loan by 37% (91 to 57 payments)!

After that debt is paid off, the $150 that is now free goes to the next debt which is accruing the most interest, and the cycle repeats until my debt is deleted!

