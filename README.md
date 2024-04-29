# go-masonry-gallery

This repo contains a CLI program that scans a directory for images and generates a simple HTML gallery using the experimental CSS [Masonry layout](https://developer.mozilla.org/en-US/docs/Web/CSS/CSS_grid_layout/Masonry_layout). The gallery file generated is left in the directory, so in the future, you can just open it in a browser.

 
## UI

```sh
gall ./some-dir  	# Generate default gallery in directory
gall ./some-dir -f 	# launch with firefox
gall ./some-dir -b 	# launch default browser
```

## Firefox
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
