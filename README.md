# golden-flask
Lets make a mini-project in golang

1. We will scrap the a mp3 audio site; taking input of https url as input.
2. Download mp3 files and send them to a folder.

Note: We donot plan to scrap to full site and perform mapping (maybe later?). 
Instead, we will keep the url such that we will just parse through the url input given
in CLI and store it at the location specified.

## Excluded file
The module ```gvars``` present in ./ is excluded. It has a variable ```AudioBooksDatabase```, which is an array of AudioBook struct.

The struct ```AudioBook``` is now a seperate package, sepearted grom ```gvars```.

#### Assumption
1. The pages with audios will be in format ```https://www.demosite.com/book/page1```, ```https://www.demosite.com/book/page2```, ```https://www.demosite.com/book/page3```, and so on for an Audio Book with site ```https://www.demosite.com/book/```, with audios on each page

## Credits
Credit to [nevermosby](https://gist.github.com/nevermosby) for creating code to [concurrently download content](https://gist.github.com/nevermosby/b54d473ea9153bb75eebd14d8d816544) from given URL. It was quite helpful.

## Notice/ Disclaimer
Use of this or any of the derived codes is up for private use and the writer/publisher of this code shall not be held liable or responsible for the way any underlying concepts or parts of code used as full or in parts.