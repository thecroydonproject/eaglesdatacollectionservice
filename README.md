
## intro

Croydon (UK) tech city is maturing fast and the tech community is expanding.  There a number of interest groups that are

clustered around topic of interest. The GO & AWS interest group is a very active group. The motivation of this group is

exchange ideas and work with the community to solve tech problems.   To semi formalise the meet ups and work on concrete

 ideas, the GO-AWS Croydon group plans to have a work programme. in 2016, the project will stand up a modern API to

 give access to local football club historical results.


 the API project is split into a number work streams.  the data collection module is responsible for collecting data

 from various sites. This code is more or less what i have used in thecroydon project.



## Be good

Before you grab information from a website make sure you are not breaking legal or ethical rules.it's importnat to be

a virtual  citizen and protect the internet planet. if you are intending to gain from your data collection exploits i

suppose  you need to check with the owners and comply with IRP.The code in this module collects data relating to game

outcomes of my football team (some say, the best FC known to humanity). As expected  Crystal Palace FC historical

results are included in supporters  and semi official portals. i have not been able to find a RESTful API to query.




## How does it work?

- use http.GET to get hold of pages of interest.

- use html tokeniser to get to operate on html pages

- use regexp and strings libraries to format data into desired format

- Map data to struct and Marshal to Json


simple .. but if you have questions, I will be more than happy to help.


## It all about knowing the data structure

I spent a great deal of time trying to figure out the structure of my source websites. i found this is useful.

in fact i wrote a free text native to get the code logic clear in my head.

## License

 Free to use but refer to License file please.

 ## known issues

 none.

 ## to do

 - make the deep nested simpler (too much of combined for, if and switch logic )
 - improve error handling

## to run

- clone
- go run scraper.go

will produce json extracts into new folder scanfolder_2

Made in Croydon