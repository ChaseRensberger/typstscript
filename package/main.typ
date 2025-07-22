#import "template.typ": *

#show: screenplay.with(
  title: "{{.ScriptTitle}}",
)

#scene("FADE IN:")

#scene("INT. LOCATION - DAY")

#action[
  Describe the scene and action.
]

#dialogue_block[
  #character("CHARACTER NAME")
  #parenthetical("parenthetical")
  #line[This is how dialogue looks.]
]

#action[
  More action and description.
]
