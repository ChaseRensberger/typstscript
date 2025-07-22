#let screenplay(
  title: "UNTITLED",
  content
) = {
  set page(
    paper: "us-letter",
    margin: (left: 1in, right: 1in, top: 1in, bottom: 1in),
  )
  
  set text(font: "Courier Prime", size: 12pt)
  set par(justify: false, leading: 0.5em)
  
  align(center)[
    #v(3in)
    #underline(offset: 0.2em)[#title]
  ]
  
  pagebreak()
  
  content
}

#let scene(heading) = {
  v(0.5em)
  align(left)[
    #heading
  ]
  v(0.5em)
}

#let action(content) = {
  block(width: 100%, content)
  v(0.5em)
}

#let dialogue_block(content) = {
  v(0.5em)
  align(center)[
    #block(width: 60%, content)
  ]
  v(0.5em)
}

#let character(name) = {
  align(center)[#name]
}

#let line(content) = {
  align(left)[#content]
}

#let parenthetical(content) = {
  align(center)[
    (#content)
  ]
}
