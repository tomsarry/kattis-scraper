# kattis-scraper

🔎 Kattis web scraper.

(Currently) A CLI kattis Web Scraper that gives all solved problems of a user, along with their difficulty and their URL.

## Getting started

Make sure you have Go installed on your machine (1.14+), and that $PATH leads to your binary go files.

```
$ go get github.com/tomsarry/kattis-scraper
$ go install github.com/tomsarry/kattis-scraper
```

You should add a `.env` file with two parameters : 
```
EMAIL:email
PASSWORD:password
```
where _email_ is your Kattis email, and _password_ is your kattis password.

## How to use it

Simply run the following 
```
$ kattis-scrapper
```

Currently returns an alphabetically sorted list of your solved problems.

**⚠️ Note: Make sure to have the .env file in that location.**

### Example

```
$ kattis-scrapper
1: 3D Printed Statues, 1.9, https://open.kattis.com/problems/3dprinter
2: Accounting, 4.1, https://open.kattis.com/problems/bokforing
3: Are You Listening?, 2.6, https://open.kattis.com/problems/areyoulistening
4: Bachet's Game, 2.5, https://open.kattis.com/problems/bachetsgame
5: Backspace, 2.9, https://open.kattis.com/problems/backspace
6: Baloni, 3.4, https://open.kattis.com/problems/baloni
7: Bing It On, 3.7, https://open.kattis.com/problems/bing
8: Birds on a Wire, 3.3, https://open.kattis.com/problems/birds
9: Bounding Robots, 1.6, https://open.kattis.com/problems/boundingrobots
10: Bus Numbers, 3.1, https://open.kattis.com/problems/busnumbers
11: Closest Sums, 2.8, https://open.kattis.com/problems/closestsums
12: Cold-puter Science, 1.3, https://open.kattis.com/problems/cold
```

## Behaviour for Unexpected Results

Here are some explanations for weird results of the program:

### Difficulty of 0 on a problem ?
Scraper could not parse the difficulty of the problem, it may be related to some weird notations like this one (TODO):

(Unprecise difficulty)[https://github.com/tomsarry/kattis-scraper/tree/master/assets/pb_ex1.PNG]

### Not a valid link for a problem ?
Scraper could not parse the link associated with the problem.

## Todo's 
Here are the next steps for this project, ranked by difficulty (assumption):

### **EASY**
* Add sort by difficulty, increasing or decreasing. 
* Add sort by inverse alphabetical order. 

### **MEDIUM**
* Turn it into an API -> json response file.

### **HARD**
* Make it work with 100+ solved problems (multiple pages).

