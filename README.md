# go-masonry-gallery
 
## UI

```sh
gal ./some-dir  # Generate default gallery in directory
gal ./some-dir -f # launch with firefox
gal ./some-dir -b # launch default browser
```

## Firefox
Set `layout.css.grid-template-masonry-value.enabled` to true in `about:config`

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
