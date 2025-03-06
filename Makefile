init:
	cp ./examples/a-real-pain-intro.typ ./main.typ

build:
	typst compile --font-path ./fonts main.typ

build-example:
	typst compile --font-path ./fonts examples/a-real-pain-intro.typ

clean-example:
	rm -f examples/a-real-pain-intro.pdf

watch:
	typst watch --font-path ./fonts main.typ

clean:
	rm -f main.pdf
