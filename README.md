<p align="center">
  <img src="https://raw.githubusercontent.com/m9s7/OISA_2x_sistem/main/media/logo.jpg?token=GHSAT0AAAAAABZE6CZGEK3XJ73U7S23SAOCY6MAGFA" width="300" height="300">
</p>

# OISA_2x_sistem

This is a software that offers a service of providing arbitrage opportunities to users via telegram messages. 
Bookies supported are all local serbian books: [Mozzart](https://www.mozzartbet.com/sr#/betting), [Maxbet](https://www.maxbet.rs/ibet-web-client/#/home#top), [Soccerbet](https://soccerbet.rs/#kladjenje), [Merkurxtip](https://www.merkurxtip.rs/desk/sr/sportsko-kladjenje).

Sports and two outcome games supported are:
- Soccer
  + total goals, total goals first half, total goals second half, 
  + those 3 for the home and away team individualy
  + GG, NG, 1GG, 1NG, 2GG, 2NG
- Basketball
  + game outcome with OT included
- Tennis
  + final outcome, 1st set outcome, 2st set outcome
  + tie break 1st set yes/no, tie break 2nd set yes/no
  + tie break in match yes/no
  
Once a user is added to a telegram group, he starts recieveing messages like this:

<img src="https://raw.githubusercontent.com/m9s7/OISA_2x_sistem/main/media/what_does_this_mean/origigi/1671719995940.jpg?token=GHSAT0AAAAAABZE6CZG7PFUFYPHZK3IVMQQY6MAHNQ" height="300">
  
Premium group users have 2 additional features
- an arbitrage calculator (they reply to an arbitrage with how much they want to invest and they get how much to put on each outcome)
- an overview of the state at all the other books by replying /extra to an arb they are interested in.

Those 2 features in action look like this:

<img src="https://github.com/m9s7/OISA_2x_sistem/blob/main/media/what_does_this_mean/origigi/1671727655564.jpg?raw=true" height="600"> <img src="https://github.com/m9s7/OISA_2x_sistem/blob/main/media/what_does_this_mean/origigi/1671727669502.jpg?raw=true" height="600">

If you would like to use this feel free to contact me to help you set it up.


Example of one loop of the program, scraping is parallelized so the scraping, merging and matching of **50,820** tips took less than 2 minutes:

<pre>
...scraping merkurxtip -  Tenis
...scraping soccerbet -  Tenis
...scraping maxb -  Tenis
...scraping mozzart -  Tenis
@MOZZART------------------
Matches scraped:  29      
Tips scraped:  76
--- 421.2794ms seconds ---
@SOCCERBET----------------
Matches scraped:  46      
Tips scraped:  67
--- 1.9728695s seconds ---
@MAXBET-------------------
Matches scraped:  42      
Tips scraped:  117        
--- 4.8968544s seconds ---
@MERKURXTIP---------------
Matches scraped:  54
Tips scraped:  103
--- 7.3106007s seconds ---      
... merging scraped data - Tenis
maxbet :  117
merkurxtip :  103
mozzart :  76
soccerbet :  67
Successfully merged: 50 records
---  30.301ms  ---
... finding arbitrage opportunities
-------------------------
---  544.7µs  ---
...scraping soccerbet -  Košarka
...scraping merkurxtip -  Košarka
...scraping maxb -  Košarka
...scraping mozzart -  Košarka
@MOZZART------------------
Matches scraped:  97
Tips scraped:  97
--- 906.2901ms seconds ---
@SOCCERBET----------------
Matches scraped:  172
Tips scraped:  151
--- 9.1809854s seconds ---
@MAXBET-------------------
Matches scraped:  156
Tips scraped:  154
--- 23.2470718s seconds ---
@MERKURXTIP---------------
Matches scraped:  323
Tips scraped:  313
--- 54.8678239s seconds ---
... merging scraped data - Košarka
merkurxtip :  313
maxbet :  154
soccerbet :  151
mozzart :  97
Successfully merged: 202 records
---  105.5096ms  ---
... finding arbitrage opportunities
-------------------------
---  538.6µs  ---
KOŠARKA, NB I
Szedeak vs Alba Fehervar
========================
KI_1_W/OT
kvota: 2.18 @ MOZZART
ulog = ukupno * 0.472
========================
KI_2_W/OT
kvota: 1.95 @ MAXBET
ulog = ukupno * 0.528
========================
Play first @ mozzart
ROI: 2.93%

...scraping merkurxtip -  Fudbal
...scraping soccerbet -  Fudbal
...scraping maxb -  Fudbal
...scraping mozzart -  Fudbal
@MOZZART------------------
Matches scraped:  127
Tips scraped:  3118
--- 13.5658308s seconds ---
@SOCCERBET----------------
Matches scraped:  446
Tips scraped:  9711
--- 28.3003609s seconds ---
@MAXBET-------------------
Matches scraped:  569
Tips scraped:  17029
--- 1m23.9319605s seconds ---
@MERKURXTIP---------------
Matches scraped:  671
Tips scraped:  19884
--- 1m47.7813139s seconds ---
... merging scraped data - Fudbal
merkurxtip :  19884
maxbet :  17029
soccerbet :  9711
mozzart :  3118
Successfully merged: 10803 records
---  6.4759886s  ---
... finding arbitrage opportunities
-------------------------
---  5.5655ms  ---
FUDBAL, PREMIER LEAGUE
Crystal Palace vs Newcastle
===========================
UG_TIM1_0-1
kvota: 1.32 @ SOCCERBET
ulog = ukupno * 0.765
===========================
UG_TIM1_2+
kvota: 4.30 @ MOZZART
ulog = ukupno * 0.235
===========================
Play first @ mozzart
ROI: 1.00%

FUDBAL, PREMIER LEAGUE
ASFA Yennega vs Sonabel Ouagadougou
=================================
GG
kvota: 2.62 @ MERKURXTIP
ulog = ukupno * 0.386
=================================
NG
kvota: 1.65 @ MOZZART
ulog = ukupno * 0.614
=================================
Play first @ merkurxtip
ROI: 1.24%

FUDBAL, CYMRU PREMIER
Cardiff Met vs Airbus
=====================
UG_1P_TIM1_0-1
kvota: 1.60 @ MOZZART
ulog = ukupno * 0.644
=====================
UG_1P_TIM1_2+
kvota: 2.90 @ MAXBET
ulog = ukupno * 0.356
=====================
Play first @ maxbet
ROI: 3.11%

</pre>
