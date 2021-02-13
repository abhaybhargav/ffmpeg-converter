## FFMpeg Bulk Converter

> Small utility to convert mp4s to any other ffmpeg supported format

### What does it do? 
* Given a directory, it recursively reads and identifies mp4 files
* Based on the format chosen in the CLI args `-format`, it will convery each mp4 file into the format chosen
* It will generate the output files in the directory selected in the `-out` path

### CLI Usage

```bash

ffmpeg-converter -input /some/directory -out /some/other/directory -format avi

```

> you need to have ffmpeg installed on your machine for this to work