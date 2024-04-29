# go-masonry-gallery

This repo contains a CLI program that scans a directory for images and generates a simple HTML gallery using the experimental CSS [Masonry layout](https://developer.mozilla.org/en-US/docs/Web/CSS/CSS_grid_layout/Masonry_layout). The gallery file generated is left in the directory, so in the future, you can just open it in a browser.

### Features
- Generate an HTML gallery with a masonry layout
- Randomize the order of the images
- Select and filter a subset of the images

 
## UI

```sh
gall ./some-dir  	# Generate default gallery in directory
gall ./some-dir -f 	# launch with firefox
gall ./some-dir -b 	# launch default browser
```

## Example
[The Works of  Zdzisław Beksiński](https://archive.org/details/ZdzislawBeksinski/)
```
PS C:\Users\Dom\Pictures\ZdzislawBeksinski> gall .
2024/04/29 11:40:36 Directory to be scanned: .
2024/04/29 11:40:36 Found 88 images in the directory
2024/04/29 11:40:36 1 non-images were excluded from the gallery
2024/04/29 11:40:36 Gallery created: C:\Users\Dom\Pictures\ZdzislawBeksinski\gal.html
```

<img src="./docs/example.png">

## Enabling CSS Masonry in Firefox
At the time of writing this, to use the Masonry layout in Firefox, navigate to [about:config](about:config) and set `layout.css.grid-template-masonry-value.enabled` to true.

## Notes
Example images generates with https://unsample.net/


## Template replacements
|Placeholder|Content|
|---|---|
|`<!-- GALLERY_CONTENTS -->`|Gallery images|
|`<!-- GALLERY_TITLE -->`|Gallery title|
|`<!-- GALLERY_INFO-->`|Gallery info|


## Context menu (windows)
Nilesoft shell https://nilesoft.org/docs/get-started

Add to shell.nss
```
menu(type='back|dir' mode="multiple"  title='dominickp/gall' image=\uE1F4)
{
	item(title = 'Generate Gallery' cmd='gall' arg='.')
	item(title = 'Generate Gallery (launch Firefox)' cmd='gall' arg='. -f')
}
```
