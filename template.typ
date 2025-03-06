#let screenplay(
  title: "UNTITLED",
  content
) = {
  // Set up the document with appropriate margins and font
  set page(
    paper: "us-letter",
    margin: (left: 1in, right: 1in, top: 1in, bottom: 1in),
  )
  
  set text(font: "Courier Prime", size: 12pt)
  set par(justify: false, leading: 0.5em)
  
  // Title page
  align(center)[
    #v(3in)
    #underline(offset: 0.2em)[#title]
  ]
  
  pagebreak()
  
  // Main content
  content
}

// Scene heading
#let scene(heading) = {
  v(0.5em)
  align(left)[
    #heading
  ]
  v(0.5em)
}

// Action/description
#let action(content) = {
  block(width: 100%, content)
  v(0.5em)
}

// Dialogue
#let dialogue(name, content) = {
  v(0.5em)
  align(center)[
    #block(width: 60%, [
      #align(center)[#name]
      #v(-0.2em)
      #align(left)[#content]
    ])
  ]
  v(0.5em)
}

// Parenthetical
#let parenthetical(content) = {
  align(center + horizon)[
    #text(style: "italic")[
      (#content)
    ]
  ]
}

// Transition
#let transition(content) = {
  align(right)[
    #text(weight: "bold")[#content]
  ]
  v(0.5em)
}