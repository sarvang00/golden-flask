# golden-flask
Lets make a mini-project in golang

1. We will scrap the a mp3 audio site; taking input of https url as input.
2. Later phase: try to download mp3 and set them to a folder.

Amendment: We donot plan to scrap to full site and perform mapping (maybe later?). 
Instead, we will keep the url such that we will just parse through the url input given
in CLI and store it at the location specified.

## Excluded file
The module gvars present in ./ is excluded. It has a struct and the values stored in struct.

The struct was:
```
type AudioBook struct {
	BookName  string
	Author    string
	Reader    string
	HomeUrl   string
	StorePath string
}
```

The variable ```AudioBooksDatabase``` is an array of AudioBook struct. This may be restructured later for better visibility and understanding.

## Credits
Credit to [nevermosby](https://gist.github.com/nevermosby) for creating code to [concurrently download content](https://gist.github.com/nevermosby/b54d473ea9153bb75eebd14d8d816544) from given URL. It was quite helpful.

## Notice/ Disclaimer
Use of this or any of the derived codes is up for private use and the writer/publisher of this code shall not be held liable or responsible for the way any underlying concepts or parts of code in full or as whole are used.